package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LocalizedLabelRepository struct {
	odm.AbstractRepository[models.LocalizedLabelModel]
}
