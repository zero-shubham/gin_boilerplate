package accesscontrol

import (
	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Enforcer casbin.Enforcer

func SetupCasbin(db *gorm.DB) error {

	adp, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return err
	}
	enfcr, err := casbin.NewEnforcer("public/rbac_model.conf", adp)
	if err != nil {
		return err
	}
	Enforcer = *enfcr
	Enforcer.LoadPolicy()
	return nil
}
