package data

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"time"
)

type Post struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

type PostRequest struct{
	Body string
	UserId int
	ThreadId int
}

func (post PostRequest) CreatePost() (createdPost Post, err error) {
	stmt, err := Db.Prepare("INSERT INTO posts(uuid, body, user_id, thread_id, created_at) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Print(err)
		return
	}

		u4, _ := uuid.NewV4()
		_, err = stmt.Exec(u4.String(),post.Body, post.UserId, post.ThreadId, time.Now())
		if err != nil {
			log.Print(err)
			return
		} else {
			err = Db.QueryRow("SELECT * FROM posts WHERE uuid=?", u4.String()).Scan(
				&createdPost.Id, &createdPost.Uuid, &createdPost.Body, &createdPost.UserId, &createdPost.ThreadId, &createdPost.CreatedAt)
			if err != nil {
				log.Print(err)
				return
			}
			return
		}

}

func PostsByThreadId(threadId int) (posts []Post, err error) {
	rows, err := Db.Query("SELECT * FROM posts WHERE thread_id=? ORDER BY created_at ASC", threadId)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		conv := Post{}
		err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Body, &conv.UserId,&conv.ThreadId, &conv.CreatedAt)
		if err != nil {
			log.Print(err)
			return
		}
		posts = append(posts, conv)
	}
	rows.Close()
	return
}