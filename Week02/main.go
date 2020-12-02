package main

import (
	"database/sql"
	"log"
)
import "github.com/pkg/errors"

type User struct {
	Id uint64
	Name string
}

type Dao struct {}

func (*Dao) GetUserById(id uint64) (*User, error) {
	err := sql.ErrNoRows
	return &User{}, errors.Wrapf(err, "User id: %v not exists", id)
}

func Service(id uint64) (*User, error) {
	d := &Dao{}
	user, err := d.GetUserById(id)
	if err != nil {
		return nil, errors.WithMessage(err, "Get user failed")
	}
	return user, nil
}

func main() {
	var id uint64 = 1024
	_, err := Service(id)
	if err != nil {
		log.Printf("original error: \n%T %v\n", errors.Cause(err), errors.Cause(err))
		log.Printf("stack error: \n%+v\n", err)
		// return logic, it's up to biz
		//    either nil data and error desc
		//    or replaced data and nil error
	}
	// return user data and nil error
}