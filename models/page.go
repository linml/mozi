package models

// 分页查询条件
type PageParams struct {
	PageRow     int               `json:"page_row"`
	CurrentPage int               `json:"current_page"`
	Params      map[string]string `json:"params"`
}

// 分页查询结果
type PageResult struct {
	TotalCount  int           `json:"total_count"`
	PageCount   int           `json:"page_count"`
	CurrentPage int           `json:"current_page"`
	PageRow     int           `json:"page_row"`
	RecordData  []interface{} `json:"record_data"`
}
