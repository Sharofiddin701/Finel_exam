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

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepoI {
	return &StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(ctx context.Context, req *user_service.CreateStudent) (resp *user_service.StudentPrimaryKey, err error) {

	var id = uuid.New().String()
	login := helper.GenerateLogin("S", "student", c.db)
	query := `INSERT INTO "student" (
				id,
				full_name,
				phone,
				password,
				login,
				group_id,
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
		req.Password,
		login,
		req.GroupId,
		req.BranchId,
		"15025cb6-f51a-40a2-aa65-c34e5b642379",
	)

	if err != nil {
		return nil, err
	}
	return &user_service.StudentPrimaryKey{Id: id}, nil
}

func (c *StudentRepo) GetByPKey(ctx context.Context, req *user_service.StudentPrimaryKey) (resp *user_service.Student, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			password,
			login,
			group_id,
			branch_id,
			role_id,
			created_at,
			updated_at
		FROM "student"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		full_name sql.NullString
		phone     sql.NullString
		password  sql.NullString
		login     sql.NullString
		group_id  sql.NullString
		branch_id sql.NullString
		role_id   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&password,
		&login,
		&group_id,
		&branch_id,
		&role_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Student{
		Id:        id.String,
		FullName:  full_name.String,
		Phone:     phone.String,
		Password:  password.String,
		Login:     login.String,
		GroupId:   group_id.String,
		BranchId:  branch_id.String,
		RoleId:    role_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *StudentRepo) GetAll(ctx context.Context, req *user_service.GetListStudentRequest) (resp *user_service.GetListStudentResponse, err error) {

	resp = &user_service.GetListStudentResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			full_name,
			phone,
			password,
			login,
			group_id,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "student"
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
		filter += " AND group_id ='" + req.GetSearch() + "'" + " OR branch_id ='" + req.GetSearch() + "'"
	}

	query += filter + offset + limit

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
			group_id  sql.NullString
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
			&group_id,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Students = append(resp.Students, &user_service.Student{
			Id:        id.String,
			FullName:  full_name.String,
			Phone:     phone.String,
			Password:  password.String,
			Login:     login.String,
			GroupId:   group_id.String,
			BranchId:  branch_id.String,
			RoleId:    role_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *StudentRepo) Update(ctx context.Context, req *user_service.UpdateStudent) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "student"
			SET
				full_name = :full_name,
				phone = :phone,
				password = :password,
				group_id = :group_id,
				branch_id =:branch_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":        req.GetId(),
		"full_name": req.GetFullName(),
		"phone":     req.GetPhone(),
		"password":  req.GetPassword(),
		"group_id":  req.GetGroupId(),
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

func (c *StudentRepo) Delete(ctx context.Context, req *user_service.StudentPrimaryKey) error {

	query := `DELETE FROM "student" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c *StudentRepo) GetStudetReport(ctx context.Context, req *user_service.StudentReportRequest) (resp *user_service.StudentReportResponse, err error) {

	resp = &user_service.StudentReportResponse{}

	query := `
				SELECT
					s.full_name,
					s.phone,
					p.paid_sum,
					p.course_count,
					p.total_sum
				FROM student AS s
				JOIN payment AS p ON s.id = p.student_id 
				WHERE s.branch_id = $1 


	`
	rows, err := c.db.Query(ctx, query, req.BranchId)
	defer rows.Close()
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		fmt.Println("ok1")

		var (
			full_name    sql.NullString
			phone        sql.NullString
			paid_sum     sql.NullFloat64
			course_count sql.NullInt64
			total_sum    sql.NullFloat64
		)

		err := rows.Scan(
			&full_name,
			&phone,
			&paid_sum,
			&course_count,
			&total_sum,
		)
		if err != nil {
			return resp, err
		}

		resp.Students = append(resp.Students, &user_service.StudentReport{
			FullName:    full_name.String,
			Phone:       phone.String,
			PaidSum:     paid_sum.Float64,
			CourseCount: course_count.Int64,
			TotalSum:    total_sum.Float64,
		})
	}
	return

}
