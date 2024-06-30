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

type EventService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedEventServiceServer
}

func NewEventService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventService {
	return &EventService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *EventService) Create(ctx context.Context, req *schedule_service.CreateEvent) (resp *schedule_service.Event, err error) {

	i.log.Info("---CreateEvent------>", logger.Any("req", req))

	pKey, err := i.strg.Event().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateEvent->Event->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Event().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyEvent->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *EventService) GetByID(ctx context.Context, req *schedule_service.EventPrimaryKey) (resp *schedule_service.Event, err error) {

	i.log.Info("---GetEventByID------>", logger.Any("req", req))

	resp, err = i.strg.Event().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetEventByID->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *EventService) GetList(ctx context.Context, req *schedule_service.GetListEventRequest) (resp *schedule_service.GetListEventResponse, err error) {

	i.log.Info("---GetEvents------>", logger.Any("req", req))

	resp, err = i.strg.Event().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetEvents->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *EventService) Update(ctx context.Context, req *schedule_service.UpdateEvent) (resp *schedule_service.Event, err error) {

	i.log.Info("---UpdateEvent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Event().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateEvent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Event().GetByPKey(ctx, &schedule_service.EventPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetEvent->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *EventService) Delete(ctx context.Context, req *schedule_service.EventPrimaryKey) (resp *schedule_service.EventEmpty, err error) {

	i.log.Info("---DeleteEvent------>", logger.Any("req", req))

	err = i.strg.Event().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteEvent->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.EventEmpty{}, nil
}
