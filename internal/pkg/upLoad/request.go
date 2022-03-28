package upload

type Uploadfile struct {
	AttachType     string `json:"AttachType"  example:"head"`
	ProjectID      string `json:"ProjectID"  example:"ProID"`
	FileOldName    string `json:"FileOldName"  example:"Name"`
	FileName       string `json:"FileName"  example:"commandline"`
	FilePath       string `json:"FilePath"  example:"test"`
	FileSize       string `json:"FileSize"  example:"10MB"`
	Status         string `json:"Status"  example:"1"`
	CreateAt       string `json:"CreateAt"  example:"2006-01-02"`
	CreateUser     string `json:"CreateUser"  example:"66866"`
	LastUpdateAt   string `json:"LastUpdateAt"  example:"2006-01-02"`
	LastUpdateUser string `json:"LastUpdateUser"  example:"66866"`
}
