package repository

import (
	"embed"
	"encoding/json"
	"htmx-contacts/internal/model"
	"strings"
)

type ContactRepository interface {
	Search(text string) []model.Contact
	All() []model.Contact
}

type contactRepository struct {
	lastId   int
	contacts []model.Contact
}

//go:embed *.json
var dummyDataFS embed.FS

func NewContactRepository() ContactRepository {
	data, err := dummyDataFS.ReadFile("dummy_contacts.json")
	if err != nil {
		panic(err)
	}

	// Read in the initial set of dummy contacts
	var cs []model.Contact
	err = json.Unmarshal(data, &cs)
	if err != nil {
		panic(err)
	}

	// Find the last ID in use
	var lastId int
	for _, c := range cs {
		lastId = max(lastId, c.Id)
	}

	return &contactRepository{
		lastId:   lastId,
		contacts: cs,
	}
}

func (r *contactRepository) Search(text string) []model.Contact {
	ret := make([]model.Contact, 0)
	for _, c := range r.contacts {
		if strings.Contains(c.First, text) {
			ret = append(ret, c)
		} else if strings.Contains(c.Last, text) {
			ret = append(ret, c)
		} else if strings.Contains(c.Email, text) {
			ret = append(ret, c)
		} else if strings.Contains(c.Phone, text) {
			ret = append(ret, c)
		}
	}
	return ret
}

func (r *contactRepository) All() []model.Contact {
	return r.contacts
}
