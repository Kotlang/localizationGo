package main

import (
	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/service"
)

type Inject struct {
	LocalizationDb *db.LocalizationDb

	LocalizationService      *service.LocalizationService
	LocalizationAdminService *service.LocalizationAdminService
}

func NewInject() *Inject {
	inj := &Inject{}
	inj.LocalizationDb = &db.LocalizationDb{}

	inj.LocalizationService = service.NewLocalizationService(inj.LocalizationDb)
	inj.LocalizationAdminService = service.NewLocalizationAdminService(inj.LocalizationDb)
	return inj
}
