package service

import (
	"exam/dao"
	"log"
)

func IsNameExist(name string) bool {
	_, err := dao.NameQuery(name)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
