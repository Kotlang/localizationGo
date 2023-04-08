package utils

import (
	"strings"

	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLanguageUsingISOCode(languageList *db.LanguageListRepository, isoCode string) (*models.LanguageListModel, error) {
	resChan, errChan := languageList.FindOne(bson.M{"isocode": isoCode})
	select {
	case res := <-resChan:
		res.Language = strings.ToLower(res.Language)
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}
