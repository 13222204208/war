package response

// Equipment 响应, 用于返回给前端
type Equipments struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	ParentID int    `json:"parentId"`
	Status   int    `json:"status"`
}
