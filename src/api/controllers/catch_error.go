package controllers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	"github.com/gin-gonic/gin"
)

func CatchError(c *gin.Context, err error) {
	if err != nil {
		appErr := apperrors.Parse(err)
		c.JSON(appErr.Code(), gin.H{
			"error": appErr.Error(),
			"code":  appErr.Code(),
		})
	}
}
