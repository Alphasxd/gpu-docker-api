package model

type Bind struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
}

type VolumeCreate struct {
	Name string `json:"name,omitempty"`
	Size string `json:"size,omitempty"`
}