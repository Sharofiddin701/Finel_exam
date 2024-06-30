package service

import (
	"context"
	"schedule_service/config"
	schedule_service "schedule_service/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/pkg/logger"
	"schedule_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedGroupServiceServer
}

func NewGroupService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupService {
	return &GroupService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *GroupService) Create(ctx context.Context, req *schedule_service.CreateGroup) (resp *schedule_service.Group, err error) {

	i.log.Info("---CreateGroup------>", logger.Any("req", req))

	pKey, err := i.strg.Group().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateGroup->Group->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Group().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyGroup->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *GroupService) GetByID(ctx context.Context, req *schedule_service.GroupPrimaryKey) (resp *schedule_service.Group, err error) {

	i.log.Info("---GetGroupByID------>", logger.Any("req", req))

	resp, err = i.strg.Group().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetGroupByID->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *GroupService) GetList(ctx context.Context, req *schedule_service.GetListGroupRequest) (resp *schedule_service.GetListGroupResponse, err error) {

	i.log.Info("---GetGroups------>", logger.Any("req", req))

	resp, err = i.strg.Group().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetGroups->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *GroupService) Update(ctx context.Context, req *schedule_service.UpdateGroup) (resp *schedule_service.Group, err error) {

	i.log.Info("---UpdateGroup------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Group().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateGroup--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Group().GetByPKey(ctx, &schedule_service.GroupPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetGroup->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *GroupService) Delete(ctx context.Context, req *schedule_service.GroupPrimaryKey) (resp *schedule_service.GroupEmpty, err error) {

	i.log.Info("---DeleteGroup------>", logger.Any("req", req))

	err = i.strg.Group().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteGroup->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.GroupEmpty{}, nil
}
