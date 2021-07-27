package api

import (
	apiv1 "basic/api/api_v1"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(apiVrsn string) (*Router, error) {
	r := Router{
		engine: gin.Default(),
	}

	v1 := r.engine.Group(apiVrsn)

	apiv1.AddIndex(v1)
	apiv1.AddRoles(v1)
	apiv1.AddUsers(v1)
	apiv1.AddAuth(v1)
	return &r, nil
}

func (r *Router) Run(addr ...string) error {
	return r.engine.Run(addr...)
}
