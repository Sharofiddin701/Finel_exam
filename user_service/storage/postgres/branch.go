package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &BranchRepo{
		db: db,
	}
}

func (c *BranchRepo) Create(ctx context.Context, req *user_service.CreateBranch) (resp *user_service.BranchPrimaryKey, err error) {

	var id = uuid.New().String()

	query := `INSERT INTO "branch" (
				id,
				name,
				updated_at
			) VALUES ($1, $2, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.Name,
	)

	if err != nil {
		return nil, err
	}
	return &user_service.BranchPrimaryKey{Id: id}, nil
}

func (c *BranchRepo) GetByPKey(ctx context.Context, req *user_service.BranchPrimaryKey) (resp *user_service.Branch, err error) {

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "branch"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Branch{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *BranchRepo) GetAll(ctx context.Context, req *user_service.GetListBranchRequest) (resp *user_service.GetListBranchResponse, err error) {

	resp = &user_service.GetListBranchResponse{}

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
			name,	
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "branch"
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
		filter += " AND name ILIKE '%" + req.GetSearch() + "%'"
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
			name      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Branches = append(resp.Branches, &user_service.Branch{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *BranchRepo) Update(ctx context.Context, req *user_service.UpdateBranch) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "branch"
			SET
				name = :name,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"name": req.GetName(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *BranchRepo) Delete(ctx context.Context, req *user_service.BranchPrimaryKey) error {

	query := `DELETE FROM "branch" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
