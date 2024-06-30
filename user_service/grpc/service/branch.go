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

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BranchService) Create(ctx context.Context, req *user_service.CreateBranch) (resp *user_service.Branch, err error) {

	i.log.Info("---CreateBranch------>", logger.Any("req", req))

	pKey, err := i.strg.Branch().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBranch->Branch->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Branch().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) GetByID(ctx context.Context, req *user_service.BranchPrimaryKey) (resp *user_service.Branch, err error) {

	i.log.Info("---GetBranchByID------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBranchByID->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) GetList(ctx context.Context, req *user_service.GetListBranchRequest) (resp *user_service.GetListBranchResponse, err error) {

	i.log.Info("---GetBranchs------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBranchs->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) Update(ctx context.Context, req *user_service.UpdateBranch) (resp *user_service.Branch, err error) {

	i.log.Info("---UpdateBranch------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Branch().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Branch().GetByPKey(ctx, &user_service.BranchPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BranchService) Delete(ctx context.Context, req *user_service.BranchPrimaryKey) (resp *user_service.BranchEmpty, err error) {

	i.log.Info("---DeleteBranch------>", logger.Any("req", req))

	err = i.strg.Branch().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.BranchEmpty{}, nil
}
