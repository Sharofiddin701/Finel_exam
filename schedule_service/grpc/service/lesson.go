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

type LessonService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedLessonServiceServer
}

func NewLessonService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *LessonService {
	return &LessonService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *LessonService) Create(ctx context.Context, req *schedule_service.CreateLesson) (resp *schedule_service.Lesson, err error) {

	i.log.Info("---CreateLesson------>", logger.Any("req", req))

	pKey, err := i.strg.Lesson().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateLesson->Lesson->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Lesson().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyLesson->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *LessonService) GetByID(ctx context.Context, req *schedule_service.LessonPrimaryKey) (resp *schedule_service.Lesson, err error) {

	i.log.Info("---GetLessonByID------>", logger.Any("req", req))

	resp, err = i.strg.Lesson().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetLessonByID->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *LessonService) GetList(ctx context.Context, req *schedule_service.GetListLessonRequest) (resp *schedule_service.GetListLessonResponse, err error) {

	i.log.Info("---GetLessons------>", logger.Any("req", req))

	resp, err = i.strg.Lesson().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetLessons->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *LessonService) Update(ctx context.Context, req *schedule_service.UpdateLesson) (resp *schedule_service.Lesson, err error) {

	i.log.Info("---UpdateLesson------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Lesson().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateLesson--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Lesson().GetByPKey(ctx, &schedule_service.LessonPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetLesson->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *LessonService) Delete(ctx context.Context, req *schedule_service.LessonPrimaryKey) (resp *schedule_service.LessonEmpty, err error) {

	i.log.Info("---DeleteLesson------>", logger.Any("req", req))

	err = i.strg.Lesson().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteLesson->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.LessonEmpty{}, nil
}
