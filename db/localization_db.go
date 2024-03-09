package db

import (
	"strings"

	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LocalizationDbInterface interface {
	LocalizedLabel(tenant, isoCode string) LocalizedLabelRepositoryInterface
	LanguageList(tenant string) LanguageListRepositoryInterface
	Tenant() TenantRepositoryInterface
}

type LocalizationDb struct{}

func (db *LocalizationDb) LocalizedLabel(tenant, isoCode string) LocalizedLabelRepositoryInterface {
	baseRepo := odm.UnimplementedBootRepository[models.LocalizedLabelModel]{
		Database:       tenant + "_localization",
		CollectionName: strings.ToLower(isoCode) + "_labels",
	}

	return &LocalizedLabelRepository{baseRepo}
}

func (db *LocalizationDb) LanguageList(tenant string) LanguageListRepositoryInterface {
	baseRepo := odm.UnimplementedBootRepository[models.LanguageListModel]{
		Database:       tenant + "_localization",
		CollectionName: "language_list",
	}

	return &LanguageListRepository{baseRepo}
}

func (db *LocalizationDb) Tenant() TenantRepositoryInterface {
	baseRepo := odm.UnimplementedBootRepository[models.TenantModel]{
		Database:       "global",
		CollectionName: "tenant",
	}
	return &TenantRepository{baseRepo}
}
