package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateManager godoc
// @ID create_Manager
// @Router /manager [POST]
// @Summary Create Manager
// @Description  Create Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateManager true "CreateManagerBody"
// @Success 200 {object} http.Response{data=user_service.Manager} "GetManagerBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateManager(c *gin.Context) {

	var Manager user_service.CreateManager
	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	err := c.ShouldBindJSON(&Manager)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Manager().Create(
		c.Request.Context(),
		&Manager,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetManagerByID godoc
// @ID get_Manager_by_id
// @Router /manager/{id} [GET]
// @Summary Get Manager  By ID
// @Description Get Manager  By ID
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=user_service.Manager} "ManagerBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetManagerByID(c *gin.Context) {

	ManagerID := c.Param("id")
	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	if !util.IsValidUUID(ManagerID) {
		h.handleResponse(c, http.InvalidArgument, "Manager id is an invalid uuid")
		return
	}

	resp, err := h.services.Manager().GetByID(
		context.Background(),
		&user_service.ManagerPrimaryKey{
			Id: ManagerID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetManagerList godoc
// @ID get_Manager_list
// @Router /manager [GET]
// @Summary Get Managers List
// @Description  Get Managers List
// @Tags Manager
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListManagerResponse} "GetAllManagerResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetManagerList(c *gin.Context) {
	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
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

	resp, err := h.services.Manager().GetList(
		context.Background(),
		&user_service.GetListManagerRequest{
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

// UpdateManager godoc
// @ID update_Manager
// @Router /manager/{id} [PUT]
// @Summary Update Manager
// @Description Update Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateManager true "UpdateManager"
// @Success 200 {object} http.Response{data=user_service.Manager} "Manager data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateManager(c *gin.Context) {

	var Manager user_service.UpdateManager
	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	Manager.Id = c.Param("id")

	if !util.IsValidUUID(Manager.Id) {
		h.handleResponse(c, http.InvalidArgument, "Manager id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Manager)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Manager().Update(
		c.Request.Context(),
		&Manager,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteManager godoc
// @ID delete_Manager
// @Router /manager/{id} [DELETE]
// @Summary Delete Manager
// @Description Delete Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Manager data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteManager(c *gin.Context) {

	ManagerId := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	if !util.IsValidUUID(ManagerId) {
		h.handleResponse(c, http.InvalidArgument, "Manager id is an invalid uuid")
		return
	}

	resp, err := h.services.Manager().Delete(
		c.Request.Context(),
		&user_service.ManagerPrimaryKey{Id: ManagerId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
