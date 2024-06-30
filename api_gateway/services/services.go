package services

import (
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CrmService() user_service.SuperAdminServiceClient
	Branch() user_service.BranchServiceClient
	Manager() user_service.ManagerServiceClient
	Teacher() user_service.TeacherServiceClient
	SupportTeacher() user_service.SupportTeacherServiceClient
	Administrator() user_service.AdministratorServiceClient
	Student() user_service.StudentServiceClient
	Group() schedule_service.GroupServiceClient
	Event() schedule_service.EventServiceClient
	AssignStudent() schedule_service.AssignStudentServiceClient
	DoTask() schedule_service.DoTaskServiceClient
	Score() schedule_service.ScoreServiceClient
	Jurnal() schedule_service.JurnalServiceClient
	Task() schedule_service.TaskServiceClient
	Schedule() schedule_service.ScheduleServiceClient
	Lesson() schedule_service.LessonServiceClient
	Payment() user_service.PaymentServiceClient
}

type grpcClients struct {
	crmService     user_service.SuperAdminServiceClient
	branch         user_service.BranchServiceClient
	manager        user_service.ManagerServiceClient
	teacher        user_service.TeacherServiceClient
	supportteacher user_service.SupportTeacherServiceClient
	adminstrator   user_service.AdministratorServiceClient
	student        user_service.StudentServiceClient
	group          schedule_service.GroupServiceClient
	event          schedule_service.EventServiceClient
	assignstudent  schedule_service.AssignStudentServiceClient
	dotask         schedule_service.DoTaskServiceClient
	score          schedule_service.ScoreServiceClient
	jurnal         schedule_service.JurnalServiceClient
	task           schedule_service.TaskServiceClient
	schedule       schedule_service.ScheduleServiceClient
	lesson         schedule_service.LessonServiceClient
	payment        user_service.PaymentServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// User Service...

	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	connScheduleService, err := grpc.Dial(
		cfg.ScheduleServiceHost+cfg.ScheduleServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		crmService:     user_service.NewSuperAdminServiceClient(connUserService),
		branch:         user_service.NewBranchServiceClient(connUserService),
		manager:        user_service.NewManagerServiceClient(connUserService),
		teacher:        user_service.NewTeacherServiceClient(connUserService),
		supportteacher: user_service.NewSupportTeacherServiceClient(connUserService),
		adminstrator:   user_service.NewAdministratorServiceClient(connUserService),
		student:        user_service.NewStudentServiceClient(connUserService),
		group:          schedule_service.NewGroupServiceClient(connScheduleService),
		event:          schedule_service.NewEventServiceClient(connScheduleService),
		assignstudent:  schedule_service.NewAssignStudentServiceClient(connScheduleService),
		dotask:         schedule_service.NewDoTaskServiceClient(connScheduleService),
		score:          schedule_service.NewScoreServiceClient(connScheduleService),
		jurnal:         schedule_service.NewJurnalServiceClient(connScheduleService),
		task:           schedule_service.NewTaskServiceClient(connScheduleService),
		schedule:       schedule_service.NewScheduleServiceClient(connScheduleService),
		lesson:         schedule_service.NewLessonServiceClient(connScheduleService),
		payment:        user_service.NewPaymentServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) CrmService() user_service.SuperAdminServiceClient {
	return g.crmService
}
func (g *grpcClients) Branch() user_service.BranchServiceClient {
	return g.branch
}
func (g *grpcClients) Manager() user_service.ManagerServiceClient {
	return g.manager
}
func (g *grpcClients) Teacher() user_service.TeacherServiceClient {
	return g.teacher
}
func (g *grpcClients) SupportTeacher() user_service.SupportTeacherServiceClient {
	return g.supportteacher
}
func (g *grpcClients) Administrator() user_service.AdministratorServiceClient {
	return g.adminstrator
}

func (g *grpcClients) Group() schedule_service.GroupServiceClient {
	return g.group
}
func (g *grpcClients) Student() user_service.StudentServiceClient {
	return g.student
}
func (g *grpcClients) Event() schedule_service.EventServiceClient {
	return g.event
}
func (g *grpcClients) AssignStudent() schedule_service.AssignStudentServiceClient {
	return g.assignstudent
}
func (g *grpcClients) DoTask() schedule_service.DoTaskServiceClient {
	return g.dotask
}
func (g *grpcClients) Score() schedule_service.ScoreServiceClient {
	return g.score
}
func (g *grpcClients) Jurnal() schedule_service.JurnalServiceClient {
	return g.jurnal
}
func (g *grpcClients) Task() schedule_service.TaskServiceClient {
	return g.task
}
func (g *grpcClients) Schedule() schedule_service.ScheduleServiceClient {
	return g.schedule
}
func (g *grpcClients) Lesson() schedule_service.LessonServiceClient {
	return g.lesson
}
func (g *grpcClients) Payment() user_service.PaymentServiceClient {
	return g.payment
}

// func (g *grpcClients) ProviderService() user_service.ProviderServiceClient {
// 	return g.provider
// }
