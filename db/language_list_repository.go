package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LanguageListRepository struct {
	odm.AbstractRepository[models.LanguageListModel]
}
