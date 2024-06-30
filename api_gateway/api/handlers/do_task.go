package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateDoTask godoc
// @ID create_dotask
// @Router /dotask [POST]
// @Summary Create DoTask
// @Description  Create DoTask
// @Tags DoTask
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateDoTask true "CreateDoTaskBody"
// @Success 200 {object} http.Response{data=schedule_service.DoTask} "GetDoTaskBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateDoTask(c *gin.Context) {

	var DoTask schedule_service.CreateDoTask

	err := c.ShouldBindJSON(&DoTask)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	task, err := h.services.Task().GetByID(context.Background(), &schedule_service.TaskPrimaryKey{Id: DoTask.TaskId})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	var (
		year  string
		month string
		day   string
	)

	for i := 0; i < 4; i++ {
		year += string(task.Deadline[i])
	}
	for i := 5; i < 7; i++ {
		month += string(task.Deadline[i])
	}
	for i := 8; i < len(task.Deadline); i++ {
		day += string(task.Deadline[i])
	}

	yearint, _ := strconv.Atoi(year)
	monthint, _ := strconv.Atoi(month)
	dayint, _ := strconv.Atoi(day)

	if time.Now().Year() == yearint && time.Now().Month() == time.Month(monthint) && time.Now().Day() <= dayint {
		DoTask.Score = float64(task.Score)

	}
	if time.Now().Year() == yearint && time.Now().Month() == time.Month(monthint) && time.Now().Day()-dayint <= 3 {
		DoTask.Score = float64(task.Score) / 2
	}
	if time.Now().Year() == yearint && time.Now().Month() == time.Month(monthint) && time.Now().Day()-dayint > 3 {
		DoTask.Score = 0

	}

	resp, err := h.services.DoTask().Create(
		c.Request.Context(),
		&DoTask,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetDoTaskByID godoc
// @ID get_dotask_by_id
// @Router /dotask/{id} [GET]
// @Summary Get DoTask  By ID
// @Description Get DoTask  By ID
// @Tags DoTask
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.DoTask} "DoTaskBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDoTaskByID(c *gin.Context) {

	DoTaskID := c.Param("id")

	if !util.IsValidUUID(DoTaskID) {
		h.handleResponse(c, http.InvalidArgument, "DoTask id is an invalid uuid")
		return
	}

	resp, err := h.services.DoTask().GetByID(
		context.Background(),
		&schedule_service.DoTaskPrimaryKey{
			Id: DoTaskID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetDoTaskList godoc
// @ID get_dotask_list
// @Router /dotask [GET]
// @Summary Get DoTasks List
// @Description  Get DoTasks List
// @Tags DoTask
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListDoTaskResponse} "GetAllDoTaskResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDoTaskList(c *gin.Context) {

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

	resp, err := h.services.DoTask().GetList(
		context.Background(),
		&schedule_service.GetListDoTaskRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateDoTask godoc
// @ID update_dotask
// @Router /dotask/{id} [PUT]
// @Summary Update DoTask
// @Description Update DoTask
// @Tags DoTask
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateDoTask true "UpdateDoTask"
// @Success 200 {object} http.Response{data=schedule_service.DoTask} "DoTask data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateDoTask(c *gin.Context) {

	var DoTask schedule_service.UpdateDoTask

	DoTask.Id = c.Param("id")

	if !util.IsValidUUID(DoTask.Id) {
		h.handleResponse(c, http.InvalidArgument, "DoTask id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&DoTask)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.DoTask().Update(
		c.Request.Context(),
		&DoTask,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteDoTask godoc
// @ID delete_dotask
// @Router /dotask/{id} [DELETE]
// @Summary Delete DoTask
// @Description Delete DoTask
// @Tags DoTask
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "DoTask data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteDoTask(c *gin.Context) {

	DoTaskId := c.Param("id")

	if !util.IsValidUUID(DoTaskId) {
		h.handleResponse(c, http.InvalidArgument, "DoTask id is an invalid uuid")
		return
	}

	resp, err := h.services.DoTask().Delete(
		c.Request.Context(),
		&schedule_service.DoTaskPrimaryKey{Id: DoTaskId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
