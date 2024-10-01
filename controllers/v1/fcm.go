package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/imbee-backend/constants"
	"github.com/kalleriakronos24/imbee-backend/dto"
	"github.com/kalleriakronos24/imbee-backend/services"
	"github.com/kalleriakronos24/imbee-backend/utils"
)

// AuthLogin godoc
// @Summary      FCM
// @Description  To send FCM Message with RBMQ Queue and Exchanges
// @Tags         FCM
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.InsertFCM true "insert fcm"
// @Router       /fcm/send [post]
func POSTInsertFCM(c *gin.Context) {
	var err error
	var p dto.InsertFCM
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err = services.Handler.InsertFcmJob(&p); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Failed to send FCM message. Try Again Later"))
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: nil})
}
