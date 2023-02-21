package db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Init()

	cmt := &Comment{UserId: 123, Content: "DSADAS", VideoId: 0xefffffff, IsValid: true,
		CreateTime: time.Now().String()}
	err := CreateComment(context.Background(), []*Comment{cmt})
	if err != nil {
		t.Fatal(err)
	}
	defer deleteAll(context.Background(), cmt.VideoId)

	res, err := QueryComments(context.Background(), cmt.VideoId)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].ID != cmt.ID {
		t.Fatal()
	}
	fmt.Println(*res[0])
	err = DeleteComment(context.Background(), cmt)
	if err != nil {
		t.Fatal(err)
	}
	res, err = QueryComments(context.Background(), cmt.VideoId)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 0 {
		t.Fatal()
	}

	// c, err := SelectComment(context.Background(), 1)
	// fmt.Println(err, c)
}
