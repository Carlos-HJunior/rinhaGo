package repository

import "rinhaGo/entities"

type PersonRepository interface {
	CreatePerson(person entities.Person) (int64, error)
	ListPersons() ([]entities.Person, error)
	SearchPersonByTerm(string) (*entities.Person, error)
}

func NewPersonRepository() PersonRepository {
	return personRepository{}
}

type personRepository struct {
	persons []entities.Person
}

func (p personRepository) CreatePerson(person entities.Person) (int64, error) {
	p.persons = append(p.persons, person)
	return 0, nil
}

func (p personRepository) ListPersons() ([]entities.Person, error) {
	return p.persons, nil
}

func (p personRepository) SearchPersonByTerm(s string) (*entities.Person, error) {
	//TODO implement me
	panic("implement me")
}
