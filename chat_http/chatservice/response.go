package chatservice

type UploadPhotoResponse struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
