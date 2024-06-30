package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @ID create_Branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description  Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body user_service.CreateBranch true "CreateBranchBody"
// @Success 200 {object} http.Response{data=user_service.Branch} "GetBranchBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {
	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	var Branch user_service.CreateBranch

	err := c.ShouldBindJSON(&Branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Branch().Create(
		c.Request.Context(),
		&Branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetBranchByID godoc
// @ID get_Branch_by_id
// @Router /branch/{id} [GET]
// @Summary Get Branch  By ID
// @Description Get Branch  By ID
// @Tags Branch
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=user_service.Branch} "BranchBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBranchByID(c *gin.Context) {

	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	BranchID := c.Param("id")

	if !util.IsValidUUID(BranchID) {
		h.handleResponse(c, http.InvalidArgument, "Branch id is an invalid uuid")
		return
	}

	resp, err := h.services.Branch().GetByID(
		context.Background(),
		&user_service.BranchPrimaryKey{
			Id: BranchID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetBranchList godoc
// @ID get_Branch_list
// @Router /branch [GET]
// @Summary Get Branchs List
// @Description  Get Branchs List
// @Tags Branch
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListBranchResponse} "GetAllBranchResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBranchList(c *gin.Context) {
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

	resp, err := h.services.Branch().GetList(
		context.Background(),
		&user_service.GetListBranchRequest{
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

// UpdateBranch godoc
// @ID update_Branch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body user_service.UpdateBranch true "UpdateBranch"
// @Success 200 {object} http.Response{data=user_service.Branch} "Branch data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBranch(c *gin.Context) {

	var Branch user_service.UpdateBranch

	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	Branch.Id = c.Param("id")

	if !util.IsValidUUID(Branch.Id) {
		h.handleResponse(c, http.InvalidArgument, "Branch id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Branch().Update(
		c.Request.Context(),
		&Branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBranch godoc
// @ID delete_Branch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Branch data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBranch(c *gin.Context) {

	BranchId := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "c6b9cac8-ecf1-4b99-b8a9-571daed55fba" {
		h.handleResponse(c, http.BadRequest, "you are not super admin")
		return
	}
	if !util.IsValidUUID(BranchId) {
		h.handleResponse(c, http.InvalidArgument, "Branch id is an invalid uuid")
		return
	}

	resp, err := h.services.Branch().Delete(
		c.Request.Context(),
		&user_service.BranchPrimaryKey{Id: BranchId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
