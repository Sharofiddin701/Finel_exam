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

// CreatePayment godoc
// @ID create_payment
// @Router /payment [POST]
// @Summary Create Payment
// @Description  Create Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param profile body user_service.CreatePayment true "CreatePaymentBody"
// @Success 200 {object} http.Response{data=user_service.Payment} "GetPaymentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePayment(c *gin.Context) {

	var Payment user_service.CreatePayment

	err := c.ShouldBindJSON(&Payment)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	student_payment, err := h.services.Payment().GetByID(context.Background(), &user_service.PaymentPrimaryKey{Id: Payment.StudentId})
	fmt.Println(student_payment)
	if err != nil {

		resp, err := h.services.Payment().Create(
			c.Request.Context(),
			&Payment,
		)
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}

		h.handleResponse(c, http.Created, resp)
	} else {

		resp, err := h.services.Payment().Update(context.Background(), &user_service.UpdatePayment{
			Id:          student_payment.Id,
			StudentId:   student_payment.StudentId,
			BranchId:    student_payment.BranchId,
			PaidSum:     student_payment.PaidSum + Payment.PaidSum,
			TotalSum:    student_payment.TotalSum,
			CourseCount: student_payment.CourseCount,
		})

		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
		fmt.Println("ok1")

		h.handleResponse(c, http.OK, resp)
	}

}

// GetPaymentByID godoc
// @ID get_payment_by_id
// @Router /payment/{id} [GET]
// @Summary Get Payment  By ID
// @Description Get Payment  By ID
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=user_service.Payment} "PaymentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPaymentByID(c *gin.Context) {

	PaymentID := c.Param("id")

	if !util.IsValidUUID(PaymentID) {
		h.handleResponse(c, http.InvalidArgument, "Payment id is an invalid uuid")
		return
	}

	resp, err := h.services.Payment().GetByID(
		context.Background(),
		&user_service.PaymentPrimaryKey{
			Id: PaymentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetPaymentList godoc
// @ID get_payment_list
// @Router /payment [GET]
// @Summary Get Payments List
// @Description  Get Payments List
// @Tags Payment
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=user_service.GetListPaymentResponse} "GetAllPaymentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPaymentList(c *gin.Context) {

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

	resp, err := h.services.Payment().GetList(
		context.Background(),
		&user_service.GetListPaymentRequest{
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

// UpdatePayment godoc
// @ID update_payment
// @Router /payment/{id} [PUT]
// @Summary Update Payment
// @Description Update Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body user_service.UpdatePayment true "UpdatePayment"
// @Success 200 {object} http.Response{data=user_service.Payment} "Payment data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePayment(c *gin.Context) {

	var Payment user_service.UpdatePayment

	Payment.Id = c.Param("id")

	if !util.IsValidUUID(Payment.Id) {
		h.handleResponse(c, http.InvalidArgument, "Payment id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Payment)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Payment().Update(
		c.Request.Context(),
		&Payment,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeletePayment godoc
// @ID delete_payment
// @Router /payment/{id} [DELETE]
// @Summary Delete Payment
// @Description Delete Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Payment data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePayment(c *gin.Context) {

	PaymentId := c.Param("id")

	if !util.IsValidUUID(PaymentId) {
		h.handleResponse(c, http.InvalidArgument, "Payment id is an invalid uuid")
		return
	}

	resp, err := h.services.Payment().Delete(
		c.Request.Context(),
		&user_service.PaymentPrimaryKey{Id: PaymentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
