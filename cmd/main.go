package main

import (
	"exam/api"
	"exam/dao"
)

func main() {
	dao.OpenDb()
	api.User()
}
