package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/qinsheng99/grpc-example/grpc-example/route"

	"google.golang.org/grpc"
)

func RunFirst(client pb.RouteGuideClient)  {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude: 310235001,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}
func main() {
	coon, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalln("client can not grpc server")
	}
	defer coon.Close()

	client := pb.NewRouteGuideClient(coon)

	RunFirst(client)

}
