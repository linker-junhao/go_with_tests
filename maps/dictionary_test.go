package maps

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertErr(t *testing.T, got, want DictionaryErr) {
	t.Helper()
	if got.Error() != want.Error() {
		t.Errorf("got error '%s' want '%s' ", got.Error(), want.Error())
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, keyword, definition string) {
	t.Helper()
	got, err := dictionary.Search(keyword)
	want := definition

	if err != "" {
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
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		keyword := "test"
		definition := "this is just a test"

		err := dictionary.Add(keyword, definition)

		assertErr(t, err, DictionaryErr(""))
		assertDefinition(t, dictionary, keyword, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertErr(t, err, errWordAlreadyExist)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		keyword := "test"
		definition := "this is just a test"
		dictionary := Dictionary{keyword: definition}
		newDefinition := "new definition"

		err := dictionary.Update(keyword, newDefinition)

		assertErr(t, err, "")
		assertDefinition(t, dictionary, keyword, newDefinition)
	})

	t.Run("new word update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErr(t, err, errKeywordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	keyword := "test"
	dictionary := Dictionary{keyword: "test definition"}

	dictionary.Delete(keyword)

	_, err := dictionary.Search(keyword)

	assertErr(t, err, errKeywordNotFind)
}
