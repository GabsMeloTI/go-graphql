// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BasicInfo struct {
	ID             string  `json:"id"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	AdditionalName string  `json:"additionalName"`
	Pronouns       *string `json:"pronouns,omitempty"`
	HeadLine       *string `json:"headLine,omitempty"`
}

type BasicInfoInput struct {
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	AdditionalName *string `json:"additionalName,omitempty"`
	Pronouns       *string `json:"pronouns,omitempty"`
	HeadLine       *string `json:"headLine,omitempty"`
}

type Mutation struct {
}

type Query struct {
}
