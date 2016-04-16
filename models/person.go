package models

type Person struct {
  FirstName string
  LastName string
  Address string
  PhoneNumber string
  Email string
}

func GetAllContacts()([]*Person, error)  {
  rows, err := db.Query("SELECT * FROM contacts")
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  persons := make([]*Person, 0)

  for rows.Next() {
    person := new(Person)

    err := rows.Scan(&person.FirstName, &person.LastName, &person.Address, &person.PhoneNumber, &person.Email)
    if err != nil {
      return nil, err
    }

    persons = append(persons, person)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return persons, nil

}
