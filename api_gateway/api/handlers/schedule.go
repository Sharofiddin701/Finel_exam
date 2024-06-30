package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateSchedule godoc
// @ID create_schedule
// @Router /schedule [POST]
// @Summary Create Schedule
// @Description  Create Schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateSchedule true "CreateScheduleBody"
// @Success 200 {object} http.Response{data=schedule_service.Schedule} "GetScheduleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSchedule(c *gin.Context) {

	var Schedule schedule_service.CreateSchedule

	err := c.ShouldBindJSON(&Schedule)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Schedule().Create(
		c.Request.Context(),
		&Schedule,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetScheduleByID godoc
// @ID get_schedule_by_id
// @Router /schedule/{id} [GET]
// @Summary Get Schedule  By ID
// @Description Get Schedule  By ID
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.Schedule} "ScheduleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleByID(c *gin.Context) {

	ScheduleID := c.Param("id")

	if !util.IsValidUUID(ScheduleID) {
		h.handleResponse(c, http.InvalidArgument, "Schedule id is an invalid uuid")
		return
	}

	resp, err := h.services.Schedule().GetByID(
		context.Background(),
		&schedule_service.SchedulePrimaryKey{
			Id: ScheduleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	lesson, err := h.services.Lesson().GetByID(context.Background(), &schedule_service.LessonPrimaryKey{
		Id: ScheduleID,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Lesson = lesson
	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetScheduleList godoc
// @ID get_schedule_list
// @Router /schedule [GET]
// @Summary Get Schedules List
// @Description  Get Schedules List
// @Tags Schedule
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListScheduleResponse} "GetAllScheduleResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleList(c *gin.Context) {

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

	resp, err := h.services.Schedule().GetList(
		context.Background(),
		&schedule_service.GetListScheduleRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)
	for _, schedule := range resp.Schedules {
		lesson, err := h.services.Lesson().GetByID(context.Background(), &schedule_service.LessonPrimaryKey{
			Id: schedule.Id,
		})
		tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
			Limit:  10000,
			Search: lesson.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		lesson.Tasks = tasks.Tasks

		schedule.Lesson = lesson
	}

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateSchedule godoc
// @ID update_schedule
// @Router /schedule/{id} [PUT]
// @Summary Update Schedule
// @Description Update Schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateSchedule true "UpdateSchedule"
// @Success 200 {object} http.Response{data=schedule_service.Schedule} "Schedule data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSchedule(c *gin.Context) {

	var Schedule schedule_service.UpdateSchedule

	Schedule.Id = c.Param("id")

	if !util.IsValidUUID(Schedule.Id) {
		h.handleResponse(c, http.InvalidArgument, "Schedule id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Schedule)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Schedule().Update(
		c.Request.Context(),
		&Schedule,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteSchedule godoc
// @ID delete_schedule
// @Router /schedule/{id} [DELETE]
// @Summary Delete Schedule
// @Description Delete Schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Schedule data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSchedule(c *gin.Context) {

	ScheduleId := c.Param("id")

	if !util.IsValidUUID(ScheduleId) {
		h.handleResponse(c, http.InvalidArgument, "Schedule id is an invalid uuid")
		return
	}

	resp, err := h.services.Schedule().Delete(
		c.Request.Context(),
		&schedule_service.SchedulePrimaryKey{Id: ScheduleId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
