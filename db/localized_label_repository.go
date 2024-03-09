package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
)

type LocalizedLabelRepositoryInterface interface {
	odm.BootRepository[models.LocalizedLabelModel]
}

type LocalizedLabelRepository struct {
	odm.UnimplementedBootRepository[models.LocalizedLabelModel]
}
