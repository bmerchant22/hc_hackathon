package appointment

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func fetchAppointmentsHandler(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer ws.Close()
	var appointments []Appointment
	err = fetchAppointments(ctx, &appointments)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch appointments for the first time
	for _, appointment := range appointments {
		if err := ws.WriteJSON(appointment); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	for {
		// Fetch new appointments
		var appointments []Appointment
		if err := fetchAppointments(ctx, &appointments); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Send appointments over WebSocket
		for _, appointment := range appointments {
			if err := ws.WriteJSON(appointment); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// Sleep for a while before fetching appointments again
		time.Sleep(5 * time.Second)
	}
}
