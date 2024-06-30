package storage

import (
	"context"
	"user_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	SuperAdmin() SuperAdminRepoI
	Branch() BranchRepoI
	Manager() ManagerRepoI
	Teacher() TeacherRepoI
	SupportTeacher() SupportTeacherRepoI
	Administrator() AdministratorRepoI
	Student() StudentRepoI
	Payment() PaymentRepoI
}

type SuperAdminRepoI interface {
	Create(ctx context.Context, req *user_service.CreateSuperAdmin) (resp *user_service.SuperAdminPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.SuperAdminPrimaryKey) (resp *user_service.SuperAdmin, err error)
	GetAll(ctx context.Context, req *user_service.GetListSuperAdminRequest) (resp *user_service.GetListSuperAdminResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateSuperAdmin) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.SuperAdminPrimaryKey) error
}

type BranchRepoI interface {
	Create(ctx context.Context, req *user_service.CreateBranch) (resp *user_service.BranchPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.BranchPrimaryKey) (resp *user_service.Branch, err error)
	GetAll(ctx context.Context, req *user_service.GetListBranchRequest) (resp *user_service.GetListBranchResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateBranch) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.BranchPrimaryKey) error
}

type ManagerRepoI interface {
	Create(ctx context.Context, req *user_service.CreateManager) (resp *user_service.ManagerPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.ManagerPrimaryKey) (resp *user_service.Manager, err error)
	GetAll(ctx context.Context, req *user_service.GetListManagerRequest) (resp *user_service.GetListManagerResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateManager) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.ManagerPrimaryKey) error
}

type TeacherRepoI interface {
	Create(ctx context.Context, req *user_service.CreateTeacher) (resp *user_service.TeacherPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.TeacherPrimaryKey) (resp *user_service.Teacher, err error)
	GetAll(ctx context.Context, req *user_service.GetListTeacherRequest) (resp *user_service.GetListTeacherResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateTeacher) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.TeacherPrimaryKey) error
	// TeacherPanel(ctx context.Context, req *user_service.TeacherPanelRequest) (resp *user_service.TeacherPanelResponse, err error)
}

type SupportTeacherRepoI interface {
	Create(ctx context.Context, req *user_service.CreateSupportTeacher) (resp *user_service.SupportTeacherPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) (resp *user_service.SupportTeacher, err error)
	GetAll(ctx context.Context, req *user_service.GetListSupportTeacherRequest) (resp *user_service.GetListSupportTeacherResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateSupportTeacher) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) error
}

type AdministratorRepoI interface {
	Create(ctx context.Context, req *user_service.CreateAdministrator) (resp *user_service.AdministratorPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.AdministratorPrimaryKey) (resp *user_service.Administrator, err error)
	GetAll(ctx context.Context, req *user_service.GetListAdministratorRequest) (resp *user_service.GetListAdministratorResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateAdministrator) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.AdministratorPrimaryKey) error
}

type StudentRepoI interface {
	Create(ctx context.Context, req *user_service.CreateStudent) (resp *user_service.StudentPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.StudentPrimaryKey) (resp *user_service.Student, err error)
	GetAll(ctx context.Context, req *user_service.GetListStudentRequest) (resp *user_service.GetListStudentResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateStudent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.StudentPrimaryKey) error
	GetStudetReport(ctx context.Context, req *user_service.StudentReportRequest) (resp *user_service.StudentReportResponse, err error)
}

type PaymentRepoI interface {
	Create(ctx context.Context, req *user_service.CreatePayment) (resp *user_service.PaymentPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.PaymentPrimaryKey) (resp *user_service.Payment, err error)
	GetAll(ctx context.Context, req *user_service.GetListPaymentRequest) (resp *user_service.GetListPaymentResponse, err error)
	Update(ctx context.Context, req *user_service.UpdatePayment) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.PaymentPrimaryKey) error
}
