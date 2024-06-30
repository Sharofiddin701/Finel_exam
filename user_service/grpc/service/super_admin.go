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

type SuperAdminService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedSuperAdminServiceServer
}

func NewSuperAdminService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SuperAdminService {
	return &SuperAdminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *SuperAdminService) Create(ctx context.Context, req *user_service.CreateSuperAdmin) (resp *user_service.SuperAdmin, err error) {

	i.log.Info("---CreateSuperAdmin------>", logger.Any("req", req))

	pKey, err := i.strg.SuperAdmin().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSuperAdmin->SuperAdmin->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.SuperAdmin().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeySuperAdmin->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SuperAdminService) GetByID(ctx context.Context, req *user_service.SuperAdminPrimaryKey) (resp *user_service.SuperAdmin, err error) {

	i.log.Info("---GetSuperAdminByID------>", logger.Any("req", req))

	resp, err = i.strg.SuperAdmin().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSuperAdminByID->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SuperAdminService) GetList(ctx context.Context, req *user_service.GetListSuperAdminRequest) (resp *user_service.GetListSuperAdminResponse, err error) {

	i.log.Info("---GetSuperAdmins------>", logger.Any("req", req))

	resp, err = i.strg.SuperAdmin().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSuperAdmins->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SuperAdminService) Update(ctx context.Context, req *user_service.UpdateSuperAdmin) (resp *user_service.SuperAdmin, err error) {

	i.log.Info("---UpdateSuperAdmin------>", logger.Any("req", req))

	rowsAffected, err := i.strg.SuperAdmin().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateSuperAdmin--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.SuperAdmin().GetByPKey(ctx, &user_service.SuperAdminPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetSuperAdmin->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SuperAdminService) Delete(ctx context.Context, req *user_service.SuperAdminPrimaryKey) (resp *user_service.SuperAdminEmpty, err error) {

	i.log.Info("---DeleteSuperAdmin------>", logger.Any("req", req))

	err = i.strg.SuperAdmin().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSuperAdmin->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.SuperAdminEmpty{}, nil
}
