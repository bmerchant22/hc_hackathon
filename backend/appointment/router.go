package appointment

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	appointment := r.Group("/api/appointment")
	{
		appointment.POST("", CreateAppointmentHandler)
	}
}
