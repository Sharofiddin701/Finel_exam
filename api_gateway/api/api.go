package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api_gateway/api/docs"
	"api_gateway/api/handlers"
	"api_gateway/config"
)

// New
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(5000))
	// r.Use(h.CheckPasswordMiddleware())

	// r.POST("/login", h.Login)

	// USER

	r.POST("/super-admin", h.CreateSuperAdmin)
	r.GET("/super-admin/:id", h.GetSuperAdminByID)
	r.GET("/super-admin", h.GetSuperAdminList)
	r.PUT("/super-admin/:id", h.UpdateSuperAdmin)
	r.DELETE("/super-admin/:id", h.DeleteSuperAdmin)

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranchByID)
	r.GET("/branch", h.GetBranchList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/manager", h.CreateManager)
	r.GET("/manager/:id", h.GetManagerByID)
	r.GET("/manager", h.GetManagerList)
	r.PUT("/manager/:id", h.UpdateManager)
	r.DELETE("/manager/:id", h.DeleteManager)

	r.POST("/teacher", h.CreateTeacher)
	r.GET("/teacher/:id", h.GetTeacherByID)
	r.GET("/teacher", h.GetTeacherList)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.DELETE("/teacher/:id", h.DeleteTeacher)
	r.GET("/teacherpanel/:teacher_id", h.TeacherPanel)
	r.GET("/scheduleweek/:teacher_id", h.GetScheduleReportTeacher)
	r.GET("/schedulemonth/:teacher_id", h.GetScheduleMonthReportTeacher)

	r.POST("/support-teacher", h.CreateSupportTeacher)
	r.GET("/support-teacher/:id", h.GetSupportTeacherByID)
	r.GET("/support-teacher", h.GetSupportTeacherList)
	r.PUT("/support-teacher/:id", h.UpdateSupportTeacher)
	r.DELETE("/support-teacher/:id", h.DeleteSupportTeacher)
	r.GET("/support-teacher-panel/:support_teacher_id", h.SupportTeacherPanel)
	r.GET("/schedule_week/:support_teacher_id", h.GetScheduleReportSteacher)
	r.GET("/schedule_month/:support_teacher_id", h.GetScheduleMonthReportSTeacher)

	r.POST("/administrator", h.CreateAdministrator)
	r.GET("/administrator/:id", h.GetAdministratorByID)
	r.GET("/administrator", h.GetAdministratorList)
	r.PUT("/administrator/:id", h.UpdateAdministrator)
	r.DELETE("/administrator/:id", h.DeleteAdministrator)
	r.GET("/adminstrator-panel/:id", h.AdministratorPanel)

	r.POST("/student", h.CreateStudent)
	r.GET("/student/:id", h.GetStudentByID)
	r.GET("/student", h.GetStudentList)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)
	r.GET("/student-panel/:id", h.StudentPanel)

	r.POST("/group", h.CreateGroup)
	r.GET("/group/:id", h.GetGroupByID)
	r.GET("/group", h.GetGroupList)
	r.PUT("/group/:id", h.UpdateGroup)
	r.DELETE("/group/:id", h.DeleteGroup)

	r.POST("/event", h.CreateEvent)
	r.GET("/event/:id", h.GetEventByID)
	r.GET("/event", h.GetEventList)
	r.PUT("/event/:id", h.UpdateEvent)
	r.DELETE("/event/:id", h.DeleteEvent)

	r.POST("/assign-student", h.CreateAssignStudent)
	r.GET("/assign-student/:id", h.GetAssignStudentByID)
	r.GET("/assign-student", h.GetAssignStudentList)
	r.PUT("/assign-student/:id", h.UpdateAssignStudent)
	r.DELETE("/assign-student/:id", h.DeleteAssignStudent)

	r.POST("/dotask", h.CreateDoTask)
	r.GET("/dotask/:id", h.GetDoTaskByID)
	r.GET("/dotask", h.GetDoTaskList)
	r.PUT("/dotask/:id", h.UpdateDoTask)
	r.DELETE("/dotask/:id", h.DeleteDoTask)

	r.POST("/score", h.CreateScore)
	r.GET("/score/:id", h.GetScoreByID)
	r.GET("/score", h.GetScoreList)
	r.PUT("/score/:id", h.UpdateScore)
	r.DELETE("/score/:id", h.DeleteScore)

	r.POST("/jurnal", h.CreateJurnal)
	r.GET("/jurnal/:id", h.GetJurnalByID)
	r.GET("/jurnal", h.GetJurnalList)
	r.PUT("/jurnal/:id", h.UpdateJurnal)
	r.DELETE("/jurnal/:id", h.DeleteJurnal)

	r.POST("/task", h.CreateTask)
	r.GET("/task/:id", h.GetTaskByID)
	r.GET("/task", h.GetTaskList)
	r.PUT("/task/:id", h.UpdateTask)
	r.DELETE("/task/:id", h.DeleteTask)

	r.POST("/schedule", h.CreateSchedule)
	r.GET("/schedule/:id", h.GetScheduleByID)
	r.GET("/schedule", h.GetScheduleList)
	r.PUT("/schedule/:id", h.UpdateSchedule)
	r.DELETE("/schedule/:id", h.DeleteSchedule)

	r.POST("/lesson", h.CreateLesson)
	r.GET("/lesson/:id", h.GetLessonByID)
	r.GET("/lesson", h.GetLessonList)
	r.PUT("/lesson/:id", h.UpdateLesson)
	r.DELETE("/lesson/:id", h.DeleteLesson)

	r.POST("/payment", h.CreatePayment)
	r.GET("/payment/:id", h.GetPaymentByID)
	r.GET("/payment", h.GetPaymentList)
	r.PUT("/payment/:id", h.UpdatePayment)
	r.DELETE("/payment/:id", h.DeletePayment)

	r.GET("/report-teacher/:branch_id", h.GetTeacherReport)
	r.GET("/report-support-teacher/:branch_id", h.GetSupportTeacherReport)
	r.GET("/report-administrator/:branch_id", h.GetSupportTeacherReport)
	r.GET("/report-student/:branch_id", h.GetStudentReport)
	r.GET("/report-schedule-week/:group_id", h.GetScheduleReport)
	r.GET("/report-schedule-month/:group_id", h.GetScheduleMonthReport)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}
