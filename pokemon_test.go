package main

import "testing"

func TestEnNames(t *testing.T) {
	cases := []struct {
		id 			int 
		lang, expectedName	string
	}{
		{1, "en", "Bulbasaur"},
		{400, "en", "Bibarel"},
		{721, "en", "Volcanion"},
	}
	for _, c := range cases {
		got := getName(c.id, c.lang)
		if got != c.expectedName {
			t.Errorf("getName(%v, %q) == %q, expectedName: %q", c.id, c.lang, got, c.expectedName)
		}
	}
}