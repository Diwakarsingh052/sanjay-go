package handlers

import "dummy-proj/models"

func API() {
	p := []models.Person{
		{
			Name:    "abc",
			Age:     0,
			Address: "",
		},
		{
			Name:    "xyz",
			Age:     0,
			Address: "",
		},
		{
			Name:    "efg",
			Age:     0,
			Address: "",
		},
	}

	for _, v := range p {
		v.Print()
	}

}
