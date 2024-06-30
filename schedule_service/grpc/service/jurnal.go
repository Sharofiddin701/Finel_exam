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

type JurnalService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedJurnalServiceServer
}

func NewJurnalService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *JurnalService {
	return &JurnalService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *JurnalService) Create(ctx context.Context, req *schedule_service.CreateJurnal) (resp *schedule_service.Jurnal, err error) {

	i.log.Info("---CreateJurnal------>", logger.Any("req", req))

	pKey, err := i.strg.Jurnal().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateJurnal->Jurnal->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Jurnal().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyJurnal->Jurnal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *JurnalService) GetByID(ctx context.Context, req *schedule_service.JurnalPrimaryKey) (resp *schedule_service.Jurnal, err error) {

	i.log.Info("---GetJurnalByID------>", logger.Any("req", req))

	resp, err = i.strg.Jurnal().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetJurnalByID->Jurnal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *JurnalService) GetList(ctx context.Context, req *schedule_service.GetListJurnalRequest) (resp *schedule_service.GetListJurnalResponse, err error) {

	i.log.Info("---GetJurnals------>", logger.Any("req", req))

	resp, err = i.strg.Jurnal().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetJurnals->Jurnal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *JurnalService) Update(ctx context.Context, req *schedule_service.UpdateJurnal) (resp *schedule_service.Jurnal, err error) {

	i.log.Info("---UpdateJurnal------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Jurnal().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateJurnal--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Jurnal().GetByPKey(ctx, &schedule_service.JurnalPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetJurnal->Jurnal->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return resp, err
}

func (i *JurnalService) Delete(ctx context.Context, req *schedule_service.JurnalPrimaryKey) (resp *schedule_service.JurnalEmpty, err error) {

	i.log.Info("---DeleteJurnal------>", logger.Any("req", req))

	err = i.strg.Jurnal().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteJurnal->Jurnal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.JurnalEmpty{}, nil
}
