package pondservice

type UploadPhotoResponse struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
type UploadFileResponse struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
