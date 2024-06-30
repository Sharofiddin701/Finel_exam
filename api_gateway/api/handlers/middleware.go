package handlers

import (
	"errors"
	"net/http"

	"api_gateway/config"

	"github.com/gin-gonic/gin"
)

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type HasAccessModel struct {
	Id     string
	RoleId string
}

// func (h *Handler) BasicAuthMiddleware(c *gin.Context) {

// 	if h.cfg.Environment == config.TestMode {
// 		h.log.Debug("BasicAuthMiddelware passed... testMode")
// 		c.Next()
// 		return
// 	}

// 	var auth HasAccessModel
// 	ok := h.hasAccess(c, &auth)
// 	if !ok {
// 		c.Abort()
// 		return
// 	}

// 	c.Request.Header.Add("user_id", auth.Id)
// 	c.Request.Header.Add("role_id", auth.RoleId)

// 	c.Next()
// }

// // hasAccess ...
// func (h *Handler) hasAccess(c *gin.Context, result *HasAccessModel) (ok bool) {

// 	platformID := c.GetHeader("Platform-Id")

// 	fmt.Println(platformID)
// 	if !util.IsValidUUIDV1(platformID) {
// 		h.handleErrorResponse(c, 422, "validation error", "platform-id")
// 		return false
// 	}

// 	bearerToken := c.GetHeader("Authorization")

// 	strArr := strings.Split(bearerToken, " ")

// 	if len(strArr) != 2 {
// 		h.handleErrorResponse(c, 403, "token error", "wrong format")
// 		return false
// 	}

// 	resp, err := h.services.SessionService().HasAccess(c.Request.Context(),
// 		&auth_service.HasAccessRequest{
// 			AccessToken: strArr[1],
// 		},
// 	)

// 	if err != nil {
// 		h.handleErrorResponse(c, 500, err.Error(), "internal server error")
// 		return false
// 	}

// 	result.Id = resp.Id
// 	result.RoleId = resp.RoleId

// 	return true
// }

func (h *Handler) CheckPasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		password := c.GetHeader("API-KEY")
		if password != config.SecureApiKey {
			c.AbortWithError(http.StatusForbidden, errors.New("The request requires an user authentication"))
			return
		}

		c.Next()
	}
}
