package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Hank"},
			[]string{"Hank"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Hank", "Seoul"},
			[]string{"Hank", "Seoul"},
		},
		{
			"Struct with Non-string fields",
			struct {
				Name string
				Age  int
			}{"Hank", 20},
			[]string{"Hank"},
		},
		{
			"Nested fields",
			Person{
				"Hank",
				Profile{
					20, "Seoul",
				},
			},
			[]string{"Hank", "Seoul"},
		},
		{
			"Pointers fields",
			&Person{
				"Hank",
				Profile{
					20, "Seoul",
				},
			},
			[]string{"Hank", "Seoul"},
		},
		{
			"Slices fields",
			[]Profile{
				{20, "Seoul"},
				{21, "San Francisco"},
			},
			[]string{"Seoul", "San Francisco"},
		},
		{
			"Arrays",
			[2]Profile{
				{20, "Seoul"},
				{21, "San Francisco"},
			},
			[]string{"Seoul", "San Francisco"},
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

	// New test for maps without ordering concerns
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		// Test membership (without caring about ordering)
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	// channels
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Berlin"}
			aChannel <- Profile{31, "Moscow"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Moscow"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, got []string, expected string) {
	t.Helper()
	contains := false
	for _, x := range got {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but not found", got, expected)
	}
}
