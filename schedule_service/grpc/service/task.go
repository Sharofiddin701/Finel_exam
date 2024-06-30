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

type TaskService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedTaskServiceServer
}

func NewTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TaskService {
	return &TaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TaskService) Create(ctx context.Context, req *schedule_service.CreateTask) (resp *schedule_service.Task, err error) {

	i.log.Info("---CreateTask------>", logger.Any("req", req))

	pKey, err := i.strg.Task().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTask->Task->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Task().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTask->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TaskService) GetByID(ctx context.Context, req *schedule_service.TaskPrimaryKey) (resp *schedule_service.Task, err error) {

	i.log.Info("---GetTaskByID------>", logger.Any("req", req))

	resp, err = i.strg.Task().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTaskByID->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TaskService) GetList(ctx context.Context, req *schedule_service.GetListTaskRequest) (resp *schedule_service.GetListTaskResponse, err error) {

	i.log.Info("---GetTasks------>", logger.Any("req", req))

	resp, err = i.strg.Task().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTasks->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TaskService) Update(ctx context.Context, req *schedule_service.UpdateTask) (resp *schedule_service.Task, err error) {

	i.log.Info("---UpdateTask------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Task().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateTask--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Task().GetByPKey(ctx, &schedule_service.TaskPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetTask->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TaskService) Delete(ctx context.Context, req *schedule_service.TaskPrimaryKey) (resp *schedule_service.TaskEmpty, err error) {

	i.log.Info("---DeleteTask------>", logger.Any("req", req))

	err = i.strg.Task().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTask->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.TaskEmpty{}, nil
}
