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

type DoTaskService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedDoTaskServiceServer
}

func NewDoTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *DoTaskService {
	return &DoTaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *DoTaskService) Create(ctx context.Context, req *schedule_service.CreateDoTask) (resp *schedule_service.DoTask, err error) {

	i.log.Info("---CreateDoTask------>", logger.Any("req", req))

	pKey, err := i.strg.DoTask().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateDoTask->DoTask->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.DoTask().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyDoTask->DoTask->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *DoTaskService) GetByID(ctx context.Context, req *schedule_service.DoTaskPrimaryKey) (resp *schedule_service.DoTask, err error) {

	i.log.Info("---GetDoTaskByID------>", logger.Any("req", req))

	resp, err = i.strg.DoTask().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetDoTaskByID->DoTask->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *DoTaskService) GetList(ctx context.Context, req *schedule_service.GetListDoTaskRequest) (resp *schedule_service.GetListDoTaskResponse, err error) {

	i.log.Info("---GetDoTasks------>", logger.Any("req", req))

	resp, err = i.strg.DoTask().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetDoTasks->DoTask->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *DoTaskService) Update(ctx context.Context, req *schedule_service.UpdateDoTask) (resp *schedule_service.DoTask, err error) {

	i.log.Info("---UpdateDoTask------>", logger.Any("req", req))

	rowsAffected, err := i.strg.DoTask().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateDoTask--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.DoTask().GetByPKey(ctx, &schedule_service.DoTaskPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetDoTask->DoTask->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *DoTaskService) Delete(ctx context.Context, req *schedule_service.DoTaskPrimaryKey) (resp *schedule_service.DoTaskEmpty, err error) {

	i.log.Info("---DeleteDoTask------>", logger.Any("req", req))

	err = i.strg.DoTask().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteDoTask->DoTask->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.DoTaskEmpty{}, nil
}
