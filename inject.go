package main

import (
	"github.com/SaiNageswarS/go-api-boot/cloud"
	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/service"
)

type Inject struct {
	LocalizationDb *db.LocalizationDb
	CloudFns       cloud.Cloud

	LocalizationService      *service.LocalizationService
	LocalizationAdminService *service.LocalizationAdminService
}

func NewInject() *Inject {
	inj := &Inject{}
	inj.LocalizationDb = &db.LocalizationDb{}
	inj.CloudFns = &cloud.GCP{}

	inj.LocalizationService = service.NewLocalizationService(inj.LocalizationDb)
	inj.LocalizationAdminService = service.NewLocalizationAdminService(inj.LocalizationDb)
	return inj
}
