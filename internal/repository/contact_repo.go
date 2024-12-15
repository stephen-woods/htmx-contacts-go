package repository

import (
	"encoding/json"
	"htmx-contacts/internal/model"
	"strings"
)

type ContactRepository interface {
	Search(text string) []model.Contact
	All() []model.Contact
}

type contactRepository struct {
	contacts []model.Contact
}

func NewContactRepository() ContactRepository {
	var cs []model.Contact
	err := json.Unmarshal(dummyContacts, &cs)
	if err != nil {
		panic(err)
	}
	return &contactRepository{
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

var dummyContacts = []byte(`
[
  {
    "id": 2,
    "first": "Carson",
    "last": "Gross",
    "phone": "123-456-7890",
    "email": "carson@example.comz",
    "errors": {}
  },
  {
    "id": 3,
    "first": "",
    "last": "",
    "phone": "",
    "email": "joe@example2.com",
    "errors": {}
  },
  {
    "id": 5,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe@example.com",
    "errors": {}
  },
  {
    "id": 6,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe1@example.com",
    "errors": {}
  },
  {
    "id": 7,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe2@example.com",
    "errors": {}
  },
  {
    "id": 8,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe3@example.com",
    "errors": {}
  },
  {
    "id": 9,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe4@example.com",
    "errors": {}
  },
  {
    "id": 10,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe5@example.com",
    "errors": {}
  },
  {
    "id": 11,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe6@example.com",
    "errors": {}
  },
  {
    "id": 12,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe7@example.com",
    "errors": {}
  },
  {
    "id": 13,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe8@example.com",
    "errors": {}
  },
  {
    "id": 14,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe9@example.com",
    "errors": {}
  },
  {
    "id": 15,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe10@example.com",
    "errors": {}
  },
  {
    "id": 16,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe11@example.com",
    "errors": {}
  },
  {
    "id": 17,
    "first": "Joe",
    "last": "Blow",
    "phone": "123-456-7890",
    "email": "joe12@example.com",
    "errors": {}
  },
  {
    "id": 18,
    "first": null,
    "last": null,
    "phone": null,
    "email": "restexample1@example.com",
    "errors": {}
  },
  {
    "id": 19,
    "first": null,
    "last": null,
    "phone": null,
    "email": "restexample2@example.com",
    "errors": {}
  }
]`)
