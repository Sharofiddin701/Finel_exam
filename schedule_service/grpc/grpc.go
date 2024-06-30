package grpc

import (
	"schedule_service/config"
	"schedule_service/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/grpc/service"
	"schedule_service/pkg/logger"
	"schedule_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	schedule_service.RegisterGroupServiceServer(grpcServer, service.NewGroupService(cfg, log, strg, srvc))
	schedule_service.RegisterEventServiceServer(grpcServer, service.NewEventService(cfg, log, strg, srvc))
	schedule_service.RegisterAssignStudentServiceServer(grpcServer, service.NewAssignStudentService(cfg, log, strg, srvc))
	schedule_service.RegisterDoTaskServiceServer(grpcServer, service.NewDoTaskService(cfg, log, strg, srvc))
	schedule_service.RegisterScoreServiceServer(grpcServer, service.NewScoreService(cfg, log, strg, srvc))
	schedule_service.RegisterJurnalServiceServer(grpcServer, service.NewJurnalService(cfg, log, strg, srvc))
	schedule_service.RegisterTaskServiceServer(grpcServer, service.NewTaskService(cfg, log, strg, srvc))
	schedule_service.RegisterScheduleServiceServer(grpcServer, service.NewScheduleService(cfg, log, strg, srvc))
	schedule_service.RegisterLessonServiceServer(grpcServer, service.NewLessonService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
