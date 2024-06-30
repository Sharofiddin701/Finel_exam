package postgres

import (
	"context"
	"schedule_service/storage"
	"fmt"
	"log"
	"schedule_service/config"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db     *pgxpool.Pool
	group  storage.GroupRepoI
	dotask storage.DoTaskRepoI
	event  storage.EventRepoI

	assign_student storage.AssignStudentRepoI
	score          storage.ScoreRepoI
	journal        storage.JurnalRepoI
	task           storage.TaskRepoI
	schedule       storage.ScheduleRepoI
	lesson         storage.LessonRepoI
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

func (s *Store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}

	return s.group
}

func (s *Store) Event() storage.EventRepoI {
	if s.event == nil {
		s.event = NewEventRepo(s.db)
	}

	return s.event
}

func (s *Store) AssignStudent() storage.AssignStudentRepoI {
	if s.assign_student == nil {
		s.assign_student = NewAssignStudentRepo(s.db)
	}

	return s.assign_student
}
func (s *Store) DoTask() storage.DoTaskRepoI {
	if s.dotask == nil {
		s.dotask = NewDoTaskRepo(s.db)
	}

	return s.dotask
}

func (s *Store) Score() storage.ScoreRepoI {
	if s.score == nil {
		s.score = NewScoreRepo(s.db)
	}

	return s.score
}

func (s *Store) Jurnal() storage.JurnalRepoI {
	if s.journal == nil {
		s.journal = NewJurnalRepo(s.db)
	}

	return s.journal
}
func (s *Store) Task() storage.TaskRepoI {
	if s.task == nil {
		s.task = NewTaskRepo(s.db)
	}

	return s.task
}
func (s *Store) Schedule() storage.ScheduleRepoI {
	if s.schedule == nil {
		s.schedule = NewScheduleRepo(s.db)
	}

	return s.schedule
}

func (s *Store) Lesson() storage.LessonRepoI {
	if s.lesson == nil {
		s.lesson = NewLessonRepo(s.db)
	}

	return s.lesson
}
