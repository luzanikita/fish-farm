package stats

import (
	"context"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client StatsServiceClient
}

const SERVER_ADDR = "127.0.0.1:6000"

func InitGrpcConnection() (*GrpcClient, error) {
	conn, err := grpc.Dial(SERVER_ADDR, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewStatsServiceClient(conn)
	return &GrpcClient{conn, client}, nil
}

func (g *GrpcClient) MyStats(text string) (string, error) {
	req := StatsRequest{
		Text: text,
	}

	res, err := g.client.GetStats(context.Background(), &req)
	if err != nil {
		return "", err
	}

	return res.Text, nil
}
