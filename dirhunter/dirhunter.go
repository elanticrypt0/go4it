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

func New(filepath string) DirHunter {
	dh := DirHunter{}
	dh.MainFilepath = filepath
	dh.AddMainFilepath()
	// to count the dirs
	dh.Fetch(filepath, dh.Directories[0])
	return dh
}

func (dh *DirHunter) RenameAll() {
	for _, dItem := range dh.Directories {
		if dItem.ID != 0 {
			dh.RenameDir(dItem)
		}
		if dItem.HasFiles {
			for _, fItem := range dItem.Files {
				fmt.Println(" -> file ", fItem.Path)
				dh.RenemeFile(dItem, fItem)
			}
		}
	}
}

// Fetch the content of a directory
// the results are stored in  Directories or Files

func (dh *DirHunter) Fetch(currentFilepath string, parent *Directory) ([]fs.DirEntry, error) {
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
				if dh.IsMainFilepath(currentFilepath) {
					dh.Directories[0].HasSubDir = true
				}
				dh.AddDirectory(parent.ID, currentFilepath, dirItem)
				dh.Fetch(currentFilepath+"/"+dirItem.Name(), dh.Directories[len(dh.Directories)-1])
			} else {
				dh.AddFile(dh, currentFilepath, dirItem)
			}
		}
		return content, nil
	}
}

func (dh *DirHunter) IsMainFilepath(dir string) bool {
	return dh.MainFilepath == dir
}

func (dh *DirHunter) AddDirectory(parentID uint, parent string, fsDir fs.DirEntry) {
	dirName := fsDir.Name()
	uid := uuid.New()

	// check if the current dir is the main
	hasSubdirs := dh.IsMainFilepath(dirName)
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

func (dh *DirHunter) AddMainFilepath() {
	newDir := &Directory{
		ID:        0,
		UID:       uuid.New(),
		Name:      dh.RemoveRootFromName(dh.MainFilepath),
		Path:      dh.MainFilepath,
		HasParent: false,
		HasFiles:  false,
		HasSubDir: false,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (dh *DirHunter) AddFile(dhParent *DirHunter, parentPath string, fsFile fs.DirEntry) {
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
	currentDirKey, _ := dh.GetCurrentDir(parentPath)
	if currentDirKey > -1 {
		dh.Directories[currentDirKey].HasFiles = true
		dh.Directories[currentDirKey].Files = append(dh.Directories[currentDirKey].Files, newFile)
	}
}

func (dh *DirHunter) RenameDir(dir *Directory) {
	parent := dh.Directories[dir.ParentID]
	oldPath := parent.Path + "/" + dir.Name

	dir.Name = dir.UID.String()
	dir.Path = parent.Path + "/" + dir.Name
	dir.ParentPath = parent.Path
	// renamed
	dh.Rename(oldPath, dir.Path)

}

func (dh *DirHunter) RenemeFile(dir *Directory, file *File) {
	fmt.Println("     PARENT .", dir.Path)
	oldFullpath := dir.Path + "/" + file.Path
	file.Path = file.UID.String() + "." + file.Extension
	file.FullPath = dir.Path + "/" + file.Path
	// renamed
	dh.Rename(oldFullpath, file.FullPath)
}

func (dh *DirHunter) Rename(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	if err != nil {
		fmt.Println(err)
	}
}

// returns the current dir key and value
func (dh *DirHunter) GetCurrentDir(path string) (int, *Directory) {
	for key, val := range dh.Directories {
		if val.Path == path {
			return key, val
		}
	}
	return -1, &Directory{}
}

// remove the "./" from name
func (dh *DirHunter) RemoveRootFromName(name string) string {
	return name[2:]
}

func (dh *DirHunter) GetAsJSON() string {
	json, _ := json.Marshal(dh.Directories)
	return string(json)
}

func (dh *DirHunter) GetDirAsJSON(d *Directory) string {
	json, _ := json.Marshal(d)
	return string(json)
}
