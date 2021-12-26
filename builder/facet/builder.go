package facet

type Person struct {
	// address
	StreetAddress, Postcode, City string
	//job
	CompanyName, Position string
	AnnualIncome          int
}
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// Used to swap between builders
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Work() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// DSL
type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}
func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}
