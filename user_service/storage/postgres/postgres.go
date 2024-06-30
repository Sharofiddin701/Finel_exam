package postgres

import (
	"context"
	"fmt"
	"log"
	"user_service/config"
	"user_service/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db             *pgxpool.Pool
	superadmin     storage.SuperAdminRepoI
	branch         storage.BranchRepoI
	manager        storage.ManagerRepoI
	teacher        storage.TeacherRepoI
	suppurtteacher storage.SupportTeacherRepoI
	adminstrator   storage.AdministratorRepoI
	student        storage.StudentRepoI
	payment        storage.PaymentRepoI
}

type Pool struct {
	db *pgxpool.Pool
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) SuperAdmin() storage.SuperAdminRepoI {
	if s.superadmin == nil {
		s.superadmin = NewSuperAdminRepo(s.db)
	}

	return s.superadmin
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}
func (s *Store) Manager() storage.ManagerRepoI {
	if s.manager == nil {
		s.manager = NewManagerRepo(s.db)
	}

	return s.manager
}

func (s *Store) Teacher() storage.TeacherRepoI {
	if s.teacher == nil {
		s.teacher = NewTeacherRepo(s.db)
	}

	return s.teacher
}
func (s *Store) SupportTeacher() storage.SupportTeacherRepoI {
	if s.suppurtteacher == nil {
		s.suppurtteacher = NewSupportTeacherRepo(s.db)
	}

	return s.suppurtteacher
}

func (s *Store) Administrator() storage.AdministratorRepoI {
	if s.adminstrator == nil {
		s.adminstrator = NewAdministratorRepo(s.db)
	}

	return s.adminstrator
}

func (s *Store) Student() storage.StudentRepoI {
	if s.student == nil {
		s.student = NewStudentRepo(s.db)
	}

	return s.student
}

func (s *Store) Payment() storage.PaymentRepoI {
	if s.payment == nil {
		s.payment = NewPaymentRepo(s.db)
	}

	return s.payment
}
