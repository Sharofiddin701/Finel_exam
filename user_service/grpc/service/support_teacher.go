package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/pkg/logger"
	"user_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedSupportTeacherServiceServer
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *SupportTeacherService) Create(ctx context.Context, req *user_service.CreateSupportTeacher) (resp *user_service.SupportTeacher, err error) {

	i.log.Info("---CreateSupportTeacher------>", logger.Any("req", req))

	pKey, err := i.strg.SupportTeacher().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSupportTeacher->SupportTeacher->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.SupportTeacher().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeySupportTeacher->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupportTeacherService) GetByID(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) (resp *user_service.SupportTeacher, err error) {

	i.log.Info("---GetSupportTeacherByID------>", logger.Any("req", req))

	resp, err = i.strg.SupportTeacher().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupportTeacherByID->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupportTeacherService) GetList(ctx context.Context, req *user_service.GetListSupportTeacherRequest) (resp *user_service.GetListSupportTeacherResponse, err error) {

	i.log.Info("---GetSupportTeachers------>", logger.Any("req", req))

	resp, err = i.strg.SupportTeacher().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupportTeachers->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupportTeacherService) Update(ctx context.Context, req *user_service.UpdateSupportTeacher) (resp *user_service.SupportTeacher, err error) {

	i.log.Info("---UpdateSupportTeacher------>", logger.Any("req", req))

	rowsAffected, err := i.strg.SupportTeacher().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateSupportTeacher--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.SupportTeacher().GetByPKey(ctx, &user_service.SupportTeacherPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetSupportTeacher->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SupportTeacherService) Delete(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) (resp *user_service.SupportTeacherEmpty, err error) {

	i.log.Info("---DeleteSupportTeacher------>", logger.Any("req", req))

	err = i.strg.SupportTeacher().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSupportTeacher->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.SupportTeacherEmpty{}, nil
}
