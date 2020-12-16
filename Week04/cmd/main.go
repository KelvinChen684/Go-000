package main

import (
	"context"
	pb "github.com/KelvinChen684/Go-000/tree/main/Week04/api/v1"
	kratos "github.com/KelvinChen684/Go-000/tree/main/Week04/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	userSrv, err := InitializeUserService()
	if err != nil {
		log.Fatalf("Initial failed: %s\n", err)
	}
	lis, _ := net.Listen("tcp", ":8000")
	grpcSrv := grpc.NewServer()
	pb.RegisterUserServer(grpcSrv, userSrv)

	start := func(ctx context.Context) error {
		return grpcSrv.Serve(lis)
	}
	stop := func(ctx context.Context) error {
		grpcSrv.Stop()
		return nil
	}

	app := kratos.New()
	app.Append(kratos.Hook{start, stop})

	if err = app.Run(); err != nil {
		log.Printf("User service run failed: %v\n", err)
	}

	/*
	lis, _ := net.Listen("tcp", ":8000")
	s := grpc.NewServer()
	//pb.RegisterUserServer(s, &grpcserver.GrpcService{})
	go s.Serve(lis)

	conn, _ := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	grpcCli := pb.NewUserClient(conn)


	r := gin.Default()

	r.GET("/getname", func(c *gin.Context) {
		name := c.Query("name")
		req := &pb.HelloReq{Name: name}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		fmt.Println("before grpc")
		resp, _ := grpcCli.SayHello(ctx, req)
		fmt.Println("after grpc")
		c.String(200, resp.Content)

	})

	r.Run(":8080")
	 */
}
