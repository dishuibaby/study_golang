package person

type Person struct {
	fisrtName string
	lastName  string
}

func (p *Person) FisrtName() string {
	return p.fisrtName
}

func (p *Person) SetFirstName(newName string) {
	p.fisrtName = newName
}
