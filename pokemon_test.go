package pokemon

import "testing"

func TestNames(t *testing.T) {
	cases := []struct {
		id                 int
		lang, expectedName string
	}{
		{1, "en", "Bulbasaur"},
		{400, "en", "Bibarel"},
		{721, "en", "Volcanion"},
		{1, "zh-Hant", "妙蛙種子"},
		{400, "zh-Hant", "大尾狸"},
		{721, "zh-Hant", "波爾凱尼恩"},
	}
	for _, c := range cases {
		got := PokemonName(c.id, c.lang)
		if got != c.expectedName {
			t.Errorf("getName(%v, %q) == %q, expectedName: %q", c.id, c.lang, got, c.expectedName)
		}
	}
}

func TestIds(t *testing.T) {
	cases := []struct {
		expectedID int
		lang, name string
	}{
		{1, "en", "Bulbasaur"},
		{400, "en", "Bibarel"},
		{721, "en", "Volcanion"},
		{1, "zh-Hant", "妙蛙種子"},
		{400, "zh-Hant", "大尾狸"},
		{721, "zh-Hant", "波爾凱尼恩"},
	}
	for _, c := range cases {
		got := PokemonId(c.name, c.lang)
		if got != c.expectedID {
			t.Errorf("getId(%q, %q) == %v, expectedID: %v", c.name, c.lang, got, c.expectedID)
		}
	}
}
