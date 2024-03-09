package db

import (
	"github.com/SaiNageswarS/go-api-boot/logger"
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/kotlang/localizationGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type TenantRepositoryInterface interface {
	odm.BootRepository[models.TenantModel]
	FindOneByToken(token string) chan *models.TenantModel
}

type TenantRepository struct {
	odm.UnimplementedBootRepository[models.TenantModel]
}

func (t *TenantRepository) FindOneByToken(token string) chan *models.TenantModel {
	ch := make(chan *models.TenantModel)

	go func() {
		resultChan, errorChan := t.FindOne(bson.M{"token": token})

		select {
		case res := <-resultChan:
			ch <- res
		case err := <-errorChan:
			logger.Error("Error fetching tenant info", zap.Error(err))
			ch <- nil
		}
	}()
	return ch
}
