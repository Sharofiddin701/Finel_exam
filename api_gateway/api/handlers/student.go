package handlers

import (
	"api_gateway/api/http"
	"api_gateway/api/models"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @ID create_Student
// @Router /student [POST]
// @Summary Create Student
// @Description  Create Student
// @Tags Student
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateStudent true "CreateStudentBody"
// @Success 200 {object} http.Response{data=user_service.Student} "GetStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateStudent(c *gin.Context) {

	var Student user_service.CreateStudent
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}

	err := c.ShouldBindJSON(&Student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Student().Create(
		c.Request.Context(),
		&Student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetStudentByID godoc
// @ID get_Student_by_id
// @Router /student/{id} [GET]
// @Summary Get Student  By ID
// @Description Get Student  By ID
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=user_service.Student} "StudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentByID(c *gin.Context) {

	StudentID := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(StudentID) {
		h.handleResponse(c, http.InvalidArgument, "Student id is an invalid uuid")
		return
	}

	resp, err := h.services.Student().GetByID(
		context.Background(),
		&user_service.StudentPrimaryKey{
			Id: StudentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetStudentList godoc
// @ID get_Student_list
// @Router /student [GET]
// @Summary Get Students List
// @Description  Get Students List
// @Tags Student
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListStudentResponse} "GetAllStudentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentList(c *gin.Context) {
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

	resp, err := h.services.Student().GetList(
		context.Background(),
		&user_service.GetListStudentRequest{
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

// UpdateStudent godoc
// @ID update_Student
// @Router /student/{id} [PUT]
// @Summary Update Student
// @Description Update Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateStudent true "UpdateStudent"
// @Success 200 {object} http.Response{data=user_service.Student} "Student data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateStudent(c *gin.Context) {
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}

	var Student user_service.UpdateStudent

	Student.Id = c.Param("id")

	if !util.IsValidUUID(Student.Id) {
		h.handleResponse(c, http.InvalidArgument, "Student id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Student().Update(
		c.Request.Context(),
		&Student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteStudent godoc
// @ID delete_Student
// @Router /student/{id} [DELETE]
// @Summary Delete Student
// @Description Delete Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Student data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteStudent(c *gin.Context) {

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	StudentId := c.Param("id")

	if !util.IsValidUUID(StudentId) {
		h.handleResponse(c, http.InvalidArgument, "Student id is an invalid uuid")
		return
	}

	resp, err := h.services.Student().Delete(
		c.Request.Context(),
		&user_service.StudentPrimaryKey{Id: StudentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// StudentPanel godoc
// @ID student_panel
// @Router /student-panel/{id} [GET]
// @Summary student_panel
// @Description Panel Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} http.Response{data=models.StudentPanelResponse} "StudentPanelResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) StudentPanel(c *gin.Context) {

	studentId := c.Query("id")
	var resp models.StudentPanelResponse

	student, err := h.services.Student().GetByID(context.Background(), &user_service.StudentPrimaryKey{Id: studentId})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	jurnal, err := h.services.Jurnal().GetByID(context.Background(), &schedule_service.JurnalPrimaryKey{Id: student.GroupId})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	schedules, err := h.services.Schedule().GetList(context.Background(), &schedule_service.GetListScheduleRequest{
		Limit:  2000,
		Search: jurnal.Id,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	for _, schedules := range schedules.Schedules {
		lesson, err := h.services.Lesson().GetByID(context.Background(), &schedule_service.LessonPrimaryKey{Id: schedules.Id})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		schedules.Lesson = lesson
		tasks, err := h.services.Task().GetList(context.Background(), &schedule_service.GetListTaskRequest{
			Limit:  2000,
			Search: lesson.Id,
		})
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		lesson.Tasks = tasks.Tasks
	}
	scores, err := h.services.Score().GetList(context.Background(), &schedule_service.GetListScoreRequest{
		Limit:  2000,
		Search: studentId,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	events, err := h.services.AssignStudent().GetList(context.Background(), &schedule_service.GetListAssignStudentRequest{
		Limit:  2000,
		Search: studentId,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Scores = scores.Scores
	resp.EventsRegistration = events.AssignStudents
	resp.Schedules = schedules.Schedules

	h.handleResponse(c, http.OK, &resp)

}
