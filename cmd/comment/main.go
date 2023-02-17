package main

import (
	comment "github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/comment/kitex_gen/comment/commentsrv"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
