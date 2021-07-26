package schemas

type AddRolesPost struct {
	Roles []string `json:"roles" binding:"required"`
}
