package dirhunter

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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
	dh.Fetch(filepath, false)
	return dh
}

// Fetch the content of a directory
// the results are stored in  Directories or Files

func (dh *DirHunter) Fetch(currentFilepath string, hasParent bool) ([]fs.DirEntry, error) {
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
				dh.AddDirectory(currentFilepath, dirItem)
				dh.Fetch(currentFilepath+"/"+dirItem.Name(), true)
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

func (dh *DirHunter) GetAllContent() {
}

// func (dh *DirHunter) GetAllContentOrdered() {
// 	for _, item := range dh.Directories {
// 		// preguntar en files si alguno lo tiene como path si lo tiene

// 	}
// }

func (dh *DirHunter) AddDirectory(parent string, fsDir fs.DirEntry) {
	// check if the current dir is the main
	hasSubdirs := dh.IsMainFilepath(fsDir.Name())

	path := parent + "/" + fsDir.Name()
	newDir := &Directory{
		ID:        uuid.New(),
		Name:      fsDir.Name(),
		Path:      path,
		HasParent: true,
		Parent:    parent,
		HasFiles:  false,
		HasSubDir: hasSubdirs,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (dh *DirHunter) AddMainFilepath() {
	newDir := &Directory{
		ID:        uuid.New(),
		Name:      dh.RemoveRootFromName(dh.MainFilepath),
		Path:      dh.MainFilepath,
		HasParent: false,
		HasFiles:  false,
		HasSubDir: false,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (dh *DirHunter) AddFile(dhParent *DirHunter, parentPath string, fsFile fs.DirEntry) {
	fullpath := parentPath + "/" + fsFile.Name()
	if parentPath != dh.MainFilepath {
		fullpath = parentPath + "/" + fsFile.Name()
	}
	fi, err := os.Stat(fullpath)

	if err != nil {
		fmt.Println(err)
		return
	}
	newFile := &File{
		ID:        uuid.New(),
		Name:      fsFile.Name(),
		Path:      parentPath,
		FullPath:  fullpath,
		Size:      fi.Size(),
		Extension: filepath.Ext(fi.Name())[1:],
	}
	currentDirKey, _ := dh.GetCurrentDir(parentPath)
	if currentDirKey > -1 {
		dh.Directories[currentDirKey].HasFiles = true
		dh.Directories[currentDirKey].Files = append(dh.Directories[currentDirKey].Files, newFile)
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
