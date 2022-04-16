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
			}{"Srdan"},
			[]string{"Srdan"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Srdan", "Auckland"},
			[]string{"Srdan", "Auckland"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Srdan", 36},
			[]string{"Srdan"},
		},
		{
			"Struct with nested fields",
			Person{
				"Srdan",
				Profile{36, "Auckland"},
			},
			[]string{"Srdan", "Auckland"},
		},
		{
			"Pointer to objects",
			&Person{
				"Srdan",
				Profile{36, "Auckland"},
			},
			[]string{"Srdan", "Auckland"},
		},
		{
			"Slices of objects",
			[]Profile{
				{36, "Auckland"},
				{30, "Wellington"},
			},
			[]string{"Auckland", "Wellington"},
		},
		{
			"Arrays of objects",
			[2]Profile{
				{36, "Auckland"},
				{30, "Wellington"},
			},
			[]string{"Auckland", "Wellington"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("testing with channels", func(t *testing.T) {
		aChan := make(chan Profile)

		go func() {
			aChan <- Profile{36, "Auckland"}
			aChan <- Profile{30, "Wellington"}
			close(aChan)
		}()

		var got []string
		want := []string{"Auckland", "Wellington"}

		walk(aChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("testing with a function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{36, "Auckland"}, Profile{30, "Wellington"}
		}

		var got []string
		want := []string{"Auckland", "Wellington"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q, but it didn't", haystack, needle)
	}
}
