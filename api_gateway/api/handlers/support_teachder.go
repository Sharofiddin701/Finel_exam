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
// @ID create_SupportTeacher
// @Router /support-teacher [POST]
// @Summary Create SupportTeacher
// @Description  Create SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateTeacher true "CreateTeacherBody"
// @Success 200 {object} http.Response{data=user_service.SupportTeacher} "GetTeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSupportTeacher(c *gin.Context) {

	var SupportTeacher user_service.CreateSupportTeacher

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	err := c.ShouldBindJSON(&SupportTeacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SupportTeacher().Create(
		c.Request.Context(),
		&SupportTeacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetTeacherByID godoc
// @ID get_SupportTeacher_by_id
// @Router /support-teacher/{id} [GET]
// @Summary Get SupportTeacher  By ID
// @Description Get SupportTeacher  By ID
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=user_service.SupportTeacher} "TeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSupportTeacherByID(c *gin.Context) {

	TeacherID := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(TeacherID) {
		h.handleResponse(c, http.InvalidArgument, "SupportTeacher id is an invalid uuid")
		return
	}

	resp, err := h.services.SupportTeacher().GetByID(
		context.Background(),
		&user_service.SupportTeacherPrimaryKey{
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
// @ID get_Teacher_list
// @Router /support-teacher [GET]
// @Summary Get Teachers List
// @Description  Get Teachers List
// @Tags SupportTeacher
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
func (h *Handler) GetSupportTeacherList(c *gin.Context) {

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

	resp, err := h.services.SupportTeacher().GetList(
		context.Background(),
		&user_service.GetListSupportTeacherRequest{
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
// @ID update_Teacher
// @Router /support-teacher/{id} [PUT]
// @Summary Update SupportTeacher
// @Description Update SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateTeacher true "UpdateTeacher"
// @Success 200 {object} http.Response{data=user_service.SupportTeacher} "SupportTeacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSupportTeacher(c *gin.Context) {

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	var SupportTeacher user_service.UpdateSupportTeacher

	SupportTeacher.Id = c.Param("id")

	if !util.IsValidUUID(SupportTeacher.Id) {
		h.handleResponse(c, http.InvalidArgument, "SupportTeacher id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&SupportTeacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SupportTeacher().Update(
		c.Request.Context(),
		&SupportTeacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteTeacher godoc
// @ID delete_teacher
// @Router /support-teacher/{id} [DELETE]
// @Summary Delete SupportTeacher
// @Description Delete SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "SupportTeacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSupportTeacher(c *gin.Context) {

	TeacherId := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(TeacherId) {
		h.handleResponse(c, http.InvalidArgument, "SupportTeacher id is an invalid uuid")
		return
	}

	resp, err := h.services.SupportTeacher().Delete(
		c.Request.Context(),
		&user_service.SupportTeacherPrimaryKey{Id: TeacherId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// SupportTeacherPanel godoc
// @ID support_teacher_panel
// @Router /support-teacher-panel/{support_teacher_id} [GET]
// @Summary support_teacher_panel
// @Description Panel Teacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param support_teacher_id query string true "support_teacher_id"
// @Success 200 {object} http.Response{data=models.SupportTeacherPanelResponse} "SupportTeacherPanelResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SupportTeacherPanel(c *gin.Context) {

	TeacherId := c.Query("support_teacher_id")
	var resp models.SupportTeacherPanelResponse
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
