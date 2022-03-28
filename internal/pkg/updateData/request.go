package updateData

type UpdateData struct {
	Config_id       string `json:"config_id"  example:"0001"`
	Group_id        string `json:"group_id"  example:"group1"`
	Create_date     string `json:"create_date"  example:"2022-01-01"`
	Create_time     string `json:"create_time"  example:"00:00:00"`
	Create_user     string `json:"create_user"  example:"66861"`
	Node_name_th    string `json:"node_name_th"  example:"th"`
	Node_name_en    string `json:"node_name_en"  example:"en"`
	Node_desc       string `json:"node_desc"  example:"newupdate"`
	Node_url        string `json:"node_url"  example:"/test/newurl"`
	Node_ref_value  string `json:"node_ref_value"  example:"newupdate"`
	Node_ref_value2 string `json:"node_ref_value2"  example:"newupdate"`
	Node_sort       int    `json:"node_sort"  example:"1"`
	Node_status     int    `json:"node_status"  example:"1"`
	Log_status      string `json:"log_status"  example:"newupdate"`
	Log_date  string `json:"log_date"  example:"2006-01-02"`
	Log_time        string `json:"log_time"  example:"00:00:00"`
	Log_user        string `json:"log_user"  example:"66861"`
}
