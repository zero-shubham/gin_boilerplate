package schemas

type AuthPost struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type AuthPostResp struct {
	Token string `json:"token"`
}
