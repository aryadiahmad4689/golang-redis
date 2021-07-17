package main

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

type UserRepository struct {
	redis *redis.Pool
}

func NewUserRepo(conn *redis.Pool) UserInterface {
	return &UserRepository{redis: conn}
}

func (redis *UserRepository) Store(user *User) error {
	conn := redis.redis.Get()
	defer conn.Close()
	serialized, _ := json.Marshal(&user)

	_, err := conn.Do("SETEX", "name", 60*100, serialized)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", "ariadi", user, err.Error())
		return err
	}
	return err
}

func (rediss *UserRepository) Get() (user interface{}, err error) {
	conn := rediss.redis.Get()
	defer conn.Close()
	s, errors := redis.String(conn.Do("GET", "name"))
	if errors != nil {
		log.Printf("ERROR: " + errors.Error())
		return user, err
	}
	return s, err
}
