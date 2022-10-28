package main

import (
	"fmt"

	"github.com/shuymn-sandbox/go-generic-encoding-sample/conv"
)

type Animal struct {
	Name string
}

type Cat struct {
	Name string
}

var (
	_ conv.Marshaler[*Animal]   = (*Cat)(nil)
	_ conv.Unmarshaler[*Animal] = (*Cat)(nil)
)

func (c *Cat) Marshal() (*Animal, error) {
	return &Animal{Name: c.Name}, nil
}

func (c *Cat) Unmarshal(animal *Animal) error {
	if animal == nil {
		return nil
	}
	c.Name = animal.Name
	return nil
}

func main() {
	c1 := &Cat{Name: "tama"}
	a1, _ := conv.Marshal[*Animal](c1)
	fmt.Printf("%#v\n", a1)

	var c2 Cat
	a2 := &Animal{Name: "shiro"}
	_ = conv.Unmarshal[*Animal](a2, &c2)
	fmt.Printf("%#v\n", &c2)
}
