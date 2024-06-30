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

type TeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedTeacherServiceServer
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TeacherService) Create(ctx context.Context, req *user_service.CreateTeacher) (resp *user_service.Teacher, err error) {

	i.log.Info("---CreateTeacher------>", logger.Any("req", req))

	pKey, err := i.strg.Teacher().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTeacher->Teacher->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Teacher().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TeacherService) GetByID(ctx context.Context, req *user_service.TeacherPrimaryKey) (resp *user_service.Teacher, err error) {

	i.log.Info("---GetTeacherByID------>", logger.Any("req", req))

	resp, err = i.strg.Teacher().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTeacherByID->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TeacherService) GetList(ctx context.Context, req *user_service.GetListTeacherRequest) (resp *user_service.GetListTeacherResponse, err error) {

	i.log.Info("---GetTeachers------>", logger.Any("req", req))

	resp, err = i.strg.Teacher().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTeachers->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TeacherService) Update(ctx context.Context, req *user_service.UpdateTeacher) (resp *user_service.Teacher, err error) {

	i.log.Info("---UpdateTeacher------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Teacher().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateTeacher--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Teacher().GetByPKey(ctx, &user_service.TeacherPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TeacherService) Delete(ctx context.Context, req *user_service.TeacherPrimaryKey) (resp *user_service.TeacherEmpty, err error) {

	i.log.Info("---DeleteTeacher------>", logger.Any("req", req))

	err = i.strg.Teacher().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.TeacherEmpty{}, nil
}
