package infra

import (
	"database/sql"
	"github.com/GabsMeloTI/go-graphql/infra/database"
	"github.com/GabsMeloTI/go-graphql/infra/database/db_postgresql"
)

type ContainerDI struct {
	Config ConfigEnv
	ConnDB *sql.DB
}

func NewContainerDI(config ConfigEnv) *ContainerDI {
	container := &ContainerDI{Config: config}
	container.db()
	container.buildRepository()
	container.buildService()
	container.buildHandler()
	return container
}

func (c *ContainerDI) db() {
	dbConfig := database.Config{
		Host:        c.Config.DBHost,
		Port:        c.Config.DBPort,
		User:        c.Config.DBUser,
		Password:    c.Config.DBPassword,
		Database:    c.Config.DBDatabase,
		SSLMode:     c.Config.DBSSLMode,
		Driver:      c.Config.DBDriver,
		Environment: c.Config.Environment,
	}
	c.ConnDB = db_postgresql.NewConnection(&dbConfig)
}

func (c *ContainerDI) buildRepository() {

}

func (c *ContainerDI) buildService() {

}

func (c *ContainerDI) buildHandler() {

}
