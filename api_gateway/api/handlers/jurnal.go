package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateJurnal godoc
// @ID create_jurnal
// @Router /jurnal [POST]
// @Summary Create Jurnal
// @Description  Create Jurnal
// @Tags Jurnal
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateJurnal true "CreateJurnalBody"
// @Success 200 {object} http.Response{data=schedule_service.Jurnal} "GetJurnalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJurnal(c *gin.Context) {

	var Jurnal schedule_service.CreateJurnal

	err := c.ShouldBindJSON(&Jurnal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	var (
		month string
		day   string
	)

	for i := 5; i < 7; i++ {
		month += string(Jurnal.From[i])
	}
	for i := 8; i < len(Jurnal.From); i++ {
		day += string(Jurnal.From[i])
	}

	monthint, _ := strconv.Atoi(month)
	dayint, _ := strconv.Atoi(day)
	if monthint > 12 || dayint > 31 {
		h.handleResponse(c, http.BadRequest, "check month or day it is not real")
		return
	}
	resp, err := h.services.Jurnal().Create(
		c.Request.Context(),
		&Jurnal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetJurnalByID godoc
// @ID get_jurnal_by_id
// @Router /jurnal/{id} [GET]
// @Summary Get Jurnal  By ID
// @Description Get Jurnal  By ID
// @Tags Jurnal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.Jurnal} "JurnalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJurnalByID(c *gin.Context) {

	JurnalID := c.Param("id")

	if !util.IsValidUUID(JurnalID) {
		h.handleResponse(c, http.InvalidArgument, "Jurnal id is an invalid uuid")
		return
	}

	resp, err := h.services.Jurnal().GetByID(
		context.Background(),
		&schedule_service.JurnalPrimaryKey{
			Id: JurnalID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
		Search: resp.GroupId,
		Limit:  1000,
	})

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

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	schedules, err := h.services.Schedule().GetList(context.Background(), &schedule_service.GetListScheduleRequest{
		Limit:  10000,
		Search: resp.Id,
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
	resp.Schedules = schedules.Schedules
	resp.Students = studentDs

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetJurnalList godoc
// @ID get_jurnal_list
// @Router /jurnal [GET]
// @Summary Get Jurnals List
// @Description  Get Jurnals List
// @Tags Jurnal
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListJurnalResponse} "GetAllJurnalResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJurnalList(c *gin.Context) {

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

	resp, err := h.services.Jurnal().GetList(
		context.Background(),
		&schedule_service.GetListJurnalRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)
	for _, jurnal := range resp.Jurnals {

		students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
			Limit:  10000,
			Search: jurnal.GroupId,
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

		schedules, err := h.services.Schedule().GetList(context.Background(), &schedule_service.GetListScheduleRequest{
			Limit:  10000,
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

		jurnal.Students = studentDs

	}

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateJurnal godoc
// @ID update_jurnal
// @Router /jurnal/{id} [PUT]
// @Summary Update Jurnal
// @Description Update Jurnal
// @Tags Jurnal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateJurnal true "UpdateJurnal"
// @Success 200 {object} http.Response{data=schedule_service.Jurnal} "Jurnal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJurnal(c *gin.Context) {

	var Jurnal schedule_service.UpdateJurnal

	Jurnal.Id = c.Param("id")

	if !util.IsValidUUID(Jurnal.Id) {
		h.handleResponse(c, http.InvalidArgument, "Jurnal id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Jurnal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Jurnal().Update(
		c.Request.Context(),
		&Jurnal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteJurnal godoc
// @ID delete_jurnal
// @Router /jurnal/{id} [DELETE]
// @Summary Delete Jurnal
// @Description Delete Jurnal
// @Tags Jurnal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Jurnal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJurnal(c *gin.Context) {

	JurnalId := c.Param("id")

	if !util.IsValidUUID(JurnalId) {
		h.handleResponse(c, http.InvalidArgument, "Jurnal id is an invalid uuid")
		return
	}

	resp, err := h.services.Jurnal().Delete(
		c.Request.Context(),
		&schedule_service.JurnalPrimaryKey{Id: JurnalId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
