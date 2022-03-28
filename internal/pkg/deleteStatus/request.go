package deletestatus

type DeleteStatus struct {
	Config_id   string `json:"config_id"  example:"0001"`
	Group_id    string `json:"group_id"  example:"group1"`
	Node_status int    `json:"node_status"  example:"1"`
	Log_date    string `json:"log_date"  example:"2006-01-02"`
	Log_time    string `json:"log_time"  example:"00:00:00"`
	Log_user    string `json:"log_user"  example:"66861"`
}
