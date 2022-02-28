package main

import (
	_ "gameMail/controller"
	"gameMail/router"
	_ "gameMail/utils"
)

func main() {
	//开启数据库及缓存

	//开启路由
	router.StartRouter()

}
