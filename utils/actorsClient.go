package utils

import (
	"google.golang.org/grpc"
	"github.com/free-way/riverwaveCommon/definitions"
)

var(
	Err error
	ActorsConnection *grpc.ClientConn
	ActorsClient definitions.ActorsServiceClient

	ResourcesClient definitions.ResourcesClient
)
