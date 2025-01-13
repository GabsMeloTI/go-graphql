// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"database/sql"
)

type BasicInfo struct {
	ID             int64          `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	AdditionalName sql.NullString `json:"additional_name"`
	Pronouns       sql.NullString `json:"pronouns"`
	HeadLine       sql.NullString `json:"head_line"`
}