package postgres

import (
	"context"
	"database/sql"
	"fmt"
	schedule_service "schedule_service/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) storage.GroupRepoI {
	return &GroupRepo{
		db: db,
	}
}

func (c *GroupRepo) Create(ctx context.Context, req *schedule_service.CreateGroup) (resp *schedule_service.GroupPrimaryKey, err error) {

	var groupUnuqieId string
	max_query := `SELECT MAX(unique_id) FROM "group" WHERE type = $1`
	err = c.db.QueryRow(ctx, max_query, req.GetType()).Scan(&groupUnuqieId)
	if err != nil {
		if err.Error() != "can't scan into dest[0]: cannot scan null into *string" {
			return resp, err
		} else {
			groupUnuqieId = string(req.GetType()[0]) + "0000000"
		}
	}

	digit, _ := strconv.Atoi(groupUnuqieId[1:])
	var id = uuid.New().String()
	query := `INSERT INTO "group" (
				id,
				unique_id,
				branch_id,
				type,
				teacher_id,
				support_teacher_id,
				updated_at
			) VALUES ($1, $2,$3,$4,$5,$6, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id,
		string(req.GetType()[0])+helper.GetGroupUniqueId(digit),
		req.BranchId,
		req.Type,
		req.TeacherId,
		req.SupportTeacherId,
	)

	if err != nil {
		return nil, err
	}
	return &schedule_service.GroupPrimaryKey{Id: id}, nil
}

func (c *GroupRepo) GetByPKey(ctx context.Context, req *schedule_service.GroupPrimaryKey) (resp *schedule_service.Group, err error) {

	query := `
		SELECT
			id,
			unique_id,
			branch_id,
			type,
			teacher_id,
			support_teacher_id,
			created_at,
			updated_at
		FROM "group"
		WHERE id = $1
	`

	var (
		id                 sql.NullString
		unique_id          sql.NullString
		branch_id          sql.NullString
		group_type         sql.NullString
		teacher_id         sql.NullString
		support_teacher_id sql.NullString
		createdAt          sql.NullString
		updatedAt          sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&unique_id,
		&branch_id,
		&group_type,
		&teacher_id,
		&support_teacher_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &schedule_service.Group{
		Id:               id.String,
		UniqueID:         unique_id.String,
		BranchId:         branch_id.String,
		Type:             group_type.String,
		TeacherId:        teacher_id.String,
		SupportTeacherId: support_teacher_id.String,
		CreatedAt:        createdAt.String,
		UpdatedAt:        updatedAt.String,
	}

	return
}

func (c *GroupRepo) GetAll(ctx context.Context, req *schedule_service.GetListGroupRequest) (resp *schedule_service.GetListGroupResponse, err error) {

	resp = &schedule_service.GetListGroupResponse{}

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
			unique_id,
			branch_id,
			type,
			teacher_id,
			support_teacher_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "group"
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
		filter += " AND type ILIKE '%" + req.GetSearch() + "%'" + " OR branch_id ='" + req.GetSearch() + "'" + " OR teacher_id ='" + req.GetSearch() + "'" + " OR support_teacher_id ='" + req.GetSearch() + "'"
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
			id                 sql.NullString
			unique_id          sql.NullString
			branch_id          sql.NullString
			group_type         sql.NullString
			teacher_id         sql.NullString
			support_teacher_id sql.NullString
			createdAt          sql.NullString
			updatedAt          sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&unique_id,
			&branch_id,
			&group_type,
			&teacher_id,
			&support_teacher_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Groups = append(resp.Groups, &schedule_service.Group{
			Id:               id.String,
			UniqueID:         unique_id.String,
			BranchId:         branch_id.String,
			Type:             group_type.String,
			TeacherId:        teacher_id.String,
			SupportTeacherId: support_teacher_id.String,
			CreatedAt:        createdAt.String,
			UpdatedAt:        updatedAt.String,
		})
	}

	return
}

func (c *GroupRepo) Update(ctx context.Context, req *schedule_service.UpdateGroup) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "group"
			SET
				type = :type,
				teacher_id = :teacher_id,
				support_teacher_id = :support_teacher_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":                 req.GetId(),
		"type":               req.GetType(),
		"teacher_id":         req.GetTeacherId(),
		"support_teacher_id": req.GetSupportTeacherId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	fmt.Println("rowsAffectedresult:", result.RowsAffected())

	return result.RowsAffected(), nil
}

func (c *GroupRepo) Delete(ctx context.Context, req *schedule_service.GroupPrimaryKey) error {

	query := `DELETE FROM "group" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
