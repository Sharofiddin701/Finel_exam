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

type AdministratorService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedAdministratorServiceServer
}

func NewAdministratorService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AdministratorService {
	return &AdministratorService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *AdministratorService) Create(ctx context.Context, req *user_service.CreateAdministrator) (resp *user_service.Administrator, err error) {

	i.log.Info("---CreateAdministrator------>", logger.Any("req", req))

	pKey, err := i.strg.Administrator().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateAdministrator->Administrator->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Administrator().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyAdministrator->Administrator->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AdministratorService) GetByID(ctx context.Context, req *user_service.AdministratorPrimaryKey) (resp *user_service.Administrator, err error) {

	i.log.Info("---GetAdministratorByID------>", logger.Any("req", req))

	resp, err = i.strg.Administrator().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAdministratorByID->Administrator->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AdministratorService) GetList(ctx context.Context, req *user_service.GetListAdministratorRequest) (resp *user_service.GetListAdministratorResponse, err error) {

	i.log.Info("---GetAdministrators------>", logger.Any("req", req))

	resp, err = i.strg.Administrator().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAdministrators->Administrator->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AdministratorService) Update(ctx context.Context, req *user_service.UpdateAdministrator) (resp *user_service.Administrator, err error) {

	i.log.Info("---UpdateAdministrator------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Administrator().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateAdministrator--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Administrator().GetByPKey(ctx, &user_service.AdministratorPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetAdministrator->Administrator->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *AdministratorService) Delete(ctx context.Context, req *user_service.AdministratorPrimaryKey) (resp *user_service.AdministratorEmpty, err error) {

	i.log.Info("---DeleteAdministrator------>", logger.Any("req", req))

	err = i.strg.Administrator().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteAdministrator->Administrator->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.AdministratorEmpty{}, nil
}
