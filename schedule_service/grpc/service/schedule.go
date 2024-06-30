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

type ScheduleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedScheduleServiceServer
}

func NewScheduleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScheduleService {
	return &ScheduleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ScheduleService) Create(ctx context.Context, req *schedule_service.CreateSchedule) (resp *schedule_service.Schedule, err error) {

	i.log.Info("---CreateSchedule------>", logger.Any("req", req))

	pKey, err := i.strg.Schedule().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSchedule->Schedule->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Schedule().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeySchedule->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScheduleService) GetByID(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (resp *schedule_service.Schedule, err error) {

	i.log.Info("---GetScheduleByID------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScheduleByID->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScheduleService) GetList(ctx context.Context, req *schedule_service.GetListScheduleRequest) (resp *schedule_service.GetListScheduleResponse, err error) {

	i.log.Info("---GetSchedules------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSchedules->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScheduleService) Update(ctx context.Context, req *schedule_service.UpdateSchedule) (resp *schedule_service.Schedule, err error) {

	i.log.Info("---UpdateSchedule------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Schedule().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateSchedule--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Schedule().GetByPKey(ctx, &schedule_service.SchedulePrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetSchedule->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ScheduleService) Delete(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (resp *schedule_service.ScheduleEmpty, err error) {

	i.log.Info("---DeleteSchedule------>", logger.Any("req", req))

	err = i.strg.Schedule().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSchedule->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.ScheduleEmpty{}, nil
}

func (i *ScheduleService) GetScheduleReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error) {

	i.log.Info("---GetScheduleReport------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetScheduleReport(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScheduleReport->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}
func (i *ScheduleService) GetScheduleMonthReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error) {

	i.log.Info("---GetScheduleMonthReport------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetScheduleMonthReport(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScheduleMonthReport->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}
