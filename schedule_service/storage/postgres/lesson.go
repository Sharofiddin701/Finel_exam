package postgres

import (
	"context"
	"database/sql"
	"fmt"
	schedule_service "schedule_service/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LessonRepo struct {
	db *pgxpool.Pool
}

func NewLessonRepo(db *pgxpool.Pool) storage.LessonRepoI {
	return &LessonRepo{
		db: db,
	}
}

func (c *LessonRepo) Create(ctx context.Context, req *schedule_service.CreateLesson) (resp *schedule_service.LessonPrimaryKey, err error) {

	var id = uuid.New().String()
	query := `INSERT INTO "lesson" (
				id,
				schedule_id,
				lesson,
				updated_at
			) VALUES ($1, $2,$3, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.ScheduleId,
		req.Lesson,
	)

	if err != nil {
		return nil, err
	}
	return &schedule_service.LessonPrimaryKey{Id: id}, nil
}

func (c *LessonRepo) GetByPKey(ctx context.Context, req *schedule_service.LessonPrimaryKey) (resp *schedule_service.Lesson, err error) {

	query := `
		SELECT
			id,
			schedule_id,
			lesson,
			created_at,
			updated_at
		FROM "lesson"
		WHERE id = $1 OR schedule_id = $1
	`

	var (
		id          sql.NullString
		schedule_id sql.NullString
		lesson      sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&schedule_id,
		&lesson,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Lesson{
		Id:         id.String,
		ScheduleId: schedule_id.String,
		Lesson:     lesson.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}

	return
}

func (c *LessonRepo) GetAll(ctx context.Context, req *schedule_service.GetListLessonRequest) (resp *schedule_service.GetListLessonResponse, err error) {

	resp = &schedule_service.GetListLessonResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE"
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			schedule_id,
			lesson,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "lesson"
	`

	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}

	if len(req.GetSearch()) > 0 {
		filter += " AND topic ILIKE '%" + req.GetSearch() + "%'"
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)
	defer rows.Close()

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var (
			id          sql.NullString
			schedule_id sql.NullString
			lesson      sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&schedule_id,
			&lesson,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Lessons = append(resp.Lessons, &schedule_service.Lesson{
			Id:         id.String,
			ScheduleId: schedule_id.String,
			Lesson:     lesson.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return
}

func (c *LessonRepo) Update(ctx context.Context, req *schedule_service.UpdateLesson) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "lesson"
			SET
				schedule_id = :schedule_id,
				lesson = :lesson,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":          req.Id,
		"schedule_id": req.ScheduleId,
		"lesson":      req.Lesson,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *LessonRepo) Delete(ctx context.Context, req *schedule_service.LessonPrimaryKey) error {

	query := `DELETE FROM "lesson" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
