package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionnary := Dictionary{"Autologism": "A word that describes itself."}

	t.Run("Search for a known word.", func(t *testing.T) {
		got, _ := dictionnary.Search("Autologism")
		want := "A word that describes itself."

		assertStrings(t, got, want)
	})

	t.Run("Search for an unknown word.", func(t *testing.T) {
		_, got := dictionnary.Search("Poney")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Add a word and its definition to a dictionnary.", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "Heterologism"
		definition := "A word that is the opposite of its meaning."

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("Edd an existing word to the dictionnary.", func(t *testing.T) {
		word := "Heterologism"
		definition := "A word that is the opposite of its meaning."

		dictionnary := Dictionary{word: definition}
		err := dictionnary.Add(word, "Some random new definition.")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionnary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update an existing definition", func(t *testing.T) {
		word := "Heterologism"
		definition := "A word that is the opposite of its meaning."

		dictionary := Dictionary{word: definition}
		newDefinition := "A word which is not described by itself."

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "Heterologism"
		definition := "A word that is the opposite of its meaning."

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
