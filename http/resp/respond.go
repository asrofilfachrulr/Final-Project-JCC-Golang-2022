package resp

import (
	webModels "anya-day/models/web"

	"net/http"

	"github.com/gin-gonic/gin"
)

// normal error respond wtih msg message. default error will be Bad Request
func RespErrWithMsg(ctx *gin.Context, msg string, code ...int) {
	statusCode := 0
	if len(code) == 0 {
		statusCode = http.StatusBadRequest
	}
	ctx.JSON(statusCode, webModels.ErrWithMsg(msg))
}
