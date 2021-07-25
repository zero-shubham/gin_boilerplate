package accesscontrol

import (
	"fmt"

	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var enforcer *casbin.Enforcer

func SetupCasbin(db *gorm.DB) error {

	adp, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return err
	}
	enfcr, err := casbin.NewEnforcer("public/rbac_model.conf", adp)
	if err != nil {
		return err
	}

	if enforcer == nil {
		enforcer = enfcr
	}
	enforcer.LoadPolicy()
	return nil
}

func GetEnforcer() (*casbin.Enforcer, error) {
	if enforcer == nil {
		return nil, fmt.Errorf("enforcer not instantiated")
	}
	return enforcer, nil
}
