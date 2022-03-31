package models

type AlComUserModel struct {
	User_id              string `json:"user_id"`
	Applicant_id         string `json:"applicant_id"`
	Pin_user             string `json:"pin_user"`
	Password             string `json:"password"`
	Last_chande_pwd      string `json:"last_chande_pwd"`
	Last_chande_pwd_time string `json:"last_chande_pwd_time"`
	Level_priv           string `json:"level_priv"`
	Pre_name             string `json:"pre_name"`
	User_name            string `json:"user_name"`
	Sur_name             string `json:"sur_name"`
	Mobile_no            string `json:"mobile_no"`
	Email                string `json:"email"`
	Last_active_date     string `json:"last_active_date"`
	Last_active_time     string `json:"last_active_time"`
	Last_active_ip       string `json:"Last_active_ip"`
	Address              string `json:"address"`
	Access_display       string `json:"access_display"`
	Access_add           string `json:"access_add"`
	Access_edit          string `json:"access_edit"`
	Access_delete        string `json:"access_delete"`
	Access_operation     string `json:"access_operation"`
	Access_quality       string `json:"access_quality"`
	Access_payment       string `json:"access_payment"`
	Access_post          string `json:"access_post"`
	Access_report        string `json:"access_report"`
	Access_audit         string `json:"access_audit"`
	Access_super_user    string `json:"access_super_user"`
	Date_time_ex         string `json:"date_time_ex"`
	Operate_user         string `json:"operate_user"`
	Operate_date         string `json:"operate_date"`
	Operate_time         string `json:"operate_time"`
	User_status          int    `json:"user_status"`
	Log_status           string `json:"log_status"`
	Log_user             string `json:"log_user"`
	Log_date             string `json:"log_date"`
	Log_time             string `json:"log_time"`
}

type AlLogConnectModel struct {
	Seqno       int    `json:"seqno" gorm:"primary_key"`
	Log_type    string `json:"log_type"`
	Log_ip      string `json:"log_ip"`
	Log_level   string `json:"log_level"`
	Log_user_id string `json:"log_user_id"`
	Log_status  string `json:"log_status"`
	Log_date    string `json:"log_date"`
	Log_time    string `json:"log_time"`
}

type ConfComConstModel struct {
	Config_id       string `json:"config_id" gorm:"primary_key"`
	Group_id        string `json:"group_id" gorm:"primary_key"`
	Create_date     string `json:"create_date"`
	Create_time     string `json:"create_time"`
	Create_user     string `json:"create_user"`
	Node_name_th    string `json:"node_name_th"`
	Node_name_en    string `json:"node_name_en"`
	Node_desc       string `json:"node_desc"`
	Node_url        string `json:"node_url"`
	Node_ref_value  string `json:"node_ref_value"`
	Node_ref_value2 string `json:"node_ref_value2"`
	Node_sort       int    `json:"node_sort"`
	Node_status     int    `json:"node_status"`
	Log_status      string `json:"log_status"`
	Log_date        string `json:"log_date"`
	Log_time        string `json:"log_time"`
	Log_user        string `json:"log_user"`
}

type FilesModel struct {
	AttachType     string `json:"AttachType"`
	ProjectID      string `json:"ProjectID" gorm:"primary_key"`
	FileOldName    string `json:"FileOldName"`
	FileName       string `json:"FileName"`
	FilePath       string `json:"FilePath"`
	FileSize       string `json:"FileSize"`
	Status         string `json:"Status"`
	CreateAt       string `json:"CreateAt"`
	CreateUser     string `json:"CreateUser"`
	LastUpdateAt   string `json:"LastUpdateAt"`
	LastUpdateUser string `json:"LastUpdateUser"`
}

func (FilesModel) TableName() string {
	return "files"
}

func (AlComUserModel) TableName() string {
	return "al_com_user"
}

func (AlLogConnectModel) TableName() string {
	return "al_log_connect"
}

func (ConfComConstModel) TableName() string {
	return "conf_com_const"
}

// func (ConfComConstModel) TableName() string {
// 	return "test_conf_com_const"
// }
