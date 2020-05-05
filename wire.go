//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"cloud.google.com/go/firestore"

	"local.packages/src"
	"local.packages/plan"
	"local.packages/handler"
)

func InitializeServer(client *firestore.Client, ctx context.Context) (*src.Server, error) {
	wire.Build(
		Plan.NewRepository,
		Plan.NewService,
		Handler.NewPlanHandler,
		src.NewServer,
	)

	return nil, nil
}
