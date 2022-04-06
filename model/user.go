package model

type User struct {
	FirstName string `json:"firstName" pact:"example=Sally"`
	LastName  string `json:"lastName" pact:"example=McSmiley Face😀😍"`
	Username  string `json:"username" pact:"example=sally"`
	Type      string `json:"type" pact:"example=admin,regex=^(admin|user|guest)$"`
	ID        int    `json:"id" pact:"example=10"`
}
