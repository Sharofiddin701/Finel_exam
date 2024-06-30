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

type PaymentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedPaymentServiceServer
}

func NewPaymentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *PaymentService {
	return &PaymentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *PaymentService) Create(ctx context.Context, req *user_service.CreatePayment) (resp *user_service.Payment, err error) {

	i.log.Info("---CreatePayment------>", logger.Any("req", req))

	pKey, err := i.strg.Payment().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreatePayment->Payment->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Payment().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyPayment->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *PaymentService) GetByID(ctx context.Context, req *user_service.PaymentPrimaryKey) (resp *user_service.Payment, err error) {

	i.log.Info("---GetPaymentByID------>", logger.Any("req", req))

	resp, err = i.strg.Payment().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPaymentByID->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *PaymentService) GetList(ctx context.Context, req *user_service.GetListPaymentRequest) (resp *user_service.GetListPaymentResponse, err error) {

	i.log.Info("---GetPayments------>", logger.Any("req", req))

	resp, err = i.strg.Payment().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPayments->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *PaymentService) Update(ctx context.Context, req *user_service.UpdatePayment) (resp *user_service.Payment, err error) {

	i.log.Info("---UpdatePayment------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Payment().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdatePayment--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Payment().GetByPKey(ctx, &user_service.PaymentPrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetPayment->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *PaymentService) Delete(ctx context.Context, req *user_service.PaymentPrimaryKey) (resp *user_service.PaymentEmpty, err error) {

	i.log.Info("---DeletePayment------>", logger.Any("req", req))

	err = i.strg.Payment().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeletePayment->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_service.PaymentEmpty{}, nil
}
