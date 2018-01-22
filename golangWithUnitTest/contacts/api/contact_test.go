package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactToJSON(t *testing.T) {
	contact := Contact{Name: "John", Phone: "1", SSN: "1"}
	json := contact.ToJSON()

	assert.Equal(t, `{"name":"John","phone":"1","ssn":"1"}`, string(json), "Contact JSON marshalling wrong.")
}

func TestContactFromJSON(t *testing.T) {
	json := []byte(`{"name":"John","phone":"1","ssn":"1"}`)
	contact := FromJSON(json)
	assert.Equal(t, Contact{Name: "John", Phone: "1", SSN: "1"}, contact, "Contact JSON unmarshalling wrong.")
}

func TestAllContacts(t *testing.T) {
	contacts := AllContacts()
	assert.Len(t, contacts, 2, "Wrong number of contacts.")
}

func TestCreateNewContact(t *testing.T) {
	contact := Contact{Name: "Josh", Phone: "5", SSN: "5"}
	ssn, created := CreateContact(contact)
	assert.True(t, created, "Contact was not created.")
	assert.Equal(t, "5", ssn, "Wrong SSN.")
}

func TestDoNotCreateExistingContact(t *testing.T) {
	contact := Contact{SSN: "1"}
	_, created := CreateContact(contact)
	assert.False(t, created, "Contact was created.")
}

func TestUpdateExistingContact(t *testing.T) {
	contact := Contact{Name: "John B", Phone: "New 1", SSN: "1", Address: "Annapolis"}
	updated := UpdateContact("1", contact)
	assert.True(t, updated, "Contact not updated.")

	contact, _ = GetContact("1")
	assert.Equal(t, "New 1", contact.Phone, "Phone not updated.")
	assert.Equal(t, "John B", contact.Name, "Name not updated.")
	assert.Equal(t, "Annapolis", contact.Address, "Address not updated.")
}

func TestDeleteContact(t *testing.T) {
	DeleteContact("1")
	assert.Len(t, AllContacts(), 2, "Wrong number of contacts after delete.")
}