package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"schedule_service/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) storage.EventRepoI {
	return &EventRepo{
		db: db,
	}
}

func (c *EventRepo) Create(ctx context.Context, req *schedule_service.CreateEvent) (resp *schedule_service.EventPrimaryKey, err error) {

	var id = uuid.New().String()
	query := `INSERT INTO "event" (
				id,
				topic,
				date,
				start_time,
				branch_id,
				updated_at
			) VALUES ($1, $2,$3,$4,$5, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.Topic,
		req.Date,
		req.StartTime,
		req.BranchId,
	)

	if err != nil {
		return nil, err
	}
	return &schedule_service.EventPrimaryKey{Id: id}, nil
}

func (c *EventRepo) GetByPKey(ctx context.Context, req *schedule_service.EventPrimaryKey) (resp *schedule_service.Event, err error) {

	query := `
		SELECT
			id,
			topic,
			date,
			start_time,
			branch_id,
			created_at,
			updated_at
		FROM "event"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		topic      sql.NullString
		date       sql.NullString
		start_time sql.NullString
		branch_id  sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&topic,
		&date,
		&start_time,
		&branch_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Event{
		Id:        id.String,
		Topic:     topic.String,
		Date:      date.String,
		StartTime: start_time.String,
		BranchId:  branch_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *EventRepo) GetAll(ctx context.Context, req *schedule_service.GetListEventRequest) (resp *schedule_service.GetListEventResponse, err error) {

	resp = &schedule_service.GetListEventResponse{}

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
			topic,
			date,
			start_time,
			branch_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "event"
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
			id         sql.NullString
			topic      sql.NullString
			date       sql.NullString
			start_time sql.NullString
			branch_id  sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&topic,
			&date,
			&start_time,
			&branch_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Events = append(resp.Events, &schedule_service.Event{
			Id:        id.String,
			Topic:     topic.String,
			Date:      date.String,
			StartTime: start_time.String,
			BranchId:  branch_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *EventRepo) Update(ctx context.Context, req *schedule_service.UpdateEvent) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "event"
			SET
				topic = :topic,
				date = :date,
				start_time = :start_time,
				branch_id =:branch_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"topic":      req.GetTopic(),
		"date":       req.GetDate(),
		"start_time": req.GetStartTime(),
		"branch_id":  req.GetBranchId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *EventRepo) Delete(ctx context.Context, req *schedule_service.EventPrimaryKey) error {

	query := `DELETE FROM "event" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
