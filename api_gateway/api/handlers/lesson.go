package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateLesson godoc
// @ID create_lesson
// @Router /lesson [POST]
// @Summary Create Lesson
// @Description  Create Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateLesson true "CreateLessonBody"
// @Success 200 {object} http.Response{data=schedule_service.Lesson} "GetLessonBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateLesson(c *gin.Context) {

	var Lesson schedule_service.CreateLesson

	err := c.ShouldBindJSON(&Lesson)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Lesson().Create(
		c.Request.Context(),
		&Lesson,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetLessonByID godoc
// @ID get_lesson_by_id
// @Router /lesson/{id} [GET]
// @Summary Get Lesson  By ID
// @Description Get Lesson  By ID
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.Lesson} "LessonBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetLessonByID(c *gin.Context) {

	LessonID := c.Param("id")

	if !util.IsValidUUID(LessonID) {
		h.handleResponse(c, http.InvalidArgument, "Lesson id is an invalid uuid")
		return
	}

	resp, err := h.services.Lesson().GetByID(
		context.Background(),
		&schedule_service.LessonPrimaryKey{
			Id: LessonID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
		Limit:  10000,
		Search: LessonID,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	resp.Tasks = tasks.Tasks
	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetLessonList godoc
// @ID get_lesson_list
// @Router /lesson [GET]
// @Summary Get Lessons List
// @Description  Get Lessons List
// @Tags Lesson
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListLessonResponse} "GetAllLessonResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetLessonList(c *gin.Context) {

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

	resp, err := h.services.Lesson().GetList(
		context.Background(),
		&schedule_service.GetListLessonRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)
	for _, lesson := range resp.Lessons {
		tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
			Limit:  10000,
			Search: lesson.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		lesson.Tasks = tasks.Tasks
	}
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateLesson godoc
// @ID update_lesson
// @Router /lesson/{id} [PUT]
// @Summary Update Lesson
// @Description Update Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateLesson true "UpdateLesson"
// @Success 200 {object} http.Response{data=schedule_service.Lesson} "Lesson data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateLesson(c *gin.Context) {

	var Lesson schedule_service.UpdateLesson

	Lesson.Id = c.Param("id")

	if !util.IsValidUUID(Lesson.Id) {
		h.handleResponse(c, http.InvalidArgument, "Lesson id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Lesson)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Lesson().Update(
		c.Request.Context(),
		&Lesson,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteLesson godoc
// @ID delete_lesson
// @Router /lesson/{id} [DELETE]
// @Summary Delete Lesson
// @Description Delete Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Lesson data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteLesson(c *gin.Context) {

	LessonId := c.Param("id")

	if !util.IsValidUUID(LessonId) {
		h.handleResponse(c, http.InvalidArgument, "Lesson id is an invalid uuid")
		return
	}

	resp, err := h.services.Lesson().Delete(
		c.Request.Context(),
		&schedule_service.LessonPrimaryKey{Id: LessonId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
