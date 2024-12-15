package repository

import (
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
	return &contactRepository{}
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
