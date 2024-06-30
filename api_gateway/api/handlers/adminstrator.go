package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateAdministrator godoc
// @ID create_Administrator
// @Router /administrator [POST]
// @Summary Create Administrator
// @Description  Create Administrator
// @Tags Administrator
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateAdministrator true "CreateAdministratorBody"
// @Success 200 {object} http.Response{data=user_service.Administrator} "GetAdministratorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdministrator(c *gin.Context) {

	var Administrator user_service.CreateAdministrator
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	err := c.ShouldBindJSON(&Administrator)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Administrator().Create(
		c.Request.Context(),
		&Administrator,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdministratorByID godoc
// @ID get_Administrator_by_id
// @Router /administrator/{id} [GET]
// @Summary Get Administrator  By ID
// @Description Get Administrator  By ID
// @Tags Administrator
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=user_service.Administrator} "AdministratorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministratorByID(c *gin.Context) {

	AdministratorID := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(AdministratorID) {
		h.handleResponse(c, http.InvalidArgument, "Administrator id is an invalid uuid")
		return
	}

	resp, err := h.services.Administrator().GetByID(
		context.Background(),
		&user_service.AdministratorPrimaryKey{
			Id: AdministratorID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetAdministratorList godoc
// @ID get_Administrator_list
// @Router /administrator [GET]
// @Summary Get Administrators List
// @Description  Get Administrators List
// @Tags Administrator
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListAdministratorResponse} "GetAllAdministratorResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministratorList(c *gin.Context) {

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

	resp, err := h.services.Administrator().GetList(
		context.Background(),
		&user_service.GetListAdministratorRequest{
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

// UpdateAdministrator godoc
// @ID update_Administrator
// @Router /administrator/{id} [PUT]
// @Summary Update Administrator
// @Description Update Administrator
// @Tags Administrator
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateAdministrator true "UpdateAdministrator"
// @Success 200 {object} http.Response{data=user_service.Administrator} "Administrator data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdministrator(c *gin.Context) {

	var Administrator user_service.UpdateAdministrator
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	Administrator.Id = c.Param("id")

	if !util.IsValidUUID(Administrator.Id) {
		h.handleResponse(c, http.InvalidArgument, "Administrator id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Administrator)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Administrator().Update(
		c.Request.Context(),
		&Administrator,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdministrator godoc
// @ID delete_Administrator
// @Router /administrator/{id} [DELETE]
// @Summary Delete Administrator
// @Description Delete Administrator
// @Tags Administrator
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Administrator data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdministrator(c *gin.Context) {

	AdministratorId := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(AdministratorId) {
		h.handleResponse(c, http.InvalidArgument, "Administrator id is an invalid uuid")
		return
	}

	resp, err := h.services.Administrator().Delete(
		c.Request.Context(),
		&user_service.AdministratorPrimaryKey{Id: AdministratorId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AdministratorPanel godoc
// @ID adminstrator_panel
// @Router /adminstrator-panel/{id} [GET]
// @Summary adminstrator_panel
// @Description Panel Administrator
// @Tags Administrator
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} http.Response{data=user_service.AdministratorPanelResponse} "AdministratorPanelResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AdministratorPanel(c *gin.Context) {

	admin_id := c.Query("id")

	var resp user_service.AdministratorPanelResponse

	adminstrator, err := h.services.Administrator().GetByID(context.Background(), &user_service.AdministratorPrimaryKey{
		Id: admin_id,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	support_teacher, err := h.services.SupportTeacher().GetList(context.Background(), &user_service.GetListSupportTeacherRequest{
		Limit:  2000,
		Search: adminstrator.BranchId,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	students, err := h.services.Student().GetList(context.Background(), &user_service.GetListStudentRequest{
		Limit:  2000,
		Search: adminstrator.BranchId,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	payments, err := h.services.Payment().GetList(context.Background(), &user_service.GetListPaymentRequest{
		Limit:  2000,
		Search: adminstrator.BranchId,
	})
	fmt.Println(payments, "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>")

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.SupportTeachers = support_teacher.SupportTeachers
	resp.Students = students.Students
	resp.Payments = payments.Payments

	h.handleResponse(c, http.OK, &resp)

}
