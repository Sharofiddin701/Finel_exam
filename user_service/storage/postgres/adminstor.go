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

type AdministratorRepo struct {
	db *pgxpool.Pool
}

func NewAdministratorRepo(db *pgxpool.Pool) storage.AdministratorRepoI {
	return &AdministratorRepo{
		db: db,
	}
}

func (c *AdministratorRepo) Create(ctx context.Context, req *user_service.CreateAdministrator) (resp *user_service.AdministratorPrimaryKey, err error) {

	var id = uuid.New().String()
	login := helper.GenerateLogin("A", "administrator", c.db)
	query := `INSERT INTO "administrator" (
				id,
				full_name,
				phone,
				salary,
				password,
				login,
				ielts_score,
				branch_id,
				role_id,
				updated_at
			) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.FullName,
		req.Phone,
		req.Salary,
		req.Password,
		login,
		req.IeltsScore,
		req.BranchId,
		"c6b9cac8-ecf1-4b99-b8a9-571daed55fb5",
	)

	if err != nil {
		return nil, err
	}
	return &user_service.AdministratorPrimaryKey{Id: id}, nil
}

func (c *AdministratorRepo) GetByPKey(ctx context.Context, req *user_service.AdministratorPrimaryKey) (resp *user_service.Administrator, err error) {

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
		FROM "administrator"
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

	resp = &user_service.Administrator{
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

func (c *AdministratorRepo) GetAll(ctx context.Context, req *user_service.GetListAdministratorRequest) (resp *user_service.GetListAdministratorResponse, err error) {

	resp = &user_service.GetListAdministratorResponse{}

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
			password,
			login,
			salary,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "administrator"
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
		filter += " AND full_name ILIKE '%" + req.GetSearch() + "%'" + " OR branch_id ='" + req.GetSearch() + "'"
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
			password  sql.NullString
			login     sql.NullString
			salary    sql.NullFloat64
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
			&password,
			&login,
			&salary,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Administrators = append(resp.Administrators, &user_service.Administrator{
			Id:        id.String,
			FullName:  full_name.String,
			Phone:     phone.String,
			Password:  password.String,
			Login:     login.String,
			Salary:    salary.Float64,
			BranchId:  branch_id.String,
			RoleId:    role_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *AdministratorRepo) Update(ctx context.Context, req *user_service.UpdateAdministrator) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "administrator"
			SET
				full_name = :full_name,
				phone = :phone,
				password = :password,
				salary = :salary,
				ielts_score = :ielts_score,
				branch_id =:branch_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":          req.GetId(),
		"full_name":   req.GetFullName(),
		"phone":       req.GetPhone(),
		"password":    req.GetPassword(),
		"salary":      req.GetSalary(),
		"ielts_score": req.GetIeltsScore(),
		"branch_id":   req.GetBranchId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *AdministratorRepo) Delete(ctx context.Context, req *user_service.AdministratorPrimaryKey) error {

	query := `DELETE FROM "administrator" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
