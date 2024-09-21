
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init()  {
	var err error
	conn,err = grpc.Dial("127.0.0.1:50051",grpc.WithInsecure())
	if err!= nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

func test(){
	//user,err :=userClient.Create(context.Background(),&proto.CreateUserInfo{
	//	NickName: "yb",
	//	Password: "123456",
	//	Mobile: "18696544688",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(*user)
	user,err := userClient.GetUserById(context.Background(),&proto.IdRequest{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(*user)
	var userListRps *proto.UserListResponse
	userListRps,err = userClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: 1,
		PSize: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(*userListRps)


}

func main()  {
	Init()
	test()
	conn.Close()
}

