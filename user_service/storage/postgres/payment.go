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

type PaymentRepo struct {
	db *pgxpool.Pool
}

func NewPaymentRepo(db *pgxpool.Pool) storage.PaymentRepoI {
	return &PaymentRepo{
		db: db,
	}
}

func (c *PaymentRepo) Create(ctx context.Context, req *user_service.CreatePayment) (resp *user_service.PaymentPrimaryKey, err error) {

	var id = uuid.New().String()
	query := `INSERT INTO "payment" (
				id,
				student_id,
				branch_id,
				paid_sum,
				total_sum,
				course_count,
				updated_at
			) VALUES ($1, $2,$3,$4,$5,$6, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.StudentId,
		req.BranchId,
		req.PaidSum,
		float64(req.CourseCount)*300000,
		req.CourseCount,
	)

	if err != nil {
		return nil, err
	}
	return &user_service.PaymentPrimaryKey{Id: id}, nil
}

func (c *PaymentRepo) GetByPKey(ctx context.Context, req *user_service.PaymentPrimaryKey) (resp *user_service.Payment, err error) {

	query := `
		SELECT
			id,
			student_id,
			branch_id,
			paid_sum,
			total_sum,
			course_count,	
			created_at,
			updated_at
		FROM payment
		WHERE id = $1 OR student_id = $1
	`

	var (
		id           sql.NullString
		student_id   sql.NullString
		branch_id    sql.NullString
		paid_sum     sql.NullFloat64
		total_sum    sql.NullFloat64
		course_count sql.NullInt16
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&student_id,
		&branch_id,
		&paid_sum,
		&total_sum,
		&course_count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Payment{
		Id:          id.String,
		StudentId:   student_id.String,
		BranchId:    branch_id.String,
		PaidSum:     paid_sum.Float64,
		TotalSum:    total_sum.Float64,
		CourseCount: int64(course_count.Int16),
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	return
}

func (c *PaymentRepo) GetAll(ctx context.Context, req *user_service.GetListPaymentRequest) (resp *user_service.GetListPaymentResponse, err error) {

	resp = &user_service.GetListPaymentResponse{}

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
			student_id,
			branch_id,
			paid_sum,
			total_sum,
			course_count,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "payment"
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
		filter += " AND branch_id ='" + req.GetSearch() + "'"
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
			id           sql.NullString
			student_id   sql.NullString
			branch_id    sql.NullString
			paid_sum     sql.NullFloat64
			total_sum    sql.NullFloat64
			course_count sql.NullInt64
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&student_id,
			&branch_id,
			&paid_sum,
			&total_sum,
			&course_count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Payments = append(resp.Payments, &user_service.Payment{
			Id:          id.String,
			StudentId:   student_id.String,
			BranchId:    branch_id.String,
			PaidSum:     paid_sum.Float64,
			TotalSum:    total_sum.Float64,
			CourseCount: (course_count.Int64),
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return
}

func (c *PaymentRepo) Update(ctx context.Context, req *user_service.UpdatePayment) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "payment"
			SET
				student_id = :student_id,
				branch_id = :branch_id,
				paid_sum = :paid_sum,
				total_sum =:total_sum,
				course_count =:course_count,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":           req.GetId(),
		"student_id":   req.GetStudentId(),
		"branch_id":    req.GetBranchId(),
		"paid_sum":     req.GetPaidSum(),
		"total_sum":    req.GetTotalSum(),
		"course_count": req.GetCourseCount(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *PaymentRepo) Delete(ctx context.Context, req *user_service.PaymentPrimaryKey) error {

	query := `DELETE FROM "payment" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
