package maps

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertErr(t *testing.T, got, want error) {
	t.Helper()
	if got.Error() != want.Error() {
		t.Errorf("got error '%s' want '%s' ", got.Error(), want.Error())
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, keyword, definition string) {
	t.Helper()
	got, err := dictionary.Search(keyword)
	want := definition

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known keyword", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown keyword", func(t *testing.T) {
		_, err := dictionary.Search("nu")

		assertErr(t, err, errKeywordNotFind)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	keyword := "test"
	definition := "this is just a test"
	dictionary.Add(keyword, definition)

	assertDefinition(t, dictionary, keyword, definition)
}
