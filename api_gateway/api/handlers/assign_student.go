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

// CreateAssignStudent godoc
// @ID create_assignstudent
// @Router /assign-student [POST]
// @Summary Create AssignStudent
// @Description  Create AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateAssignStudent true "CreateAssignStudentBody"
// @Success 200 {object} http.Response{data=schedule_service.AssignStudent} "GetAssignStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAssignStudent(c *gin.Context) {

	var AssignStudent schedule_service.CreateAssignStudent
	event, err := h.services.Event().GetByID(context.Background(), &schedule_service.EventPrimaryKey{
		Id: AssignStudent.EventId,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	var reqhour string
	var reqminute string
	for i := 0; i < 2; i++ {
		reqhour += string(event.StartTime[i])
	}
	for i := 3; i < 5; i++ {
		reqminute += string(event.StartTime[i])
	}

	reqhourint, _ := strconv.Atoi(reqhour)
	reqminuteint, _ := strconv.Atoi(reqminute)

	reqhourint -= 3
	reqminuteint -= 1
	hour := time.Now().Hour()
	minute := time.Now().Minute()

	if hour > reqhourint || minute > reqminuteint {
		h.handleResponse(c, http.BadRequest, "3 hours left to the event you canot register to the even")
		return
	}
	err = c.ShouldBindJSON(&AssignStudent)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AssignStudent().Create(
		c.Request.Context(),
		&AssignStudent,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAssignStudentByID godoc
// @ID get_assignstudent_by_id
// @Router /assign-student/{id} [GET]
// @Summary Get AssignStudent  By ID
// @Description Get AssignStudent  By ID
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.AssignStudent} "AssignStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAssignStudentByID(c *gin.Context) {

	AssignStudentID := c.Param("id")

	if !util.IsValidUUID(AssignStudentID) {
		h.handleResponse(c, http.InvalidArgument, "AssignStudent id is an invalid uuid")
		return
	}

	resp, err := h.services.AssignStudent().GetByID(
		context.Background(),
		&schedule_service.AssignStudentPrimaryKey{
			Id: AssignStudentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetAssignStudentList godoc
// @ID get_assignstudent_list
// @Router /assign-student [GET]
// @Summary Get AssignStudents List
// @Description  Get AssignStudents List
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListAssignStudentResponse} "GetAllAssignStudentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAssignStudentList(c *gin.Context) {

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

	resp, err := h.services.AssignStudent().GetList(
		context.Background(),
		&schedule_service.GetListAssignStudentRequest{
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

// UpdateAssignStudent godoc
// @ID update_assignstudent
// @Router /assign-student/{id} [PUT]
// @Summary Update AssignStudent
// @Description Update AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateAssignStudent true "UpdateAssignStudent"
// @Success 200 {object} http.Response{data=schedule_service.AssignStudent} "AssignStudent data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAssignStudent(c *gin.Context) {

	var AssignStudent schedule_service.UpdateAssignStudent

	AssignStudent.Id = c.Param("id")

	if !util.IsValidUUID(AssignStudent.Id) {
		h.handleResponse(c, http.InvalidArgument, "AssignStudent id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&AssignStudent)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AssignStudent().Update(
		c.Request.Context(),
		&AssignStudent,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAssignStudent godoc
// @ID delete_assignstudent
// @Router /assign-student/{id} [DELETE]
// @Summary Delete AssignStudent
// @Description Delete AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "AssignStudent data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAssignStudent(c *gin.Context) {

	AssignStudentId := c.Param("id")

	if !util.IsValidUUID(AssignStudentId) {
		h.handleResponse(c, http.InvalidArgument, "AssignStudent id is an invalid uuid")
		return
	}

	resp, err := h.services.AssignStudent().Delete(
		c.Request.Context(),
		&schedule_service.AssignStudentPrimaryKey{Id: AssignStudentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
