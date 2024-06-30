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

// CreateSuperAdmin godoc
// @ID create_SuperAdmin
// @Router /super-admin [POST]
// @Summary Create SuperAdmin
// @Description  Create SuperAdmin
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param profile body user_service.CreateSuperAdmin true "CreateSuperAdminBody"
// @Success 200 {object} http.Response{data=user_service.SuperAdmin} "GetSuperAdminBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSuperAdmin(c *gin.Context) {

	var SuperAdmin user_service.CreateSuperAdmin

	err := c.ShouldBindJSON(&SuperAdmin)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	fmt.Println("superAdmin:", &SuperAdmin)

	resp, err := h.services.CrmService().Create(
		c.Request.Context(),
		&SuperAdmin,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSuperAdminByID godoc
// @ID get_SuperAdmin_by_id
// @Router /super-admin/{id} [GET]
// @Summary Get SuperAdmin  By ID
// @Description Get SuperAdmin  By ID
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=user_service.SuperAdmin} "SuperAdminBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSuperAdminByID(c *gin.Context) {

	SuperAdminID := c.Param("id")

	if !util.IsValidUUID(SuperAdminID) {
		h.handleResponse(c, http.InvalidArgument, "SuperAdmin id is an invalid uuid")
		return
	}

	resp, err := h.services.CrmService().GetByID(
		context.Background(),
		&user_service.SuperAdminPrimaryKey{
			Id: SuperAdminID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetSuperAdminList godoc
// @ID get_SuperAdmin_list
// @Router /super-admin [GET]
// @Summary Get SuperAdmins List
// @Description  Get SuperAdmins List
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListSuperAdminResponse} "GetAllSuperAdminResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSuperAdminList(c *gin.Context) {

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

	resp, err := h.services.CrmService().GetList(
		context.Background(),
		&user_service.GetListSuperAdminRequest{
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

// UpdateSuperAdmin godoc
// @ID update_SuperAdmin
// @Router /super-admin/{id} [PUT]
// @Summary Update SuperAdmin
// @Description Update SuperAdmin
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body user_service.UpdateSuperAdmin true "UpdateSuperAdmin"
// @Success 200 {object} http.Response{data=user_service.SuperAdmin} "SuperAdmin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSuperAdmin(c *gin.Context) {

	var SuperAdmin user_service.UpdateSuperAdmin

	SuperAdmin.Id = c.Param("id")

	if !util.IsValidUUID(SuperAdmin.Id) {
		h.handleResponse(c, http.InvalidArgument, "SuperAdmin id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&SuperAdmin)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CrmService().Update(
		c.Request.Context(),
		&SuperAdmin,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteSuperAdmin godoc
// @ID delete_SuperAdmin
// @Router /super-admin/{id} [DELETE]
// @Summary Delete SuperAdmin
// @Description Delete SuperAdmin
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "SuperAdmin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSuperAdmin(c *gin.Context) {

	SuperAdminId := c.Param("id")

	if !util.IsValidUUID(SuperAdminId) {
		h.handleResponse(c, http.InvalidArgument, "SuperAdmin id is an invalid uuid")
		return
	}

	resp, err := h.services.CrmService().Delete(
		c.Request.Context(),
		&user_service.SuperAdminPrimaryKey{Id: SuperAdminId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
