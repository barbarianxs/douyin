package main
import(
	"context"
	"log"
	// "time"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
)
import "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
import "github.com/cloudwego/kitex/client"
import "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user/userservice"
// type UserRequest struct {
// 	Username string `thrift:"username,1" frugal:"1,default,string" json:"username"`
// 	Password string `thrift:"password,2" frugal:"2,default,string" json:"password"`
// }
func main(){
	c, err := userservice.NewClient("RegisterUser", client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
	  log.Fatal(err)
	}
	// u := UserRequest{
	// 	Username: "guo",
	// 	Password: "123456",
	// } 
	req := &user.RegisterUserRequest{Username:"guo", Password:"123456"}
	resp, err := c.RegisterUser(context.Background(), req)
	if err != nil {
	  log.Fatal(err)
	}
	log.Println(resp)

}
