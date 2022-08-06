package main

import (
	"GinBlog/model"
	"GinBlog/routers"
	"GinBlog/utils"
)

func main() {
	model.InitDb()
	r := routers.InitRouter()
	r.Run(utils.HttpPort)
}
