package service

import (
	"context"

	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/models"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/kotlang/localizationGo/generated"
)

type LocalizationService struct {
	pb.UnimplementedLabelLocalizationServer
	db *db.LocalizationDb
}

func NewLocalizationService(db *db.LocalizationDb) *LocalizationService {
	return &LocalizationService{db: db}
}

// removing auth interceptor
func (u *LocalizationService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (u *LocalizationService) GetLabel(ctx context.Context, req *pb.GetLabelRequest) (*pb.LocalizedLabel, error) {
	if len(req.Domain) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid Domain Token")
	}

	tenantDetails := <-u.db.Tenant().FindOneByToken(req.Domain)
	if tenantDetails == nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid domain token")
	}

	resChan, errChan := u.db.LocalizedLabel(tenantDetails.Name, req.IsoCode).FindOneById(req.Key)
	select {
	case res := <-resChan:
		return &pb.LocalizedLabel{Key: res.Key, Value: res.Translation}, nil
	case err := <-errChan:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (u *LocalizationService) GetAllLabelsByISOCode(ctx context.Context, req *pb.GetAllLabelsByISOCodeRequest) (*pb.LocalizedLabelsResponse, error) {
	if len(req.Domain) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid Domain Token")
	}

	tenantDetails := <-u.db.Tenant().FindOneByToken(req.Domain)
	if tenantDetails == nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid domain token")
	}

	resChan, errChan := u.db.LocalizedLabel(tenantDetails.Name, req.IsoCode).Find(bson.M{}, nil, 0, 0)
	select {
	case res := <-resChan:
		labels := funk.Map(res, func(label models.LocalizedLabelModel) *pb.LocalizedLabel {
			return &pb.LocalizedLabel{Key: label.Key, Value: label.Translation}
		}).([]*pb.LocalizedLabel)

		return &pb.LocalizedLabelsResponse{LocalizedLabelList: labels}, nil
	case err := <-errChan:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (u *LocalizationService) GetAllLanguages(ctx context.Context, req *pb.GetAllLanguagesRequest) (*pb.GetAllLanguagesResponse, error) {
	if len(req.Domain) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid Domain Token")
	}

	tenantDetails := <-u.db.Tenant().FindOneByToken(req.Domain)
	if tenantDetails == nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid domain token")
	}

	resChan, errChan := u.db.LanguageList(tenantDetails.Name).Find(bson.M{}, nil, 0, 0)

	select {
	case res := <-resChan:
		languages := funk.Map(res, func(lang models.LanguageListModel) *pb.LanguageDetail {
			return &pb.LanguageDetail{Language: lang.Language, IsoCode: lang.IsoCode}
		}).([]*pb.LanguageDetail)

		return &pb.GetAllLanguagesResponse{LanguageList: languages}, nil
	case err := <-errChan:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
