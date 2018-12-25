package data

import "time"

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")

	for rows.Next() {
		conv := Thread{}
		err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
		 if err != nil {
		 	return
		 }
		threads = append(threads, conv)
	}
	rows.Close()
	return
}