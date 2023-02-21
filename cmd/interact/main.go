package main

import (
	interact "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact/interact/interact/interactservice"
	"log"
)

func main() {
	svr := interact.NewServer(new(InteractServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
