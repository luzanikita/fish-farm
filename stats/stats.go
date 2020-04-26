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

func (g *GrpcClient) MyStats(conditions []*Condition) (*ConditionsStatsResponse, error) {
	req := ConditionsStatsRequest{
		Conditions: conditions,
	}

	res, err := g.client.GetConditionsStats(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
