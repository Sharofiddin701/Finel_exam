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

type AssignStudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedAssignStudentServiceServer
}

func NewAssignStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AssignStudentService {
	return &AssignStudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *AssignStudentService) Create(ctx context.Context, req *schedule_service.CreateAssignStudent) (resp *schedule_service.AssignStudent, err error) {

	i.log.Info("---CreateAssignStudent------>", logger.Any("req", req))

	pKey, err := i.strg.AssignStudent().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateAssignStudent->AssignStudent->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.AssignStudent().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyAssignStudent->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AssignStudentService) GetByID(ctx context.Context, req *schedule_service.AssignStudentPrimaryKey) (resp *schedule_service.AssignStudent, err error) {

	i.log.Info("---GetAssignStudentByID------>", logger.Any("req", req))

	resp, err = i.strg.AssignStudent().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAssignStudentByID->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AssignStudentService) GetList(ctx context.Context, req *schedule_service.GetListAssignStudentRequest) (resp *schedule_service.GetListAssignStudentResponse, err error) {

	i.log.Info("---GetAssignStudents------>", logger.Any("req", req))

	resp, err = i.strg.AssignStudent().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAssignStudents->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AssignStudentService) Update(ctx context.Context, req *schedule_service.UpdateAssignStudent) (resp *schedule_service.AssignStudent, err error) {

	i.log.Info("---UpdateAssignStudent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.AssignStudent().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateAssignStudent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.AssignStudent().GetByPKey(ctx, &schedule_service.AssignStudentPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetAssignStudent->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *AssignStudentService) Delete(ctx context.Context, req *schedule_service.AssignStudentPrimaryKey) (resp *schedule_service.AssignStudentEmpty, err error) {

	i.log.Info("---DeleteAssignStudent------>", logger.Any("req", req))

	err = i.strg.AssignStudent().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteAssignStudent->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.AssignStudentEmpty{}, nil
}
