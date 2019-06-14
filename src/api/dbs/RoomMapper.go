package dbs

import (
	"api/models"
	. "core"
)

func RoomFindById(id int) models.Room {
	var room models.Room
	DB.Find(&room, id) //select * form t_room where id = 1
	return room
}
