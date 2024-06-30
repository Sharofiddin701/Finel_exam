package storage

import (
	"context"
	"schedule_service/genproto/schedule_service"
)

type StorageI interface {
	CloseDB()

	Group() GroupRepoI
	Event() EventRepoI
	AssignStudent() AssignStudentRepoI
	DoTask() DoTaskRepoI
	Score() ScoreRepoI
	Jurnal() JurnalRepoI
	Task() TaskRepoI
	Schedule() ScheduleRepoI
	Lesson() LessonRepoI
}

type GroupRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateGroup) (resp *schedule_service.GroupPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.GroupPrimaryKey) (resp *schedule_service.Group, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListGroupRequest) (resp *schedule_service.GetListGroupResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateGroup) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.GroupPrimaryKey) error
}

type DoTaskRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateDoTask) (resp *schedule_service.DoTaskPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.DoTaskPrimaryKey) (resp *schedule_service.DoTask, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListDoTaskRequest) (resp *schedule_service.GetListDoTaskResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateDoTask) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.DoTaskPrimaryKey) error
}

type EventRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateEvent) (resp *schedule_service.EventPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.EventPrimaryKey) (resp *schedule_service.Event, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListEventRequest) (resp *schedule_service.GetListEventResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateEvent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.EventPrimaryKey) error
}

type ScoreRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateScore) (resp *schedule_service.ScorePrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.ScorePrimaryKey) (resp *schedule_service.Score, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListScoreRequest) (resp *schedule_service.GetListScoreResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateScore) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.ScorePrimaryKey) error
}

type AssignStudentRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateAssignStudent) (resp *schedule_service.AssignStudentPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.AssignStudentPrimaryKey) (resp *schedule_service.AssignStudent, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListAssignStudentRequest) (resp *schedule_service.GetListAssignStudentResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateAssignStudent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.AssignStudentPrimaryKey) error
}

type JurnalRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateJurnal) (resp *schedule_service.JurnalPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.JurnalPrimaryKey) (resp *schedule_service.Jurnal, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListJurnalRequest) (resp *schedule_service.GetListJurnalResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateJurnal) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.JurnalPrimaryKey) error
}
type TaskRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateTask) (resp *schedule_service.TaskPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.TaskPrimaryKey) (resp *schedule_service.Task, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListTaskRequest) (resp *schedule_service.GetListTaskResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateTask) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.TaskPrimaryKey) error
}

type ScheduleRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateSchedule) (resp *schedule_service.SchedulePrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (resp *schedule_service.Schedule, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListScheduleRequest) (resp *schedule_service.GetListScheduleResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateSchedule) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.SchedulePrimaryKey) error
	GetScheduleReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error)
	GetScheduleMonthReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error)
}

type LessonRepoI interface {
	Create(ctx context.Context, req *schedule_service.CreateLesson) (resp *schedule_service.LessonPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *schedule_service.LessonPrimaryKey) (resp *schedule_service.Lesson, err error)
	GetAll(ctx context.Context, req *schedule_service.GetListLessonRequest) (resp *schedule_service.GetListLessonResponse, err error)
	Update(ctx context.Context, req *schedule_service.UpdateLesson) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *schedule_service.LessonPrimaryKey) error
}
