package services

import (
	"fmt"
	"sync"

	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var enforcer *casbin.Enforcer
var once sync.Once

func SetupCasbin(db *gorm.DB) error {
	var err error

	once.Do(func() {
		adp, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}
		enfcr, err := casbin.NewEnforcer("public/rbac_model.conf", adp)
		if err != nil {
			return
		}

		if enforcer == nil {
			enforcer = enfcr
		}
	})

	return err
}

func GetEnforcer() (*casbin.Enforcer, error) {
	if enforcer == nil {
		return nil, fmt.Errorf("enforcer not instantiated")
	}
	return enforcer, nil
}
