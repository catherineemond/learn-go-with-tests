package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with one int field",
			struct {
				Name string
				Age  int
			}{"Chris", 12},
			[]string{"Chris"},
		},
		{
			"Nested struct",
			Person{"Chris", Address{"Milky Way", 12345}},
			[]string{"Chris", "Milky Way"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Address{"Milky Way", 12345},
			},
			[]string{"Chris", "Milky Way"},
		},
		{
			"Slices",
			[]Address{
				{"Milky Way", 12345},
				{"Galaxy", 654321},
			},
			[]string{"Milky Way", "Galaxy"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Street string
	Number int
}
