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

type ScoreService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*schedule_service.UnimplementedScoreServiceServer
}

func NewScoreService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScoreService {
	return &ScoreService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ScoreService) Create(ctx context.Context, req *schedule_service.CreateScore) (resp *schedule_service.Score, err error) {

	i.log.Info("---CreateScore------>", logger.Any("req", req))

	pKey, err := i.strg.Score().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateScore->Score->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Score().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyScore->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScoreService) GetByID(ctx context.Context, req *schedule_service.ScorePrimaryKey) (resp *schedule_service.Score, err error) {

	i.log.Info("---GetScoreByID------>", logger.Any("req", req))

	resp, err = i.strg.Score().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScoreByID->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScoreService) GetList(ctx context.Context, req *schedule_service.GetListScoreRequest) (resp *schedule_service.GetListScoreResponse, err error) {

	i.log.Info("---GetScores------>", logger.Any("req", req))

	resp, err = i.strg.Score().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScores->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScoreService) Update(ctx context.Context, req *schedule_service.UpdateScore) (resp *schedule_service.Score, err error) {

	i.log.Info("---UpdateScore------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Score().Update(ctx, req)
	if err != nil {
		i.log.Info("!!!UpdateScore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// fmt.Println("ok1")
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	// fmt.Println("ok2")

	resp, err = i.strg.Score().GetByPKey(ctx, &schedule_service.ScorePrimaryKey{Id: req.Id})
	// fmt.Println("ok3")

	if err != nil {
		i.log.Error("!!!GetScore->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ScoreService) Delete(ctx context.Context, req *schedule_service.ScorePrimaryKey) (resp *schedule_service.ScoreEmpty, err error) {

	i.log.Info("---DeleteScore------>", logger.Any("req", req))

	err = i.strg.Score().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteScore->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &schedule_service.ScoreEmpty{}, nil
}
