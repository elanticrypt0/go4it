package dirhunter

type Directory struct {
	ID        ID      `json:"id"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	HasParent bool    `json:"has_parent"`
	Parent    string  `json:"parent"`
	HasSubDir bool    `json:"has_subdirs"`
	HasFiles  bool    `json:"has_files"`
	Files     []*File `json:"files"`
}
