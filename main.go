package main

import "douyin/common/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()
}
