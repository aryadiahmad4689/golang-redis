package main

import (
	"fmt"
	"testing"
)

func Test_getData(t *testing.T) {
	redis := CreatePoolRedis()
	repo := NewUserRepo(redis)
	user := User{
		Name:   "ariadi ahmad",
		Alamat: "Makassar",
	}
	repo.Store(&user)

	data, _ := repo.Get()
	fmt.Println(data)
}
