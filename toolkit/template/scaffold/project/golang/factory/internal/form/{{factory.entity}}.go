package form

type {{ factory.entity|title }}Create struct {
	When  string `json:"when" binding:"required" example:"时间"`
	Where string `json:"where" binding:"required" example:"地点"`
	Who   string `json:"who" binding:"required" example:"人物"`
	What  string `json:"what" binding:"required" example:"事件"`
	How   string `json:"how" binding:"required" example:"过程"`
}

type {{ factory.entity|title }}Update struct {
	When  string `json:"when" example:"时间"`
	Where string `json:"where" example:"地点"`
	Who   string `json:"who" example:"人物"`
	What  string `json:"what" example:"事件"`
	How   string `json:"how" example:"过程"`
}
