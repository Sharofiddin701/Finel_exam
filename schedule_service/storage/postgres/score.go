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

type ScoreRepo struct {
	db *pgxpool.Pool
}

func NewScoreRepo(db *pgxpool.Pool) storage.ScoreRepoI {
	return &ScoreRepo{
		db: db,
	}
}

func (c *ScoreRepo) Create(ctx context.Context, req *schedule_service.CreateScore) (resp *schedule_service.ScorePrimaryKey, err error) {

	var id = uuid.New().String()

	query := `INSERT INTO "score" (
				id,
				task_id,
				student_id,
				score,
				updated_at
			) VALUES ($1, $2, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.TaskId,
		req.StudentId,
		req.Score,
	)

	if err != nil {
		return nil, err
	}
	return &schedule_service.ScorePrimaryKey{Id: id}, nil
}

func (c *ScoreRepo) GetByPKey(ctx context.Context, req *schedule_service.ScorePrimaryKey) (resp *schedule_service.Score, err error) {

	query := `
		SELECT
			id,
			task_id,
			student_id,
			score,
			created_at,
			updated_at
		FROM "score"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		task_id    sql.NullString
		student_id sql.NullString
		score      sql.NullInt64
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&task_id,
		&student_id,
		&score,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Score{
		Id:        id.String,
		TaskId:    task_id.String,
		StudentId: student_id.String,
		Score:     score.Int64,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ScoreRepo) GetAll(ctx context.Context, req *schedule_service.GetListScoreRequest) (resp *schedule_service.GetListScoreResponse, err error) {

	resp = &schedule_service.GetListScoreResponse{}

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
			task_id,
			student_id,
			score,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "score"
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
		filter += " AND student_id ='" + req.GetSearch() + "'"
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
			id         sql.NullString
			task_id    sql.NullString
			student_id sql.NullString
			score      sql.NullInt64
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&task_id,
			&student_id,
			&score,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Scores = append(resp.Scores, &schedule_service.Score{
			Id:        id.String,
			TaskId:    task_id.String,
			StudentId: student_id.String,
			Score:     score.Int64,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ScoreRepo) Update(ctx context.Context, req *schedule_service.UpdateScore) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "score"
			SET
				task_id = :task_id,
				student_id = :student_id,
				score = :score,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"task_id":    req.GetTaskId(),
		"student_id": req.GetStudentId(),
		"score":      req.GetScore(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *ScoreRepo) Delete(ctx context.Context, req *schedule_service.ScorePrimaryKey) error {

	query := `DELETE FROM "score" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
