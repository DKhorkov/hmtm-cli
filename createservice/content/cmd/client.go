package cmd

const ClientContent = `package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/DKhorkov/hmtm-sso/api/protobuf/generated/go/sso"
	"github.com/DKhorkov/libs/requestid"
)

// TODO add client gRPC variables inside
type Client struct {}

func main() {
	clientConnection, err := grpc.NewClient(
		// TODO change on your host and port:
		fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		panic(err)
	}

	client := &Client{
		// TODO add gRPC clients here
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), requestid.Key, requestid.New())

	// TODO remove after adding logic:
	fmt.Println(client, ctx)
}
`
