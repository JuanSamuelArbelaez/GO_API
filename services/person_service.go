package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/Utils"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services/complementary"
)

func GetPeople() ([]model.Person, error) {
	return complementary.SelectAllPersons()
}

func GetNumberOfPeople() (size int, err error) {
	return complementary.CountAllPersons()
}

func AddPerson(p model.PersonRequest) (id string, err error) {
	err = CheckPersonRequest(p)
	if err != nil {
		return "", err
	}

	newId := ""
	err = Utils.GenerateId(&newId)

	if err != nil {
		return "", err
	} else {
		newPerson := model.Person{}
		newPerson.ID = newId
		newPerson.State = "ACTIVE"
		model.MapRequest(&newPerson, &p)
		complementary.InsertPerson(newPerson)
		return newId, nil
	}
}

func UpdatePerson(p model.Person) (err error) {
	exists, err := complementary.PersonByIDExists(p.ID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("person not found")
	}

	active, err := complementary.UserIsActive(p.ID)
	if err != nil {
		return err
	} else if !active {
		return fmt.Errorf("user is not active")
	}

	err = CheckPerson(p)
	if err != nil {
		return err
	}

	complementary.UpdatePerson(p)
	return nil
}

func RemovePerson(ID string) (state string, err error) {
	exists, err := complementary.PersonByIDExists(ID)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", fmt.Errorf("person not found")
	}

	active, err := complementary.UserIsActive(ID)
	if err != nil {
		return "", err
	} else if !active {
		return "", fmt.Errorf("user already inactive")
	} else {
		complementary.DeletePerson(ID)
	}

	return complementary.GetUserState(ID)
}

func RecoverPerson(ID string) (state string, err error) {
	exists, err := complementary.PersonByIDExists(ID)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", fmt.Errorf("person not found")
	}

	active, err := complementary.UserIsActive(ID)
	if err != nil {
		return "", err
	} else if active {
		return "", fmt.Errorf("user already active")
	} else {
		complementary.RecoverPerson(ID)
	}

	return complementary.GetUserState(ID)
}

func GetPerson(ID string) (product model.Person, err error) {
	return complementary.SelectPersonByID(ID)
}

func CheckPersonRequest(person model.PersonRequest) (err error) {
	var p model.Person
	model.MapRequest(&p, &person)
	p.ID = ""
	return CheckPerson(p)
}

func CheckPerson(person model.Person) (err error) {
	errorMsg := ""

	if Utils.EmptyString(person.Name) {
		errorMsg += "Invalid person name. "
	}
	if Utils.EmptyString(person.LastName) {
		errorMsg += "Invalid person last name. "
	}
	if Utils.NotEmailString(person.Email) {
		errorMsg += "Invalid person email. "
	}
	if exists, err := complementary.PersonByEmailExists(person.ID, person.Email); err != nil {
		errorMsg += fmt.Sprint(err) + " "
	} else if exists {
		errorMsg += "Email already in use. "
	}
	if Utils.EmptyString(person.Telephone) {
		errorMsg += "Invalid person telephone. "
	}
	if NotCountryCode(person.Country) {
		errorMsg += "Invalid country code. "
	}
	if errorMsg != "" {
		return fmt.Errorf(errorMsg)
	}
	return nil
}
