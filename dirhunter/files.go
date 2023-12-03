package dirhunter

type File struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	FullPath  string `json:"fullpath"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
}
