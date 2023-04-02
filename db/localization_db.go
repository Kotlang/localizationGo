package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LocalizationDb struct{}

func (db *LocalizationDb) LocalizedLabel(tenant, language string) *LocalizedLabelRepository {
	baseRepo := odm.AbstractRepository[models.LocalizedLabelModel]{
		Database:       tenant + "_localization",
		CollectionName: language + "_labels",
	}

	return &LocalizedLabelRepository{baseRepo}
}

func (db *LocalizationDb) LanguageList(tenant string) *LanguageListRepository {
	baseRepo := odm.AbstractRepository[models.LanguageListModel]{
		Database:       tenant + "_localization",
		CollectionName: "language_list",
	}

	return &LanguageListRepository{baseRepo}
}

func (db *LocalizationDb) Tenant() *TenantRepository {
	baseRepo := odm.AbstractRepository[models.TenantModel]{
		Database:       "global",
		CollectionName: "tenant",
	}
	return &TenantRepository{baseRepo}
}
