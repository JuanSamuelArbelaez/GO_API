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
	_, err := SQL.DB.Exec("UPDATE people SET Name = ?, LastName = ?, Email = ?, Telephone = ?, Country = ? WHERE ID = ? AND State = ?", p.Name, p.LastName, p.Email, p.Telephone, p.Country, p.ID, "ACTIVE")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("person updated correctly")
}

func RecoverPerson(ID string) {
	_, err := SQL.DB.Exec("UPDATE people SET State = ? WHERE ID = ? AND State = ?", "ACTIVE", ID, "INACTIVE")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("person active")
}

func DeletePerson(ID string) {
	_, err := SQL.DB.Exec("UPDATE people SET State = ? WHERE ID = ? AND State = ?", "INACTIVE", ID, "ACTIVE")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("person no longer active")
}

func SelectPersonByID(ID string) (p model.Person, e error) {
	rows, err := SQL.DB.Query("SELECT * FROM people WHERE ID = ?", ID)
	if err != nil {
		return model.Person{}, err
	} else {
		fmt.Println(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return model.Person{}, errors.New("person not found")
	}
	var person model.Person
	if err := rows.Scan(&person.ID, &person.Name, &person.LastName, &person.Email, &person.Telephone, &person.Country, &person.State); err != nil {
		return model.Person{}, err
	}
	return person, err
}

func PersonByIDExists(ID string) (bool, error) {
	rows, err := SQL.DB.Query("SELECT 1 FROM people WHERE ID = ?", ID)
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

func PersonByEmailExists(id, email string) (bool, error) {
	rows, err := SQL.DB.Query("SELECT 1 FROM people WHERE Email = ? AND ID != ?", email, id)
	if err != nil {
		return false, err
	}

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
		if err := rows.Scan(&person.ID, &person.Name, &person.LastName, &person.Email, &person.Telephone, &person.Country, &person.State); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return persons, nil
}

func CountAllPersons() (int, error) {
	var count int
	err := SQL.DB.QueryRow("SELECT COUNT(*) FROM people").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountActivePersons() (int, error) {
	var count int
	err := SQL.DB.QueryRow("SELECT COUNT(*) FROM people WHERE State = ?", "ACTIVE").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountInactivePersons() (int, error) {
	var count int
	err := SQL.DB.QueryRow("SELECT COUNT(*) FROM people WHERE State = ?", "INACTIVE").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UserIsActive(ID string) (isActive bool, err error) {
	err = SQL.DB.QueryRow("SELECT (State = ?) as isActive FROM people WHERE ID = ?", "ACTIVE", ID).Scan(&isActive)
	if err != nil {
		return false, err
	}
	return isActive, nil
}

func GetUserState(ID string) (State string, err error) {
	var state string
	err = SQL.DB.QueryRow("SELECT State FROM people WHERE ID = ?", ID).Scan(&state)
	if err != nil {
		return "", err
	}
	return state, nil
}
