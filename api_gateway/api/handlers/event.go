package handlers

import (
	"api_gateway/api/http"
	"api_gateway/config"
	"api_gateway/genproto/schedule_service"
	"api_gateway/pkg/util"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @ID create_event
// @Router /event [POST]
// @Summary Create Event
// @Description  Create Event
// @Tags Event
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param profile body schedule_service.CreateEvent true "CreateEventBody"
// @Success 200 {object} http.Response{data=schedule_service.Event} "GetEventBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateEvent(c *gin.Context) {

	var reqevent schedule_service.CreateEvent

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	err := c.ShouldBindJSON(&reqevent)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	events, err := h.services.Event().GetList(context.Background(), &schedule_service.GetListEventRequest{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	var timey string

	for _, event := range events.Events {
		for i := 0; i < 5; i++ {
			timey += string(event.StartTime[i])
		}
		if event.BranchId == reqevent.BranchId && event.Date == reqevent.Date && timey == reqevent.StartTime {
			h.handleResponse(c, http.BadRequest, ">>>>>>>> this time is already reserved for another event <<<<<<<")
			return
		}
		timey = ""
	}

	reqdate := reqevent.Date
	var (
		year  string
		month string
		day   string
	)

	for i := 0; i < 4; i++ {
		year += string(reqdate[i])
	}
	for i := 5; i < 7; i++ {
		month += string(reqdate[i])
	}

	for i := 8; i < len(reqdate); i++ {
		day += string(reqdate[i])
	}

	yearint, _ := strconv.Atoi(year)
	monthint, _ := strconv.Atoi(month)
	dayint, _ := strconv.Atoi(day)

	date := time.Date(yearint, time.Month(monthint), dayint, 00, 00, 00, 00, time.UTC)
	if date.Weekday() != time.Sunday {
		h.handleResponse(c, http.BadRequest, ">>>>>>>>> event date is not sunday,events only sunday <<<<<<<<<")
		return
	}

	resp, err := h.services.Event().Create(
		c.Request.Context(),
		&reqevent,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetEventByID godoc
// @ID get_event_by_id
// @Router /event/{id} [GET]
// @Summary Get Event  By ID
// @Description Get Event  By ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=schedule_service.Event} "EventBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEventByID(c *gin.Context) {

	EventID := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(EventID) {
		h.handleResponse(c, http.InvalidArgument, "Event id is an invalid uuid")
		return
	}

	resp, err := h.services.Event().GetByID(
		context.Background(),
		&schedule_service.EventPrimaryKey{
			Id: EventID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetEventList godoc
// @ID get_event_list
// @Router /event [GET]
// @Summary Get Events List
// @Description  Get Events List
// @Tags Event
// @Accept json
// @Produce json
// @Param role_id query string true "role_id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=schedule_service.GetListEventResponse} "GetAllEventResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEventList(c *gin.Context) {

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

	resp, err := h.services.Event().GetList(
		context.Background(),
		&schedule_service.GetListEventRequest{
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

// UpdateEvent godoc
// @ID update_event
// @Router /event/{id} [PUT]
// @Summary Update Event
// @Description Update Event
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Param profile body schedule_service.UpdateEvent true "UpdateEvent"
// @Success 200 {object} http.Response{data=schedule_service.Event} "Event data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateEvent(c *gin.Context) {

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	var Event schedule_service.UpdateEvent

	Event.Id = c.Param("id")

	if !util.IsValidUUID(Event.Id) {
		h.handleResponse(c, http.InvalidArgument, "Event id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Event)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Event().Update(
		c.Request.Context(),
		&Event,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteEvent godoc
// @ID delete_event
// @Router /event/{id} [DELETE]
// @Summary Delete Event
// @Description Delete Event
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role_id query string true "role_id"
// @Success 200 {object} http.Response{data=object{}} "Event data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteEvent(c *gin.Context) {

	EventId := c.Param("id")

	roleID := c.Query("role_id")
	if roleID != "45344572-d60f-43ae-9d03-4a16e0127e52" {
		h.handleResponse(c, http.BadRequest, "you are not manager")
		return
	}
	if !util.IsValidUUID(EventId) {
		h.handleResponse(c, http.InvalidArgument, "Event id is an invalid uuid")
		return
	}

	resp, err := h.services.Event().Delete(
		c.Request.Context(),
		&schedule_service.EventPrimaryKey{Id: EventId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
