package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"

	"github.com/GabsMeloTI/go-graphql/graph/model"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, data model.PersonInput) (*model.Person, error) {
	panic(fmt.Errorf("not implemented: CreatePerson - createPerson"))
}

// UpdatePerson is the resolver for the updatePerson field.
func (r *mutationResolver) UpdatePerson(ctx context.Context, id string, data model.PersonInput) (*model.Person, error) {
	panic(fmt.Errorf("not implemented: UpdatePerson - updatePerson"))
}

// DeletePerson is the resolver for the deletePerson field.
func (r *mutationResolver) DeletePerson(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeletePerson - deletePerson"))
}

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	panic(fmt.Errorf("not implemented: Persons - persons"))
}

// Person is the resolver for the person field.
func (r *queryResolver) Person(ctx context.Context, id string) (*model.Person, error) {
	panic(fmt.Errorf("not implemented: Person - person"))
}
