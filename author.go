package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type Author struct {
	Id        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

var authorType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstname": &graphql.Field{
			Type: graphql.String,
		},
		"lastname": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var authorInputType *graphql.InputObject = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"firstname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

func RegisterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type", "application/json")
	var author Author
	json.NewDecoder(request.Body).Decode(&author)
	author.Id = uuid.Must(uuid.NewRandom()).String()
	authors = append(authors, author)
	json.NewEncoder(response).Encode(authors)
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type", "application/json")
	var data Author
	json.NewDecoder(request.Body).Decode(&data)
	for _, author := range authors {
		if author.Username == data.Username {
			json.NewEncoder(response).Encode(author)
		}
	}
	json.NewEncoder(response).Encode(Author{})
}
