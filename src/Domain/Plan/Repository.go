package Plan

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository struct {
	cleint firestore.Client
	ctx context.Context
}

func (r Repository) Add(plan Plan) error {
	var err error
	_, _, err = r.cleint.Collection("Plans").Add(r.ctx, plan)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

func (r Repository) FindAll() ([]Plan, error) {
	var err error
	plans := make([]Plan, 0)
	iter :=  r.cleint.Collection("Plans").Documents(r.ctx)
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
