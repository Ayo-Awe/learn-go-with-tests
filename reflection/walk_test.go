package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age      int
	Location string
}

func TestWalk(t *testing.T) {

	testcases := []struct {
		name     string
		x        interface{}
		expected []string
	}{

		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			name:     "one string field",
			x:        struct{ Name string }{"Dele"},
			expected: []string{"Dele"},
		},
		{
			name: "two string field",
			x: struct {
				Fieldirstname string
				Lastname      string
			}{"Dele", "Ali"},
			expected: []string{"Dele", "Ali"},
		},
		{
			name: "non string field",
			x: struct {
				Name string
				Age  int
			}{"Dele", 23},
			expected: []string{"Dele"},
		},
		{
			name: "no string field",
			x: struct {
				Age int
			}{23},
			expected: nil,
		},
		{
			name: "nested struct fields",
			x: Person{
				Name:    "Ada",
				Profile: Profile{Age: 20, Location: "Lagos"},
			},
			expected: []string{"Ada", "Lagos"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string

			walk(tc.x, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected %v, but got %v", tc.expected, got)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		var got []string
		expected := []string{"Moo", "Baa"}
		x := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		walk(x, func(s string) {
			got = append(got, s)
		})

		// assert same length
		if len(got) != len(expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}

		// assert all elements in the expected are in the got
		for _, ele := range expected {
			if !slices.Contains(got, ele) {
				t.Errorf("expected %v, but got %v", expected, got)
			}
		}
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
