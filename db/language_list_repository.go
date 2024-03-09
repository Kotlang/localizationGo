package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LanguageListRepositoryInterface interface {
	odm.BootRepository[models.LanguageListModel]
}

type LanguageListRepository struct {
	odm.UnimplementedBootRepository[models.LanguageListModel]
}
