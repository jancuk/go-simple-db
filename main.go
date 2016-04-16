package main

import (
  "github.com/jancuk/simple/models"
  "fmt"
  "net/http"
  "encoding/json"
)

type Response struct {
  Page int                `json:"page"`
  Data []*models.Person   `json:"data"`
}

func main()  {
  models.InitDB("postgres://azhar:12345rewq@localhost/contact_person?sslmode=disable")

  http.HandleFunc("/contacts", ContactResource)
  http.ListenAndServe(":3001", nil)
}

func ContactResource(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }
  contacts, err := models.GetAllContacts()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  GetResponse := &Response{ Page: 1, Data: contacts}

  ResultSerialize, _ := json.Marshal(GetResponse)
  fmt.Fprintf(w, "%s",string(ResultSerialize))

}
