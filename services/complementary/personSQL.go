package complementary

import (
	"errors"
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/SQL"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InsertPerson(p model.Person) {
	_, err := SQL.DB.Exec("INSERT INTO people (ID, Name, LastName, Email, Telephone, Country) VALUES (?, ?, ?, ?, ?, ?)", p.ID, p.Name, p.LastName, p.Email, p.Telephone, p.Country)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("new person inserted correctly")
}

func UpdatePerson(p model.Person) {
	_, err := SQL.DB.Exec("UPDATE people SET Name = ?, LastName = ?, Email = ?, Telephone = ?, Country = ? WHERE ID = ?", p.Name, p.LastName, p.Email, p.Telephone, p.Country, p.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("person updated correctly")
}

func DeletePerson(ID string) {
	_, err := SQL.DB.Exec("DELETE FROM people WHERE ID = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("person deleted correctly")
}

func SelectPersonByID(ID string) (model.Person, error) {
	rows, err := SQL.DB.Query("SELECT * FROM people WHERE ID = ?", ID)
	if err != nil {
		return model.Person{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return model.Person{}, errors.New("person not found")
	}
	var person model.Person
	if err := rows.Scan(&person.ID, &person.Name, &person.LastName, &person.Email, &person.Telephone, &person.Country); err != nil {
		return model.Person{}, err
	}
	fmt.Println(person)
	return person, nil
}

func ContainsPersonByID(ID string) (bool, error) {
	rows, err := SQL.DB.Query("SELECT 1 FROM product WHERE ID = ?", ID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func SelectAllPersons() ([]model.Person, error) {
	rows, err := SQL.DB.Query("SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []model.Person
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.ID, &person.Name, &person.LastName, &person.Email, &person.Telephone, &person.Country); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return persons, nil
}

func CountPersons() (int, error) {
	var count int
	err := SQL.DB.QueryRow("SELECT COUNT(*) FROM people").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
