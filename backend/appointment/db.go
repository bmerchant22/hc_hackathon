package appointment

import "github.com/gin-gonic/gin"

func createAppointment(ctx *gin.Context, appointment *Appointment) error{
	tx := db.WithContext(ctx).Create(appointment)
	return tx.Error
}

func fetchAppointments(ctx *gin.Context, appointments *[]Appointment) error {
	tx := db.WithContext(ctx).Where("status_code = ?", 1).Find(appointments)
	return tx.Error
}