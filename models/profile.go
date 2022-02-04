package models

import "encoding/json"

type Profile struct {
	Name        string
	FamilyName  string
	GivenName   string
	NickName    string
	Picture     string
	Email       string
	PhoneNumber string
	Groups      []string
	Remember    bool
}

func (p Profile) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
