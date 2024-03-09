package service

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/auth"
	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/extensions"
	pb "github.com/kotlang/localizationGo/generated"
	"github.com/kotlang/localizationGo/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LocalizationAdminService struct {
	pb.UnimplementedLocalizationAdminServer
	db *db.LocalizationDb
}

func NewLocalizationAdminService(db *db.LocalizationDb) *LocalizationAdminService {
	return &LocalizationAdminService{db: db}
}

func (u *LocalizationAdminService) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	userId, tenant := auth.GetUserIdAndTenant(ctx)

	if !<-extensions.IsUserAdmin(ctx, userId) {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to add labels")
	}

	<-u.db.LocalizedLabel(tenant, req.Language).Save(&models.LocalizedLabelModel{Key: req.Key, Translation: req.Value})
	return &pb.AddLabelResponse{Status: "success"}, nil
}

func (u *LocalizationAdminService) AddLanguage(ctx context.Context, req *pb.AddLanguageRequest) (*pb.AddLanguageResponse, error) {
	userId, tenant := auth.GetUserIdAndTenant(ctx)

	if !<-extensions.IsUserAdmin(ctx, userId) {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to add labels")
	}

	<-u.db.LanguageList(tenant).Save(&models.LanguageListModel{Language: req.Language, IsoCode: req.IsoCode})
	return &pb.AddLanguageResponse{Status: "success"}, nil
}
