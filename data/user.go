package data

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"time"
)

type User struct {
	Id int
	Username string
	Uuid string
	Email string
	Password string
	CreatedAt time.Time
}

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (user User) Create() (err error){
	stmt, err := Db.Prepare("INSERT INTO users SET uuid=?, username=?, email=?, password=?, created_at=?")
	if err != nil {
		log.Print(err)
		return err
	}

	u4, err := uuid.NewV4()
	if err != nil {
		log.Print(err)
		return err
	}
	_, err = stmt.Exec(u4.String() ,user.Username, user.Email, user.Password, time.Now())
	if err != nil {
		log.Print(err)
		return err
	} else {
		log.Print("User created successfully!")
		return nil
	}
}

func UserByEmailOrUsername(emailOrUsername string) (user User, err error) {
	conv := User{}
	err = Db.QueryRow("SELECT * FROM users WHERE email=? OR username=?",emailOrUsername,emailOrUsername).Scan(&conv.Id,
		&conv.Uuid, &conv.Username,&conv.Email, &conv.Password, &conv.CreatedAt)
	if err != nil {
		log.Print(err)
		return
	}
	user = conv
	return user, nil
}

func (user User) CreateSession() (session Session, err error) {
	stmt, err := Db.Prepare("INSERT INTO sessions(uuid, email, user_id, created_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Print(err)
		return
	}

	u4, err := uuid.NewV4()
	if err != nil {
		log.Print(err)
		return
	}
	_, err = stmt.Exec(u4.String(), user.Email, user.Id, time.Now())
	if err != nil {
		log.Print(err)
		return
	} else {
		log.Print("Successfully registered the session")
		err = Db.QueryRow("SELECT * FROM sessions WHERE uuid=?",u4.String()).Scan(
			&session.Id,&session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
}

func (user User) Session() (session Session, err error) {
	conv := Session{}
	err = Db.QueryRow("SELECT * FROM sessions WHERE user_id=?",user.Id).Scan(
		&conv.Id,&conv.Uuid, &conv.Email, &conv.UserId, &conv.CreatedAt)
	if err != nil {
		log.Print(err)
		return
	}
	session = conv
	return session, nil
}

func (session Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT * FROM sessions WHERE uuid=?", session.Uuid).Scan(
		&session.Id,&session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		log.Print(err)
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (sess Session) User() (user User, err error) {
	err = Db.QueryRow("SELECT user_id FROM sessions WHERE uuid=?", sess.Uuid).Scan(
		&user.Id)

	err = Db.QueryRow("SELECT * FROM users WHERE id=?", user.Id).Scan(
		&user.Id, &user.Uuid, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Print("User corresponding to the session not found")
		return
	}
	return
}

func (sess Session) DeleteByUUID() (err error) {
	stmt, err := Db.Prepare("DELETE FROM sessions WHERE uuid=?")

	_, err = stmt.Exec(sess.Uuid)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
