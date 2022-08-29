package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
	w "videochat-project/pkg/webrtc"
)

// RoomCreate /* Creating a room and redirecting to that room */
func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

/*Room */
func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}
	suuid,_uuid,_:= CreateOrGetRoom(uuid)
}
func RoomWebSocket(c *websocket.Conn){
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
	_,_,room:= CreateOrGetRoom(uuid)
	w.RoomConn(c,room.Peers)
}


func CreateOrGetRoom(uuid string) (string,string, *w.Room) {

}


func RoomViewerSocket(c *websocket.Conn){
	
}

func RoomViewerConn(c *websocket.Conn, p *w.Peers){
	
}

type websocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}