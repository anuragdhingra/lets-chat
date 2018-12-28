package data

import (
	"log"
	"strconv"
	"time"
)

type Thread struct {
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Topic string `json:"topic"`
	UserId int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		conv := Thread{}
		err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
		 if err != nil {
			 log.Print(err)
		 	return
		 }
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func ThreadByID(threadId string) (thread Thread, err error) {
	id, err := strconv.Atoi(threadId)
	if err != nil {
		log.Print(err)
		return
	}

	conv := Thread{}
	err = Db.QueryRow("SELECT * FROM threads WHERE id = ?",id).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return
	}
	thread = conv
	return
}