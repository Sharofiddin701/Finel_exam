package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
// @ID create_group
// @Router /group [POST]
// @Summary Create Group
// @Description  Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Param roleid query string true "roleid"
// @Param profile body schedule_service.CreateGroup true "CreateGroupBody"
// @Success 200 {object} http.Response{data=schedule_service.Group} "GetGroupBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateGroup(c *gin.Context) {

	var Group schedule_service.CreateGroup

	roleID := c.Query("roleid")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	err := c.ShouldBindJSON(&Group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Group().Create(
		c.Request.Context(),
		&Group,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	from := time.Now().AddDate(0, 0, 0)
	fromstring := from.Format("2006-01-02")
	_, err = h.services.Jurnal().Create(context.Background(), &schedule_service.CreateJurnal{
		GroupId: resp.Id,
		From:    fromstring,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetGroupByID godoc
// @ID get_group_by_id
// @Router /group/{id} [GET]
// @Summary Get Group  By ID
// @Description Get Group  By ID
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=schedule_service.Group} "GroupBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroupByID(c *gin.Context) {

	GroupID := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(GroupID) {
		h.handleResponse(c, http.InvalidArgument, "Group id is an invalid uuid")
		return
	}

	resp, err := h.services.Group().GetByID(
		context.Background(),
		&schedule_service.GroupPrimaryKey{
			Id: GroupID,
		},
	)

	students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
		Limit:  10000,
		Search: GroupID,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	var studentDs []*schedule_service.StudentForRep
	for _, student := range students.Students {
		studentD := &schedule_service.StudentForRep{
			Id:        student.Id,
			FullName:  student.FullName,
			Phone:     student.Phone,
			Password:  student.Password,
			GroupId:   student.GroupId,
			BranchId:  student.BranchId,
			RoleId:    student.RoleId,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Login:     student.Login,
			///////////////////////////////////////
		}
		studentDs = append(studentDs, studentD)
	}
	jurnal, err := h.services.Jurnal().GetByID(context.Background(), &schedule_service.JurnalPrimaryKey{
		Id: GroupID,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	schedules, err := h.services.Schedule().GetList(context.Background(), &schedule_service.GetListScheduleRequest{
		Limit:  20000,
		Search: jurnal.Id,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	for _, schedule := range schedules.Schedules {
		lesson, err := h.services.Lesson().GetByID(context.Background(), &schedule_service.LessonPrimaryKey{
			Id: schedule.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
			Limit:  20000,
			Search: lesson.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		lesson.Tasks = tasks.Tasks
		schedule.Lesson = lesson

	}
	jurnal.Schedules = schedules.Schedules
	resp.Jurnal = jurnal
	resp.Students = studentDs
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetGroupList godoc
// @ID get_group_list
// @Router /group [GET]
// @Summary Get Groups List
// @Description  Get Groups List
// @Tags Group
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListGroupResponse} "GetAllGroupResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroupList(c *gin.Context) {
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if c.GetHeader("role_id") == config.RoleClient {
		h.handleResponse(c, http.OK, struct{}{})
		return
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.Group().GetList(
		context.Background(),
		&schedule_service.GetListGroupRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)

	for _, group := range resp.Groups {
		students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
			Limit:  10000,
			Search: group.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}

		var studentDs []*schedule_service.StudentForRep
		for _, student := range students.Students {
			studentD := &schedule_service.StudentForRep{
				Id:        student.Id,
				FullName:  student.FullName,
				Phone:     student.Phone,
				Password:  student.Password,
				GroupId:   student.GroupId,
				BranchId:  student.BranchId,
				RoleId:    student.RoleId,
				CreatedAt: student.CreatedAt,
				UpdatedAt: student.UpdatedAt,
				Login:     student.Login,
				///////////////////////////////////////
			}
			studentDs = append(studentDs, studentD)
		}

		jurnal, err := h.services.Jurnal().GetByID(context.Background(), &schedule_service.JurnalPrimaryKey{
			Id: group.Id,
		})

		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}

		schedules, err := h.services.Schedule().GetList(context.Background(), &schedule_service.GetListScheduleRequest{
			Limit:  20000,
			Search: jurnal.Id,
		})
		if err != nil {

			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		group.Jurnal = jurnal
		group.Students = studentDs
		jurnal.Students = studentDs

		for _, schedule := range schedules.Schedules {

			lesson, err := h.services.Lesson().GetByID(context.Background(), &schedule_service.LessonPrimaryKey{
				Id: schedule.Id,
			})

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}
			tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
				Limit:  20000,
				Search: lesson.Id,
			})

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}
			lesson.Tasks = tasks.Tasks
			schedule.Lesson = lesson
			jurnal.Schedules = schedules.Schedules

		}

	}

	fmt.Println(resp)
	h.handleResponse(c, http.OK, resp)
}

// UpdateGroup godoc
// @ID update_group
// @Router /group/{id} [PUT]
// @Summary Update Group
// @Description Update Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body schedule_service.UpdateGroup true "UpdateGroup"
// @Success 200 {object} http.Response{data=schedule_service.Group} "Group data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateGroup(c *gin.Context) {

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	var Group schedule_service.UpdateGroup

	Group.Id = c.Param("id")

	if !util.IsValidUUID(Group.Id) {
		h.handleResponse(c, http.InvalidArgument, "Group id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Group().Update(
		c.Request.Context(),
		&Group,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteGroup godoc
// @ID delete_group
// @Router /group/{id} [DELETE]
// @Summary Delete Group
// @Description Delete Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Group data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteGroup(c *gin.Context) {
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}

	GroupId := c.Param("id")

	if !util.IsValidUUID(GroupId) {
		h.handleResponse(c, http.InvalidArgument, "Group id is an invalid uuid")
		return
	}

	resp, err := h.services.Group().Delete(
		c.Request.Context(),
		&schedule_service.GroupPrimaryKey{Id: GroupId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
