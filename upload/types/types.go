package types


type PreSignedUrl struct {
	UploadURL string `json:"uploadUrl"`
	Key       string `json:"key"`
}
