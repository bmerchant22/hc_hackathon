package appointment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAppointmentHandler(ctx *gin.Context) {
	var appointment *Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createAppointment(ctx, appointment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully"})
}
