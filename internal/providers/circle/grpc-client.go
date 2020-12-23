package circle

import (
	pb "buycoins-raven/internal/providers/circle/genproto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

func CreateCardResponse(request interface{}) (string, error) {
	var opts []grpc.DialOption
	grpcPort := "50055"
	host := "0.0.0.0"
	srvAddr := fmt.Sprintf("%s:%s", host, grpcPort)

	opts = append(opts, grpc.WithAuthority(host), grpc.WithInsecure())

	conn, err := grpc.Dial(srvAddr, opts...)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewCardClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	out, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	res, err := client.CreateCard(ctx, &pb.Req{Payload: string(out)})

	if err != nil {
		got := status.Code(err)
		return "" , fmt.Errorf("error occurred in getting request:, %v ", got)
	}

	return  res.Payload, nil
}

func MakePaymentResponse(request interface{}) (string, error) {
	var opts []grpc.DialOption
	grpcPort := "50055"
	host := "0.0.0.0"
	srvAddr := fmt.Sprintf("%s:%s", host, grpcPort)

	opts = append(opts, grpc.WithAuthority(host), grpc.WithInsecure())

	conn, err := grpc.Dial(srvAddr, opts...)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewPaymentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	out, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	res, err := client.MakePayment(ctx, &pb.Req{Payload: string(out)})

	if err != nil {
		got := status.Code(err)
		return "" , fmt.Errorf("error occurred in getting request:, %v ", got)
	}

	return  res.Payload, nil
}
