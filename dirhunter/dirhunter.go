package dirhunter

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type DirHunter struct {
	MainFilepath string       `json:"main_filepath"`
	Directories  []*Directory `json:"directories"`
}

// Create a new instance and run the hunter
func New(filepath string) DirHunter {
	dh := DirHunter{}
	dh.MainFilepath = filepath
	dh.addMainFilepath()

	return dh
}

func (dh *DirHunter) Run(filepath string) {
	if filepath != "" {
		dh.MainFilepath = filepath
		dh.addMainFilepath()
	}
	dh.fetch(dh.MainFilepath, dh.Directories[0])
}

// Rename all directories
func (dh *DirHunter) renameAll() {
	for _, dItem := range dh.Directories {
		if dItem.ID != 0 {
			dh.renameDir(dItem)
		}
		if dItem.HasFiles {
			for _, fItem := range dItem.Files {
				dh.renemeFile(dItem, fItem)
			}
		}
	}
}

// fetch the content of a directory
// the results are stored in  Directories or Files
func (dh *DirHunter) fetch(currentFilepath string, parent *Directory) ([]fs.DirEntry, error) {
	// checks if dir exists
	_, err := os.Stat(currentFilepath)
	if err != nil {
		return nil, err
	} else {
		content, err := os.ReadDir(currentFilepath)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("%v", content)
		for _, dirItem := range content {
			// fmt.Printf("%v", dirItem.IsDir())
			if dirItem.IsDir() {
				if dh.isMainFilepath(currentFilepath) {
					dh.Directories[0].HasSubDir = true
				}
				dh.addDirectory(parent.ID, currentFilepath, dirItem)
				dh.fetch(currentFilepath+"/"+dirItem.Name(), dh.Directories[len(dh.Directories)-1])
			} else {
				dh.addFile(dh, currentFilepath, dirItem)
			}
		}
		return content, nil
	}
}

func (dh *DirHunter) isMainFilepath(dir string) bool {
	return dh.MainFilepath == dir
}

func (dh *DirHunter) addDirectory(parentID uint, parent string, fsDir fs.DirEntry) {
	dirName := fsDir.Name()
	uid := uuid.New()

	// check if the current dir is the main
	hasSubdirs := dh.isMainFilepath(dirName)
	// update the parent directory
	dh.Directories[parentID].HasSubDir = true
	path := parent + "/" + dirName

	newDir := &Directory{
		ID:         uint(len(dh.Directories)),
		UID:        uid,
		Name:       dirName,
		Path:       path,
		HasParent:  true,
		ParentPath: parent,
		ParentID:   parentID,
		HasFiles:   false,
		HasSubDir:  hasSubdirs,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (dh *DirHunter) addMainFilepath() {
	newDir := &Directory{
		ID:        0,
		UID:       uuid.New(),
		Name:      dh.removeRootFromName(dh.MainFilepath),
		Path:      dh.MainFilepath,
		HasParent: false,
		HasFiles:  false,
		HasSubDir: false,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (dh *DirHunter) addFile(dhParent *DirHunter, parentPath string, fsFile fs.DirEntry) {
	fileName := fsFile.Name()
	id := uuid.New()
	extension := filepath.Ext(fileName)[1:]
	fullpath := parentPath + "/" + fileName

	if parentPath != dh.MainFilepath {
		fullpath = parentPath + "/" + fileName
	}
	fi, err := os.Stat(fullpath)

	if err != nil {
		fmt.Println(err)
		return
	}
	newFile := &File{
		UID:       id,
		Name:      strings.TrimSuffix(fileName, filepath.Ext(fileName)),
		Path:      fileName,
		FullPath:  fullpath,
		Size:      fi.Size(),
		Extension: extension,
	}
	currentDirKey, _ := dh.getCurrentDir(parentPath)
	if currentDirKey > -1 {
		dh.Directories[currentDirKey].HasFiles = true
		dh.Directories[currentDirKey].Files = append(dh.Directories[currentDirKey].Files, newFile)
	}
}

func (dh *DirHunter) renameDir(dir *Directory) {
	parent := dh.Directories[dir.ParentID]
	oldPath := parent.Path + "/" + dir.Name

	dir.Name = dh.removeDash(dir.UID.String())
	dir.Path = parent.Path + "/" + dir.Name
	dir.ParentPath = parent.Path
	// renamed
	dh.rename(oldPath, dir.Path)
}

func (dh *DirHunter) renemeFile(dir *Directory, file *File) {
	oldFullpath := dir.Path + "/" + file.Path
	file.Path = dh.removeDash(file.UID.String()) + "." + file.Extension
	file.FullPath = dir.Path + "/" + file.Path
	// renamed
	dh.rename(oldFullpath, file.FullPath)
}

func (dh *DirHunter) rename(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	if err != nil {
		fmt.Println(err)
	}
}

func (dh *DirHunter) removeDash(olduuid string) string {
	return strings.ReplaceAll(olduuid, "-", "")
}

// returns the current dir key and value
func (dh *DirHunter) getCurrentDir(path string) (int, *Directory) {
	for key, val := range dh.Directories {
		if val.Path == path {
			return key, val
		}
	}
	return -1, &Directory{}
}

// remove the "./" from name
func (dh *DirHunter) removeRootFromName(name string) string {
	return name[2:]
}

// return a string in JSON format
func (dh *DirHunter) GetAllDirAsJSON() string {
	json, _ := json.Marshal(dh.Directories)
	return string(json)
}

// return one dir in JSON format
func (dh *DirHunter) GetDirAsJSON(d *Directory) string {
	json, _ := json.Marshal(d)
	return string(json)
}

func (dh *DirHunter) PrintDirData() {
	for _, dir := range dh.Directories {
		fmt.Printf("%v: \n", dir.Name)
		for _, dirFile := range dir.Files {
			fmt.Printf("    + %v\n", dirFile.Name)
		}
	}
}
