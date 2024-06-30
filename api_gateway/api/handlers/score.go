package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateScore godoc
// @ID create_score
// @Router /score [POST]
// @Summary Create Score
// @Description  Create Score
// @Tags Score
// @Accept json
// @Produce json
// @Param profile body schedule_service.CreateScore true "CreateScoreBody"
// @Success 200 {object} http.Response{data=schedule_service.Score} "GetScoreBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateScore(c *gin.Context) {

	var Score schedule_service.CreateScore

	err := c.ShouldBindJSON(&Score)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Score().Create(
		c.Request.Context(),
		&Score,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetScoreByID godoc
// @ID get_score_by_id
// @Router /score/{id} [GET]
// @Summary Get Score  By ID
// @Description Get Score  By ID
// @Tags Score
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=schedule_service.Score} "ScoreBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScoreByID(c *gin.Context) {

	ScoreID := c.Param("id")

	if !util.IsValidUUID(ScoreID) {
		h.handleResponse(c, http.InvalidArgument, "Score id is an invalid uuid")
		return
	}

	resp, err := h.services.Score().GetByID(
		context.Background(),
		&schedule_service.ScorePrimaryKey{
			Id: ScoreID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetScoreList godoc
// @ID get_score_list
// @Router /score [GET]
// @Summary Get Scores List
// @Description  Get Scores List
// @Tags Score
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListScoreResponse} "GetAllScoreResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScoreList(c *gin.Context) {

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

	resp, err := h.services.Score().GetList(
		context.Background(),
		&schedule_service.GetListScoreRequest{
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

// UpdateScore godoc
// @ID update_score
// @Router /score/{id} [PUT]
// @Summary Update Score
// @Description Update Score
// @Tags Score
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body schedule_service.UpdateScore true "UpdateScore"
// @Success 200 {object} http.Response{data=schedule_service.Score} "Score data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateScore(c *gin.Context) {

	var Score schedule_service.UpdateScore

	Score.Id = c.Param("id")

	if !util.IsValidUUID(Score.Id) {
		h.handleResponse(c, http.InvalidArgument, "Score id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Score)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Score().Update(
		c.Request.Context(),
		&Score,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteScore godoc
// @ID delete_score
// @Router /score/{id} [DELETE]
// @Summary Delete Score
// @Description Delete Score
// @Tags Score
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Score data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteScore(c *gin.Context) {

	ScoreId := c.Param("id")

	if !util.IsValidUUID(ScoreId) {
		h.handleResponse(c, http.InvalidArgument, "Score id is an invalid uuid")
		return
	}

	resp, err := h.services.Score().Delete(
		c.Request.Context(),
		&schedule_service.ScorePrimaryKey{Id: ScoreId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
