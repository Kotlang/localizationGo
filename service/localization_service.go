package service

import (
	"context"

	"github.com/kotlang/localizationGo/db"
	"github.com/kotlang/localizationGo/models"
	"github.com/kotlang/localizationGo/utils"
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

	languageListModel, err := utils.GetLanguageUsingISOCode(u.db.LanguageList(tenantDetails.Name), req.IsoCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ISO Code")
	}

	resChan, errChan := u.db.LocalizedLabel(tenantDetails.Name, languageListModel.Language).FindOneById(req.Key)
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

	languageListModel, err := utils.GetLanguageUsingISOCode(u.db.LanguageList(tenantDetails.Name), req.IsoCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ISO Code")
	}

	resChan, errChan := u.db.LocalizedLabel(tenantDetails.Name, languageListModel.Language).Find(bson.M{}, nil, 0, 0)
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
