package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @ID create_task
// @Router /task [POST]
// @Summary Create Task
// @Description  Create Task
// @Tags Task
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateTask true "CreateTaskBody"
// @Success 200 {object} http.Response{data=schedule_service.Task} "GetTaskBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateTask(c *gin.Context) {

	var Task schedule_service.CreateTask

	err := c.ShouldBindJSON(&Task)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Task().Create(
		c.Request.Context(),
		&Task,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetTaskByID godoc
// @ID get_task_by_id
// @Router /task/{id} [GET]
// @Summary Get Task  By ID
// @Description Get Task  By ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.Task} "TaskBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTaskByID(c *gin.Context) {

	TaskID := c.Param("id")

	if !util.IsValidUUID(TaskID) {
		h.handleResponse(c, http.InvalidArgument, "Task id is an invalid uuid")
		return
	}

	resp, err := h.services.Task().GetByID(
		context.Background(),
		&schedule_service.TaskPrimaryKey{
			Id: TaskID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetTaskList godoc
// @ID get_task_list
// @Router /task [GET]
// @Summary Get Tasks List
// @Description  Get Tasks List
// @Tags Task
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListTaskResponse} "GetAllTaskResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTaskList(c *gin.Context) {

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

	resp, err := h.services.Task().GetList(
		context.Background(),
		&schedule_service.GetListTaskRequest{
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

// UpdateTask godoc
// @ID update_task
// @Router /task/{id} [PUT]
// @Summary Update Task
// @Description Update Task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateTask true "UpdateTask"
// @Success 200 {object} http.Response{data=schedule_service.Task} "Task data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateTask(c *gin.Context) {

	var Task schedule_service.UpdateTask

	Task.Id = c.Param("id")

	if !util.IsValidUUID(Task.Id) {
		h.handleResponse(c, http.InvalidArgument, "Task id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Task)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Task().Update(
		c.Request.Context(),
		&Task,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteTask godoc
// @ID delete_task
// @Router /task/{id} [DELETE]
// @Summary Delete Task
// @Description Delete Task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Task data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteTask(c *gin.Context) {

	TaskId := c.Param("id")

	if !util.IsValidUUID(TaskId) {
		h.handleResponse(c, http.InvalidArgument, "Task id is an invalid uuid")
		return
	}

	resp, err := h.services.Task().Delete(
		c.Request.Context(),
		&schedule_service.TaskPrimaryKey{Id: TaskId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
