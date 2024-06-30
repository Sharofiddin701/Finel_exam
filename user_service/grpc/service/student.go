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

type StudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedStudentServiceServer
}

func NewStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *StudentService) Create(ctx context.Context, req *user_service.CreateStudent) (resp *user_service.Student, err error) {

	i.log.Info("---CreateStudent------>", logger.Any("req", req))

	pKey, err := i.strg.Student().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateStudent->Student->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Student().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyStudent->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StudentService) GetByID(ctx context.Context, req *user_service.StudentPrimaryKey) (resp *user_service.Student, err error) {

	i.log.Info("---GetStudentByID------>", logger.Any("req", req))

	resp, err = i.strg.Student().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudentByID->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StudentService) GetList(ctx context.Context, req *user_service.GetListStudentRequest) (resp *user_service.GetListStudentResponse, err error) {

	i.log.Info("---GetStudents------>", logger.Any("req", req))

	resp, err = i.strg.Student().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudents->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StudentService) Update(ctx context.Context, req *user_service.UpdateStudent) (resp *user_service.Student, err error) {

	i.log.Info("---UpdateStudent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Student().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateStudent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Student().GetByPKey(ctx, &user_service.StudentPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetStudent->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *StudentService) Delete(ctx context.Context, req *user_service.StudentPrimaryKey) (resp *user_service.StudentEmpty, err error) {

	i.log.Info("---DeleteStudent------>", logger.Any("req", req))

	err = i.strg.Student().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteStudent->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.StudentEmpty{}, nil
}

func (i *StudentService) GetStudetReport(ctx context.Context, req *user_service.StudentReportRequest) (resp *user_service.StudentReportResponse, err error) {

	i.log.Info("---GetStudentReport------>", logger.Any("req", req))

	resp, err = i.strg.Student().GetStudetReport(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudentReport->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}
