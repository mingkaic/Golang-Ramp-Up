package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Contact type with Name, Phone, SSN
type Contact struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	SSN     string `json:"ssn"`
	Email 	string `json:"email,omitempty"`
	Address string `json:"description,omitempty"`
}

var contacts = map[string]Contact{
	"1": Contact{Name: "John", Phone: "1", SSN: "1", Email: "john@infoblox.com"},
	"2": Contact{Name: "Brenda", Phone: "2", SSN: "2", Email: "brenda@infoblox.com"},
}

// ToJSON to be used for marshalling of Contact type
func (c Contact) ToJSON() []byte {
	ToJSON, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Contact type
func FromJSON(data []byte) Contact {
	contact := Contact{}
	err := json.Unmarshal(data, &contact)
	if err != nil {
		panic(err)
	}
	return contact
}

// AllContacts returns a slice of all contacts
func AllContacts() []Contact {
	values := make([]Contact, len(contacts))
	idx := 0
	for _, contact := range contacts {
		values[idx] = contact
		idx++
	}
	return values
}

// ContactsHandleFunc to be used as http.HandleFunc for Contacts API
func ContactsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		contacts := AllContacts()
		writeJSON(w, contacts)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		contact := FromJSON(body)
		ssn, created := CreateContact(contact)
		if created {
			w.Header().Add("Location", "/api/contacts/"+ssn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// ContactHandleFunc to be used as http.HandleFunc for Contact API
func ContactHandleFunc(w http.ResponseWriter, r *http.Request) {
	ssn := r.URL.Path[len("/api/contacts/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		contact, found := GetContact(ssn)
		if found {
			writeJSON(w, contact)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		contact := FromJSON(body)
		exists := UpdateContact(ssn, contact)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteContact(ssn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// GetContact returns the contact for a given SSN
func GetContact(ssn string) (Contact, bool) {
	contact, found := contacts[ssn]
	return contact, found
}

// CreateContact creates a new Contact if it does not exist
func CreateContact(contact Contact) (string, bool) {
	_, exists := contacts[contact.SSN]
	if exists {
		return "", false
	}
	contacts[contact.SSN] = contact
	return contact.SSN, true
}

// UpdateContact updates an existing contact
func UpdateContact(ssn string, contact Contact) bool {
	_, exists := contacts[ssn]
	if exists {
		contacts[ssn] = contact
	}
	return exists
}

// DeleteContact removes a contact from the map by SSN key
func DeleteContact(ssn string) {
	delete(contacts, ssn)
}