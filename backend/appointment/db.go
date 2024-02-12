package appointment

import "github.com/gin-gonic/gin"

func createAppointment(ctx *gin.Context, appointment *Appointment) error{
	tx := db.WithContext(ctx).Create(appointment)
	return tx.Error
}