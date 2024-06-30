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

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepoI {
	return &TeacherRepo{
		db: db,
	}
}

func (c *TeacherRepo) Create(ctx context.Context, req *user_service.CreateTeacher) (resp *user_service.TeacherPrimaryKey, err error) {

	var id = uuid.New().String()
	login := helper.GenerateLogin("T", "teacher", c.db)
	query := `INSERT INTO "teacher" (
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
		"596839a1-45f1-44cc-aba3-fec60c82c2f3",
	)

	if err != nil {
		return nil, err
	}
	return &user_service.TeacherPrimaryKey{Id: id}, nil
}

func (c *TeacherRepo) GetByPKey(ctx context.Context, req *user_service.TeacherPrimaryKey) (resp *user_service.Teacher, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			salary,
			password,
			login,
			ielts_score,
			branch_id,
			role_id,
			created_at,
			updated_at
		FROM "teacher"
		WHERE id = $1
	`

	var (
		id          sql.NullString
		full_name   sql.NullString
		phone       sql.NullString
		salary      sql.NullFloat64
		password    sql.NullString
		login       sql.NullString
		ielts_score sql.NullString
		branch_id   sql.NullString
		role_id     sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&salary,
		&password,
		&login,
		&ielts_score,
		&branch_id,
		&role_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Teacher{
		Id:         id.String,
		FullName:   full_name.String,
		Phone:      phone.String,
		Salary:     salary.Float64,
		Password:   password.String,
		Login:      login.String,
		IeltsScore: ielts_score.String,
		BranchId:   branch_id.String,
		RoleId:     role_id.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}

	return
}

func (c *TeacherRepo) GetAll(ctx context.Context, req *user_service.GetListTeacherRequest) (resp *user_service.GetListTeacherResponse, err error) {

	resp = &user_service.GetListTeacherResponse{}

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
			ielts_score,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "teacher"
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
		filter += " AND branch_id ='" + req.GetSearch() + "'" + " OR full_name ILIKE '%" + req.GetSearch() + "%'"
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
			full_name   sql.NullString
			phone       sql.NullString
			password    sql.NullString
			login       sql.NullString
			salary      sql.NullFloat64
			ielts_score sql.NullString
			branch_id   sql.NullString
			role_id     sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&phone,
			&password,
			&login,
			&salary,
			&ielts_score,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Teachers = append(resp.Teachers, &user_service.Teacher{
			Id:         id.String,
			FullName:   full_name.String,
			Phone:      phone.String,
			Password:   password.String,
			Login:      login.String,
			Salary:     salary.Float64,
			IeltsScore: ielts_score.String,
			BranchId:   branch_id.String,
			RoleId:     role_id.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return
}

func (c *TeacherRepo) Update(ctx context.Context, req *user_service.UpdateTeacher) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "teacher"
			SET
				full_name = :full_name,
				phone = :phone,
				password = :password,
				login = :login,
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
		"login":       req.GetLogin(),
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

func (c *TeacherRepo) Delete(ctx context.Context, req *user_service.TeacherPrimaryKey) error {

	query := `DELETE FROM "teacher" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
