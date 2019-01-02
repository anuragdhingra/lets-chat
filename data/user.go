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
