package service

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/auth"
	"github.com/kotlang/localizationGo/db"
	pb "github.com/kotlang/localizationGo/generated"
	"github.com/kotlang/localizationGo/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ALLOWED_USERS = map[string]bool{
	"NzAyMjM3NDU2OQ==": true,
	"ODE0NzQ5MDE0NA==": true,
	"ODA1NDk4NDM1MA==": true,
}

type LocalizationAdminService struct {
	pb.UnimplementedLocalizationAdminServer
	db *db.LocalizationDb
}

func NewLocalizationAdminService(db *db.LocalizationDb) *LocalizationAdminService {
	return &LocalizationAdminService{db: db}
}

func (u *LocalizationAdminService) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	userId, tenant := auth.GetUserIdAndTenant(ctx)

	if val, ok := ALLOWED_USERS[userId]; !ok || !val {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to add labels")
	}

	<-u.db.LocalizedLabel(tenant, req.Language).Save(&models.LocalizedLabelModel{Key: req.Key, Translation: req.Value})
	return &pb.AddLabelResponse{Status: "success"}, nil
}

func (u *LocalizationAdminService) AddLanguage(ctx context.Context, req *pb.AddLanguageRequest) (*pb.AddLanguageResponse, error) {
	userId, tenant := auth.GetUserIdAndTenant(ctx)

	if val, ok := ALLOWED_USERS[userId]; !ok || !val {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to add languages")
	}

	<-u.db.LanguageList(tenant).Save(&models.LanguageListModel{Language: req.Language, IsoCode: req.IsoCode})
	return &pb.AddLanguageResponse{Status: "success"}, nil
}
