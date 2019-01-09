package data

import (
	"github.com/nu7hatch/gouuid"
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

type CreateThreadRequest struct {
	Topic string
	UserId int
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

func (thread CreateThreadRequest)Create() (threadId int, err error) {
	stmt, err := Db.Prepare("INSERT INTO threads(uuid, topic, user_id, created_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Print(err)
		return
	}
	u4, err := uuid.NewV4()
	_, err = stmt.Exec(u4.String(), thread.Topic, thread.UserId, time.Now())
	if err != nil {
		log.Print(err)
		return
	} else {
		log.Print("Successfully created the thread")
		err = Db.QueryRow("SELECT id FROM threads WHERE uuid=?", u4.String()).Scan(&threadId)
		return
	}
}