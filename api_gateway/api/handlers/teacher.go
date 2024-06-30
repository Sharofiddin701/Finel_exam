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

// CreateTeacher godoc
// @ID create_Teacher
// @Router /teacher [POST]
// @Summary Create Teacher
// @Description  Create Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateTeacher true "CreateTeacherBody"
// @Success 200 {object} http.Response{data=user_service.Teacher} "GetTeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateTeacher(c *gin.Context) {

	var Teacher user_service.CreateTeacher
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	err := c.ShouldBindJSON(&Teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Teacher().Create(
		c.Request.Context(),
		&Teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetTeacherByID godoc
// @ID get_Teacher_by_id
// @Router /teacher/{id} [GET]
// @Summary Get Teacher  By ID
// @Description Get Teacher  By ID
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=user_service.Teacher} "TeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacherByID(c *gin.Context) {

	TeacherID := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(TeacherID) {
		h.handleResponse(c, http.InvalidArgument, "Teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.Teacher().GetByID(
		context.Background(),
		&user_service.TeacherPrimaryKey{
			Id: TeacherID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetTeacherList godoc
// @ID get_teacher_list
// @Router /teacher [GET]
// @Summary Get Teachers List
// @Description  Get Teachers List
// @Tags Teacher
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListTeacherResponse} "GetAllTeacherResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacherList(c *gin.Context) {

	if c.GetHeader("role_id") == config.RoleClient {
		h.handleResponse(c, http.OK, struct{}{})
		return
	}
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
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

	resp, err := h.services.Teacher().GetList(
		context.Background(),
		&user_service.GetListTeacherRequest{
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

// UpdateTeacher godoc
// @ID update_teacher
// @Router /teacher/{id} [PUT]
// @Summary Update Teacher
// @Description Update Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateTeacher true "UpdateTeacher"
// @Success 200 {object} http.Response{data=user_service.Teacher} "Teacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateTeacher(c *gin.Context) {

	var Teacher user_service.UpdateTeacher

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	Teacher.Id = c.Param("id")

	if !util.IsValidUUID(Teacher.Id) {
		h.handleResponse(c, http.InvalidArgument, "Teacher id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Teacher().Update(
		c.Request.Context(),
		&Teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteTeacher godoc
// @ID delete_Teacher
// @Router /teacher/{id} [DELETE]
// @Summary Delete Teacher
// @Description Delete Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Teacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteTeacher(c *gin.Context) {

	TeacherId := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(TeacherId) {
		h.handleResponse(c, http.InvalidArgument, "Teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.Teacher().Delete(
		c.Request.Context(),
		&user_service.TeacherPrimaryKey{Id: TeacherId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// TeacherPanel godoc
// @ID teacher_panel
// @Router /teacherpanel/{teacher_id} [GET]
// @Summary teacher_panel
// @Description Panel Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param teacher_id query string true "teacher_id"
// @Success 200 {object} http.Response{data=models.TeacherPanelResponse} "TeacherPanelResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) TeacherPanel(c *gin.Context) {

	TeacherId := c.Query("teacher_id")
	var resp models.TeacherPanelResponse
	groups, err := h.services.Group().GetList(context.Background(), &schedule_service.GetListGroupRequest{
		Limit:  20000,
		Search: TeacherId,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	for _, group := range groups.Groups {
		students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
			Limit:  200000,
			Search: group.Id,
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
		group.Jurnal = jurnal
		group.Students = studentDs
		jurnal.Students = studentDs
	}
	resp.Groups = groups.Groups

	h.handleResponse(c, http.OK, resp.Groups)

}
