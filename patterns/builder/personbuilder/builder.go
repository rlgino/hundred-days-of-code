package personbuilder

import "fmt"

// PersonBuilder it's the builder of person
type PersonBuilder interface {
	WithName(name string) PersonBuilder
	WithAge(age int) PersonBuilder
	WithAddress(address string) PersonBuilder
	Build() Person
}

// Person struct
type Person struct {
	Name    string
	Age     int
	Address string
}

func (p *Person) String() string {
	return fmt.Sprintf("\n Name: %s \n Age: %d \n Address: %s", p.Name, p.Age, p.Address)
}

type personBuilder struct {
	person Person
}

// New constructor
func New() PersonBuilder {

	return &personBuilder{
		Person{},
	}
}

func (builder *personBuilder) WithName(name string) PersonBuilder {
	builder.person.Name = name
	return builder
}

func (builder *personBuilder) WithAge(age int) PersonBuilder {
	builder.person.Age = age
	return builder
}

func (builder *personBuilder) WithAddress(address string) PersonBuilder {
	builder.person.Address = address
	return builder
}

func (builder *personBuilder) Build() Person {
	return builder.person
}
