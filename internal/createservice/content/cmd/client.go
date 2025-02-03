package cmd

const ClientContent = `package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/DKhorkov/hmtm-orders/api/protobuf/generated/go/<service-name>"
	"github.com/DKhorkov/libs/requestid"
)

type Client struct {
	<service-name>.<service-name-title>ServiceClient
}

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
		<service-name-title>ServiceClient: <service-name>.New<service-name-title>ServiceClient(clientConnection),
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), requestid.Key, requestid.New())
}
`
