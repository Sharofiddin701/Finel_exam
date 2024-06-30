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

type ManagerRepo struct {
	db *pgxpool.Pool
}

func NewManagerRepo(db *pgxpool.Pool) storage.ManagerRepoI {
	return &ManagerRepo{
		db: db,
	}
}

func (c *ManagerRepo) Create(ctx context.Context, req *user_service.CreateManager) (resp *user_service.ManagerPrimaryKey, err error) {

	var id = uuid.New().String()
	login := helper.GenerateLogin("M", "manager", c.db)
	query := `INSERT INTO "manager" (
				id,
				full_name,
				phone,
				salary,
				password,
				login,
				branch_id,
				role_id,
				updated_at
			) VALUES ($1, $2,$3,$4,$5,$6,$7,$8, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.FullName,
		req.Phone,
		req.Salary,
		req.Password,
		login,
		req.BranchId,
		"45344572-d60f-43ae-9d03-4a16e0127e52",
	)

	if err != nil {
		return nil, err
	}
	return &user_service.ManagerPrimaryKey{Id: id}, nil
}

func (c *ManagerRepo) GetByPKey(ctx context.Context, req *user_service.ManagerPrimaryKey) (resp *user_service.Manager, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			salary,
			password,
			login,
			branch_id,
			role_id,
			created_at,
			updated_at
		FROM "manager"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		full_name sql.NullString
		phone     sql.NullString
		salary    sql.NullFloat64
		password  sql.NullString
		login     sql.NullString
		branch_id sql.NullString
		role_id   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&salary,
		&password,
		&login,
		&branch_id,
		&role_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Manager{
		Id:        id.String,
		FullName:  full_name.String,
		Phone:     phone.String,
		Salary:    salary.Float64,
		Password:  password.String,
		Login:     login.String,
		BranchId:  branch_id.String,
		RoleId:    role_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ManagerRepo) GetAll(ctx context.Context, req *user_service.GetListManagerRequest) (resp *user_service.GetListManagerResponse, err error) {

	resp = &user_service.GetListManagerResponse{}

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
			phone,
			salary,
			password,
			login,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "manager"
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
			phone     sql.NullString
			salary    sql.NullFloat64
			password  sql.NullString
			login     sql.NullString
			branch_id sql.NullString
			role_id   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&phone,
			&salary,
			&password,
			&login,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Managers = append(resp.Managers, &user_service.Manager{
			Id:        id.String,
			FullName:  full_name.String,
			Phone:     phone.String,
			Salary:    salary.Float64,
			Password:  password.String,
			Login:     login.String,
			BranchId:  branch_id.String,
			RoleId:    role_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ManagerRepo) Update(ctx context.Context, req *user_service.UpdateManager) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "manager"
			SET
				full_name = :full_name,
				phone = :phone,
				salary = :salary,
				password = :password,
				login = :login,
				branch_id =:branch_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":        req.GetId(),
		"full_name": req.GetFullName(),
		"phone":     req.GetPhone(),
		"salary":    req.GetSalary(),
		"password":  req.GetPassword(),
		"login":     req.GetLogin(),
		"branch_id": req.GetBranchId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *ManagerRepo) Delete(ctx context.Context, req *user_service.ManagerPrimaryKey) error {

	query := `DELETE FROM "manager" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
