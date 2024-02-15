// web socket
package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/bmerchant22/hc_hackathon/auth"
)

func InitializeWS(c *gin.Context, u *auth.User) error {
	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}

	u.WSConnMux.Lock()
	defer u.WSConnMux.Unlock()

	if u.WSConn != nil {
		u.WSConn.Close()
	}

	u.WSConn = conn
	return nil
}

func CleanupWS(u *auth.User) {
	u.WSConnMux.Lock()
	defer u.WSConnMux.Unlock()

	if u.WSConn != nil {
		u.WSConn.Close()
	}
	u.WSConn = nil
}
