package postgres

import (
	"context"
	"database/sql"
	"fmt"
	schedule_service "schedule_service/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JurnalRepo struct {
	db *pgxpool.Pool
}

func NewJurnalRepo(db *pgxpool.Pool) storage.JurnalRepoI {
	return &JurnalRepo{
		db: db,
	}
}

func (c *JurnalRepo) Create(ctx context.Context, req *schedule_service.CreateJurnal) (resp *schedule_service.JurnalPrimaryKey, err error) {

	var id = uuid.New().String()
	to_datet := time.Now().AddDate(0, 3, 1)
	to_date := to_datet.Format("2006-01-02")

	query := `INSERT INTO "journal" (
				id,
				group_id,
				from_date,
				to_date,
				updated_at
			) VALUES ($1, $2,$3,$4, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GroupId,
		req.From,
		to_date,
	)

	if err != nil {
		return nil, err
	}

	return &schedule_service.JurnalPrimaryKey{Id: id}, nil
}

func (c *JurnalRepo) GetByPKey(ctx context.Context, req *schedule_service.JurnalPrimaryKey) (resp *schedule_service.Jurnal, err error) {

	query := `
		SELECT
			id,
			group_id,
			from_date,
			to_date,
			created_at,
			updated_at
		FROM "journal"
		WHERE id = $1 OR group_id = $1
	`

	var (
		id        sql.NullString
		group_id  sql.NullString
		from_date sql.NullString
		to_date   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&group_id,
		&from_date,
		&to_date,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Jurnal{
		Id:        id.String,
		GroupId:   group_id.String,
		From:      from_date.String,
		To:        to_date.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *JurnalRepo) GetAll(ctx context.Context, req *schedule_service.GetListJurnalRequest) (resp *schedule_service.GetListJurnalResponse, err error) {

	resp = &schedule_service.GetListJurnalResponse{}

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
			group_id,
			from_date,
			to_date,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "journal"
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
			id        sql.NullString
			group_id  sql.NullString
			from_date sql.NullString
			to_date   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&group_id,
			&from_date,
			&to_date,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Jurnals = append(resp.Jurnals, &schedule_service.Jurnal{
			Id:        id.String,
			GroupId:   group_id.String,
			From:      from_date.String,
			To:        to_date.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *JurnalRepo) Update(ctx context.Context, req *schedule_service.UpdateJurnal) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "journal"
			SET
				group_id = :group_id,
				from = :from_date,
				to = :to_date,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":       req.GetId(),
		"group_id": req.GetGroupId(),
		"from":     req.GetFrom(),
		"to":       req.GetTo(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *JurnalRepo) Delete(ctx context.Context, req *schedule_service.JurnalPrimaryKey) error {

	query := `DELETE FROM "journal" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
