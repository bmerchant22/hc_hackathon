package ws

import (
	"encoding/json"
	"errors"

	"github.com/bmerchant22/hc_hackathon/auth"
	"github.com/gorilla/websocket"
)

func Notify(message interface{}, u *auth.User) error {
	u.WSConnMux.Lock()
	defer u.WSConnMux.Unlock()

	if u.WSConn == nil {
		return errors.New("WebSocket connection is not initialized")
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = u.WSConn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		return err
	}

	return nil
}
