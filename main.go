package main

import (
  "github.com/jancuk/simple/models"
  "fmt"
  "net/http"
)

func main()  {
  models.InitDB("postgres://azhar:12345rewq@localhost/contact_person?sslmode=disable")

  http.HandleFunc("/contacts", ContactResource)
  http.ListenAndServe(":3001", nil)
}

func ContactResource(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }
  contacts, err := models.GetAllContacts()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  for _, person := range contacts {
    fmt.Fprintf(w, "%s, %s, %s, %s, %s", person.FirstName, person.LastName, person.PhoneNumber, person.Address, person.Email)
  }

}
