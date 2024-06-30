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

type ManagerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedManagerServiceServer
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ManagerService) Create(ctx context.Context, req *user_service.CreateManager) (resp *user_service.Manager, err error) {

	i.log.Info("---CreateManager------>", logger.Any("req", req))

	pKey, err := i.strg.Manager().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateManager->Manager->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Manager().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyManager->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ManagerService) GetByID(ctx context.Context, req *user_service.ManagerPrimaryKey) (resp *user_service.Manager, err error) {

	i.log.Info("---GetManagerByID------>", logger.Any("req", req))

	resp, err = i.strg.Manager().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetManagerByID->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ManagerService) GetList(ctx context.Context, req *user_service.GetListManagerRequest) (resp *user_service.GetListManagerResponse, err error) {

	i.log.Info("---GetManagers------>", logger.Any("req", req))

	resp, err = i.strg.Manager().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetManagers->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ManagerService) Update(ctx context.Context, req *user_service.UpdateManager) (resp *user_service.Manager, err error) {

	i.log.Info("---UpdateManager------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Manager().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateManager--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Manager().GetByPKey(ctx, &user_service.ManagerPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetManager->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ManagerService) Delete(ctx context.Context, req *user_service.ManagerPrimaryKey) (resp *user_service.ManagerEmpty, err error) {

	i.log.Info("---DeleteManager------>", logger.Any("req", req))

	err = i.strg.Manager().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteManager->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.ManagerEmpty{}, nil
}
