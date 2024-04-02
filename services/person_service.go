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

func GetNumberOfPeople() (size int, e error) {
	return complementary.CountPersons()
}

func AddPerson(p model.PersonRequest) (id string, e error) {
	e = CheckPerson(p)
	if e != nil {
		return "", e
	}

	newId := ""
	e = Utils.GenerateId(&newId)

	if e != nil {
		return "", e
	} else {
		newPerson := model.Person{}
		newPerson.ID = newId
		model.MapRequest(&newPerson, &p)
		complementary.InsertPerson(newPerson)
		return newId, nil
	}
}

func RemovePerson(ID string) (e error) {
	if contained, e := ContainsPerson(ID); e != nil {
		return e
	} else if !contained {
		return fmt.Errorf("person with ID:'%s' not found", ID)
	} else {
		return nil
	}
}

func ContainsPerson(ID string) (isContained bool, e error) {
	return complementary.ContainsPersonByID(ID)
}

func GetPerson(ID string) (product model.Person, e error) {
	p, _ := complementary.SelectPersonByID(ID)
	fmt.Println(p)
	return p, nil
}

func CheckPerson(person model.PersonRequest) (e error) {
	errorMsg := ""
	errorOccurred := false

	if Utils.EmptyString(person.Name) {
		errorOccurred = true
		errorMsg += "Invalid person name. "
	}
	if Utils.EmptyString(person.LastName) {
		errorOccurred = true
		errorMsg += "Invalid person last name. "
	}
	if Utils.NotEmailString(person.Email) {
		errorOccurred = true
		errorMsg += "Invalid person email. "
	}
	if Utils.EmptyString(person.Telephone) {
		errorOccurred = true
		errorMsg += "Invalid person telephone. "
	}
	if Utils.NotCountryCode(person.Country) {
		errorOccurred = true
		errorMsg += "Invalid country code. "
	}

	if errorOccurred {
		return fmt.Errorf(errorMsg)
	}
	return nil
}
