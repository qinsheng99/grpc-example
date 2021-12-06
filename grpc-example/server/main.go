package main

import (
	"context"
	"log"
	"net"

	pb "github.com/qinsheng99/grpc-example/grpc-example/route"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type RouteGuideServer struct {
	features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func NewServer() *RouteGuideServer  {
	return &RouteGuideServer{
		features: []*pb.Feature{
			{
				Name: "上海交通大学",
				Location: &pb.Point{
					Latitude: 310235000,
					Longitude: 121437403,
				},
			},
			{
				Name: "复旦大学",
				Location: &pb.Point{
					Latitude: 310235001,
					Longitude: 121437403,
				},
			},
		},
	}
}

func (s *RouteGuideServer)GetFeature(ctx context.Context,point  *pb.Point) (*pb.Feature, error)  {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

func (s *RouteGuideServer) ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error {
	return nil
}
func (s *RouteGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}
func (s *RouteGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}
func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, NewServer())
	log.Fatalln(grpcServer.Serve(lis))
}
