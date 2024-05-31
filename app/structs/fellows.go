package structs

import "strings"

type Fellows map[string]Fellow

type Fellow struct {
	Forename  string `json:"Forename"`
	Surname   string `json:"Surname"`
	Birthdate string `json:"Birthdate"`
	Email     string `json:"Email"`
	Meme      string
}

func (f *Fellows) Emails() []string {
	var emails []string
	for _, fellow := range *f {
		fellow.Email = strings.ToLower(fellow.Email)
		emails = append(emails, fellow.Email)
	}
	return emails
}
