package handlers

import (
	"api_gateway/api/http"
	"api_gateway/api/models"
	"api_gateway/genproto/schedule_service"
	"api_gateway/genproto/user_service"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTeacherReport godoc
// @ID get_teacher_report_by_branch_id
// @Router /report-teacher/{branch_id} [GET]
// @Summary get_teacher_report_by_branch_id
// @Description get_teacher_report_by_branch_id
// @Tags Report
// @Accept json
// @Produce json
// @Param branch_id query string true "branch_id"
// @Success 200 {object} http.Response{data=models.TeachersReport} "TeacherReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacherReport(c *gin.Context) {

	branchId := c.Query("branch_id")

	var teacher_report []models.TeachersReport

	teachers, err := h.services.Teacher().GetList(context.Background(), &user_service.GetListTeacherRequest{
		Search: branchId,
		Limit:  10000,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	var (
		year  string
		month string
		day   string
	)

	for _, teacher := range teachers.Teachers {

		for i := 0; i < 4; i++ {
			year += string(teacher.CreatedAt[i])
		}
		for i := 5; i < 7; i++ {
			month += string(teacher.CreatedAt[i])
		}
		for i := 8; i < 10; i++ {
			day += string(teacher.CreatedAt[i])
		}

		yearint, _ := strconv.Atoi(year)
		monthint, _ := strconv.Atoi(month)

		var monthworked = (time.Now().Year()-yearint)*12 + (int(time.Now().Month()) - monthint)

		teacher_report = append(teacher_report, models.TeachersReport{
			FullName:     teacher.FullName,
			Phone:        teacher.Phone,
			Salary:       teacher.Salary,
			MonthsWorked: int64(monthworked),
			TotalSum:     float64(monthworked) * teacher.Salary,
			IeltsScore:   teacher.IeltsScore,
		})
		monthworked = 0
		monthint = 0
		yearint = 0
		year = ""
		month = ""

	}
	h.handleResponse(c, http.Created, teacher_report)

}

// GetSupportTeacherReport godoc
// @ID get_support_teacher_report_by_branch_id
// @Router /report-support-teacher/{branch_id} [GET]
// @Summary get_support_teacher_report_by_branch_id
// @Description support_get_teacher_report_by_branch_id
// @Tags Report
// @Accept json
// @Produce json
// @Param branch_id query string true "branch_id"
// @Success 200 {object} http.Response{data=models.TeachersReport} "SupportTeacherReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSupportTeacherReport(c *gin.Context) {

	branchId := c.Query("branch_id")

	var teacher_report []models.TeachersReport

	supportteachers, err := h.services.SupportTeacher().GetList(context.Background(), &user_service.GetListSupportTeacherRequest{
		Search: branchId,
		Limit:  10000,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	var (
		year  string
		month string
		day   string
	)

	for _, teacher := range supportteachers.SupportTeachers {

		for i := 0; i < 4; i++ {
			year += string(teacher.CreatedAt[i])
		}
		for i := 5; i < 7; i++ {
			month += string(teacher.CreatedAt[i])
		}
		for i := 8; i < 10; i++ {
			day += string(teacher.CreatedAt[i])
		}

		yearint, _ := strconv.Atoi(year)
		monthint, _ := strconv.Atoi(month)

		var monthworked = (time.Now().Year()-yearint)*12 + (int(time.Now().Month()) - monthint)

		teacher_report = append(teacher_report, models.TeachersReport{
			FullName:     teacher.FullName,
			Phone:        teacher.Phone,
			Salary:       teacher.Salary,
			MonthsWorked: int64(monthworked),
			TotalSum:     float64(monthworked) * teacher.Salary,
			IeltsScore:   teacher.IeltsScore,
		})
		monthworked = 0
		monthint = 0
		yearint = 0
		year = ""
		month = ""

	}
	h.handleResponse(c, http.Created, teacher_report)

}

// GetAdministratorReport godoc
// @ID get_administrator_report_by_branch_id
// @Router /report-administrator/{branch_id} [GET]
// @Summary get_administrator_report_by_branch_id
// @Description administrator_report_by_branch_id
// @Tags Report
// @Accept json
// @Produce json
// @Param branch_id query string true "branch_id"
// @Success 200 {object} http.Response{data=models.TeachersReport} "AdministratorReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministratorReport(c *gin.Context) {

	branchId := c.Query("branch_id")

	var admins []models.TeachersReport

	administrators, err := h.services.Administrator().GetList(context.Background(), &user_service.GetListAdministratorRequest{
		Search: branchId,
		Limit:  10000,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	var (
		year  string
		month string
		day   string
	)

	for _, administrator := range administrators.Administrators {

		for i := 0; i < 4; i++ {
			year += string(administrator.CreatedAt[i])
		}
		for i := 5; i < 7; i++ {
			month += string(administrator.CreatedAt[i])
		}
		for i := 8; i < 10; i++ {
			day += string(administrator.CreatedAt[i])
		}

		yearint, _ := strconv.Atoi(year)
		monthint, _ := strconv.Atoi(month)

		var monthworked = (time.Now().Year()-yearint)*12 + (int(time.Now().Month()) - monthint)

		admins = append(admins, models.TeachersReport{
			FullName:     administrator.FullName,
			Phone:        administrator.Phone,
			Salary:       administrator.Salary,
			MonthsWorked: int64(monthworked),
			TotalSum:     float64(monthworked) * administrator.Salary,
			IeltsScore:   administrator.IeltsScore,
		})
		monthworked = 0
		monthint = 0
		yearint = 0
		year = ""
		month = ""

	}
	h.handleResponse(c, http.Created, admins)

}

// GetStudentReport godoc
// @ID get_student_report_by_branch_id
// @Router /report-student/{branch_id} [GET]
// @Summary get_student_report_by_branch_id
// @Description student_report_by_branch_id
// @Tags Report
// @Accept json
// @Produce json
// @Param branch_id query string true "branch_id"
// @Success 200 {object} http.Response{data=user_service.StudentReport} "StudentReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentReport(c *gin.Context) {

	branchId := c.Query("branch_id")
	fmt.Println(branchId)
	resp, err := h.services.Student().GetStudetReport(context.Background(), &user_service.StudentReportRequest{
		BranchId: branchId,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportWeek godoc
// @ID get_schedule_report_by_group_id_week
// @Router /report-schedule-week/{group_id} [GET]
// @Summary get_schedule_report_by_group_id
// @Description schedule_report_by_group_id
// @Tags Report
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleReport(c *gin.Context) {

	branchId := c.Query("group_id")
	fmt.Println(branchId)
	resp, err := h.services.Schedule().GetScheduleReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportMonth godoc
// @ID get_schedule_report_by_group_id_month
// @Router /report-schedule-month/{group_id} [GET]
// @Summary get_schedule_report_by_group_id
// @Description schedule_report_by_group_id
// @Tags Report
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleMonthReport(c *gin.Context) {

	branchId := c.Query("group_id")
	fmt.Println(branchId)
	resp, err := h.services.Schedule().GetScheduleMonthReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportMonth godoc
// @ID get_schedule_report_by_teacher_month
// @Router /schedulemonth/{teacher_id} [GET]
// @Summary get_schedule_report_by_teacher
// @Description schedule_report_by_teacher
// @Tags Report
// @Accept json
// @Produce json
// @Param teacher_id query string true "teacher_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleMonthReportTeacher(c *gin.Context) {

	branchId := c.Query("teacher_id")
	fmt.Println(branchId)
	resp, err := h.services.Schedule().GetScheduleMonthReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportMonth godoc
// @ID get_schedule_report_by_support_teacher_id_month
// @Router /schedule_month/{support_teacher_id} [GET]
// @Summary get_schedule_report_by_teacher
// @Description schedule_report_by_teacher
// @Tags Report
// @Accept json
// @Produce json
// @Param support_teacher_id query string true "support_teacher_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleMonthReportSTeacher(c *gin.Context) {

	branchId := c.Query("support_teacher_id")
	fmt.Println(branchId)
	resp, err := h.services.Schedule().GetScheduleMonthReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportWeek godoc
// @ID get_schedule_report_by_teacher_id_week
// @Router /scheduleweek/{teacher_id} [GET]
// @Summary get_schedule_report_by_teacher_id
// @Description schedule_report_by_teacher_id
// @Tags Report
// @Accept json
// @Produce json
// @Param teacher_id query string true "teacher_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleReportTeacher(c *gin.Context) {

	branchId := c.Query("teacher_id")
	fmt.Println(branchId)
	resp, err := h.services.Schedule().GetScheduleReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}

// GetScheduleReportWeek godoc
// @ID get_schedule_report_by_support_teacher_id_week
// @Router /schedule_week/{support_teacher_id} [GET]
// @Summary get_schedule_report_by_support_teacher_id
// @Description schedule_report_by_support_teacher_id
// @Tags Report
// @Accept json
// @Produce json
// @Param support_teacher_id query string true "support_teacher_id"
// @Success 200 {object} http.Response{data=schedule_service.ScheduleReport} "ScheduleReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScheduleReportSteacher(c *gin.Context) {

	branchId := c.Query("support_teacher_id")

	resp, err := h.services.Schedule().GetScheduleReport(context.Background(), &schedule_service.ScheduleReportRequest{
		BranchId: branchId,
	})

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)

}
