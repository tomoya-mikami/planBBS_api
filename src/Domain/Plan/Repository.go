package Plan

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type RepositoryInterface interface {
	Add(plan Plan) error
	FindAll() ([]Plan, error)
}

type Repository struct {
	client *firestore.Client
	ctx context.Context
}

func NewRepository(client *firestore.Client, ctx context.Context) *Repository {
	repository := new(Repository)
	repository.client = client
	repository.ctx = ctx
	return repository
}

func (r Repository) Add(plan *Plan) error {
	var err error
	_, _, err = r.client.Collection("Plans").Add(r.ctx, plan)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

func (r Repository) FindAll() ([]Plan, error) {
	var err error
	plans := make([]Plan, 0)
	iter :=  r.client.Collection("Plans").Documents(r.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
			return make([]Plan, 0), err
		}

		var plan Plan
		doc.DataTo(&plan)
		plans = append(plans, plan)
	}

	return plans, err
}
