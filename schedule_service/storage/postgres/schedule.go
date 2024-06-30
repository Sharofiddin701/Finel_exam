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

type ScheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) storage.ScheduleRepoI {
	return &ScheduleRepo{
		db: db,
	}
}

func (c *ScheduleRepo) Create(ctx context.Context, req *schedule_service.CreateSchedule) (resp *schedule_service.SchedulePrimaryKey, err error) {

	var id = uuid.New().String()
	query := `INSERT INTO "schedule" (
				id,
				journal_id,
				start_time,
				end_time,
				date,
				updated_at
			) VALUES ($1, $2,$3,$4,$5, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		req.JournalId,
		req.StartTime,
		req.EndTime,
		req.Date,
	)

	if err != nil {
		return nil, err
	}
	return &schedule_service.SchedulePrimaryKey{Id: id}, nil
}

func (c *ScheduleRepo) GetByPKey(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (resp *schedule_service.Schedule, err error) {

	query := `
		SELECT
			id,
			journal_id,
			start_time,
			end_time,
			date,
			created_at,
			updated_at
		FROM "schedule"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		journal_id sql.NullString
		start_time sql.NullString
		end_time   sql.NullString
		date       sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&journal_id,
		&start_time,
		&end_time,
		&date,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Schedule{
		Id:        id.String,
		JournalId: journal_id.String,
		StartTime: start_time.String,
		EndTime:   end_time.String,
		Date:      date.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ScheduleRepo) GetAll(ctx context.Context, req *schedule_service.GetListScheduleRequest) (resp *schedule_service.GetListScheduleResponse, err error) {

	resp = &schedule_service.GetListScheduleResponse{}

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
			journal_id,
			start_time,
			end_time,
			date,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "schedule"
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
		filter += " AND date ILIKE '%" + req.GetSearch() + "%'" + " OR journal_id='" + req.GetSearch() + "'"
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
			journal_id sql.NullString
			start_time sql.NullString
			end_time   sql.NullString
			date       sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&journal_id,
			&start_time,
			&end_time,
			&date,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Schedules = append(resp.Schedules, &schedule_service.Schedule{
			Id:        id.String,
			JournalId: journal_id.String,
			StartTime: start_time.String,
			EndTime:   end_time.String,
			Date:      date.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ScheduleRepo) Update(ctx context.Context, req *schedule_service.UpdateSchedule) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "schedule"
			SET
				journal_id = :journal_id,
				start_time = :start_time,
				end_time = :end_time,
				date = :date,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.Id,
		"journal_id": req.JournalId,
		"start_time": req.StartTime,
		"end_time":   req.EndTime,
		"date":       req.Date,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *ScheduleRepo) Delete(ctx context.Context, req *schedule_service.SchedulePrimaryKey) error {

	query := `DELETE FROM "schedule" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c *ScheduleRepo) GetScheduleReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error) {
	todate := time.Now().AddDate(0, 0, 7)
	todatestring := todate.Format("2006-01-02")
	fromdate := time.Now().AddDate(0, 0, 0)
	fromdatestring := fromdate.Format("2006-01-02")
	resp = &schedule_service.ScheduleReportResponse{}
	var query string
	query = `
				SELECT 
				g.unique_id,
				g.type,
				b.name,
				s.start_time,
				s.end_time,
				t.full_name,
				st.full_name
				FROM "group" AS g
				JOIN branch AS b ON g.branch_id = b.id
				JOIN journal AS j ON j.group_id = g.id
				JOIN schedule AS s ON s.journal_id =  j.id
				JOIN teacher AS t ON g.teacher_id = t.id
				JOIN support_teacher AS st ON g.support_teacher_id = st.id
				WHERE g.id = $1 OR t.id = $1 OR st.id = $1 AND s.date BETWEEN $2 AND $3
			`

	rows, err := c.db.Query(ctx, query, req.BranchId, fromdatestring, todatestring)
	fmt.Println(fromdatestring)
	fmt.Println(todatestring)

	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			unique_id          sql.NullString
			typee              sql.NullString
			number_of_students sql.NullInt64
			start_time         sql.NullString
			end_time           sql.NullString
			branch             sql.NullString
			teacher            sql.NullString
			support_teacher    sql.NullString
		)

		err := rows.Scan(
			&unique_id,
			&typee,
			&branch,
			&start_time,
			&end_time,
			&teacher,
			&support_teacher,
		)
		if err != nil {
			return resp, err
		}
		query1 := `
				SELECT
					COUNT(s.group_id)
				FROM student AS s
				JOIN "group" AS g ON g.id = s.group_id
				WHERE g.unique_id = $1 GROUP BY s.group_id
		`
		err = c.db.QueryRow(ctx, query1, unique_id).Scan(
			&number_of_students,
		)
		if err != nil {
			return resp, err
		}

		resp.Schedules = append(resp.Schedules, &schedule_service.ScheduleReport{
			GroupIdUnique:      unique_id.String,
			GroupType:          typee.String,
			NumberOfStudents:   number_of_students.Int64,
			StartTime:          start_time.String,
			EndTime:            end_time.String,
			BranchName:         branch.String,
			TeacherName:        teacher.String,
			SupportTeacherName: support_teacher.String,
		})

	}
	return
}

func (c *ScheduleRepo) GetScheduleMonthReport(ctx context.Context, req *schedule_service.ScheduleReportRequest) (resp *schedule_service.ScheduleReportResponse, err error) {
	todate := time.Now().AddDate(0, 1, 1)
	todatestring := todate.Format("2006-01-02")
	fromdate := time.Now().AddDate(0, 0, 0)
	fromdatestring := fromdate.Format("2006-01-02")
	resp = &schedule_service.ScheduleReportResponse{}

	query := `
				SELECT 
				g.unique_id,
				g.type,
				b.name,
				s.start_time,
				s.end_time,
				t.full_name,
				st.full_name
				FROM "group" AS g
				JOIN branch AS b ON g.branch_id = b.id
				JOIN journal AS j ON j.group_id = g.id
				JOIN schedule AS s ON s.journal_id =  j.id
				JOIN teacher AS t ON g.teacher_id = t.id
				JOIN support_teacher AS st ON g.support_teacher_id = st.id
				WHERE g.id = $1 OR t.id = $1 OR st.id = $1 AND s.date BETWEEN $2 AND $3
			`

	rows, err := c.db.Query(ctx, query, req.BranchId, fromdatestring, todatestring)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			unique_id          sql.NullString
			typee              sql.NullString
			number_of_students sql.NullInt64
			start_time         sql.NullString
			end_time           sql.NullString
			branch             sql.NullString
			teacher            sql.NullString
			support_teacher    sql.NullString
		)

		err := rows.Scan(
			&unique_id,
			&typee,
			&branch,
			&start_time,
			&end_time,
			&teacher,
			&support_teacher,
		)
		if err != nil {
			return resp, err
		}
		query1 := `
				SELECT
					COUNT(s.group_id)
				FROM student AS s
				JOIN "group" AS g ON g.id = s.group_id
				WHERE g.unique_id = $1 GROUP BY s.group_id
		`
		err = c.db.QueryRow(ctx, query1, unique_id).Scan(
			&number_of_students,
		)
		if err != nil {
			return resp, err
		}

		resp.Schedules = append(resp.Schedules, &schedule_service.ScheduleReport{
			GroupIdUnique:      unique_id.String,
			GroupType:          typee.String,
			NumberOfStudents:   number_of_students.Int64,
			StartTime:          start_time.String,
			EndTime:            end_time.String,
			BranchName:         branch.String,
			TeacherName:        teacher.String,
			SupportTeacherName: support_teacher.String,
		})

	}
	return
}
