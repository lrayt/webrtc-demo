package signaling

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webrtc-demo/utils"
)

type RoomHub struct {
	hub map[string]*Room
}

func NewRoomHub() *RoomHub {
	return &RoomHub{
		hub: make(map[string]*Room),
	}
}

func (h RoomHub) All(c *gin.Context) {
	var data = make([]*RoomInfo, 0, len(h.hub))
	for _, room := range h.hub {
		data = append(data, room.RoomInfo())
	}
	c.JSON(http.StatusOK, &utils.BaseResponse{
		Code: utils.OK,
		Msg:  utils.Success,
		Data: data,
	})
}
func (h RoomHub) Add(c *gin.Context) {
	var roomInfo = new(RoomInfo)
	if err := c.Bind(roomInfo); err != nil || len(roomInfo.Name) <= 0 {
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.ParamInvalid,
			Msg:  "room name is empty",
		})
		return
	}
	room := NewRoom(roomInfo.Name)
	h.hub[room.id] = room
	c.JSON(http.StatusOK, &utils.BaseResponse{
		Code: utils.OK,
		Msg:  utils.Success,
		Data: room.RoomInfo(),
	})
}
func (h RoomHub) Del(c *gin.Context) {
	roomId := c.Param("room_id")
	if len(roomId) <= 0 {
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.ParamInvalid,
			Msg:  "room id is empty",
		})
	}

	if _, ok := h.hub[roomId]; !ok {
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.BusinessError,
			Msg:  "room id is not exist",
		})
		return
	}
	delete(h.hub, roomId)
	c.JSON(http.StatusOK, &utils.BaseResponse{
		Code: utils.OK,
		Msg:  utils.Success,
	})
}

func (h RoomHub) Users(c *gin.Context) {
	roomId, ok := c.GetQuery("room_id")
	if !ok || len(roomId) <= 0 {
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.ParamInvalid,
			Msg:  "room id is empty",
		})
		return
	}

	if room, ok := h.hub[roomId]; !ok {
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.BusinessError,
			Msg:  "room id is not exist",
		})
	} else {
		users := make([]*UserInfo, 0, len(room.clients))
		for _, client := range room.clients {
			users = append(users, client.userInfo())
		}
		c.JSON(http.StatusOK, &utils.BaseResponse{
			Code: utils.OK,
			Msg:  utils.Success,
			Data: users,
		})
	}
}
func (h RoomHub) WS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("target", err)
		return
	}
	client := NewClient(conn)
	roomId, ok := c.GetQuery("roomId")
	if !ok {
		client.send <- &Message{Action: Error, Content: "room id is empty"}
		return
	}

	if room, exist := h.hub[roomId]; !exist {
		client.send <- &Message{Action: Error, Content: "room is not exist"}
		return
	} else {
		client.room = room
	}
	client.room.register <- client
}
