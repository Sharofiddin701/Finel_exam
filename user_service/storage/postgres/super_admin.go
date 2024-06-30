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

type SuperAdminRepo struct {
	db *pgxpool.Pool
}

func NewSuperAdminRepo(db *pgxpool.Pool) storage.SuperAdminRepoI {
	return &SuperAdminRepo{
		db: db,
	}
}

func (c *SuperAdminRepo) Create(ctx context.Context, req *user_service.CreateSuperAdmin) (resp *user_service.SuperAdminPrimaryKey, err error) {

	var id = uuid.New().String()
	login := helper.GenerateLogin("SA", "super_admin", c.db)
	query := `INSERT INTO "super_admin" (
				id,
				full_name,
				password,	
				login,
				role_id,
				updated_at
			) VALUES ($1, $2, $3,$4,$5, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.FullName,
		req.Password,
		login,
		"badf0142-2dae-4f2f-bda1-09ee9aec7199",
	)

	if err != nil {
		return nil, err
	}
	return &user_service.SuperAdminPrimaryKey{Id: id}, nil
}

func (c *SuperAdminRepo) GetByPKey(ctx context.Context, req *user_service.SuperAdminPrimaryKey) (resp *user_service.SuperAdmin, err error) {

	query := `
		SELECT
			id,
			full_name,
			password,	
			login,
			role_id,
			created_at,
			updated_at
		FROM "super_admin"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		full_name sql.NullString
		password  sql.NullString
		login     sql.NullString
		role_id   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&password,
		&login,
		&role_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.SuperAdmin{
		Id:        id.String,
		FullName:  full_name.String,
		Password:  password.String,
		Login:     login.String,
		RoleId:    role_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *SuperAdminRepo) GetAll(ctx context.Context, req *user_service.GetListSuperAdminRequest) (resp *user_service.GetListSuperAdminResponse, err error) {

	resp = &user_service.GetListSuperAdminResponse{}

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
			full_name,	
			password,
			login,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "super_admin"
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
		filter += " AND full_name ILIKE '%" + req.GetSearch() + "%'"
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
			full_name sql.NullString
			password  sql.NullString
			login     sql.NullString
			role_id   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&password,
			&login,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.SuperAdmins = append(resp.SuperAdmins, &user_service.SuperAdmin{
			Id:        id.String,
			FullName:  full_name.String,
			Password:  password.String,
			Login:     login.String,
			RoleId:    role_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *SuperAdminRepo) Update(ctx context.Context, req *user_service.UpdateSuperAdmin) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "super_admin"
			SET
				full_name = :full_name,
				password = :password,
				login = :login,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":        req.GetId(),
		"full_name": req.GetFullName(),
		"password":  req.GetPassword(),
		"login":     req.GetLogin(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *SuperAdminRepo) Delete(ctx context.Context, req *user_service.SuperAdminPrimaryKey) error {

	query := `DELETE FROM "super_admin" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
