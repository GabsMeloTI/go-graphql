package resolvers

import (
	"database/sql"
	"github.com/GabsMeloTI/go-graphql/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB                *sql.DB
	InterfaceResolver repository.InterfaceRepository
}
