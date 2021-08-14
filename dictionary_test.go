package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown key")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected error but did not get one")
		}

		assertErrorEq(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word, def := "test", "this is a test"
		err := dictionary.Add(word, def)

		assertErrorEq(t, err, nil)
		assertDef(t, dictionary, word, def)
	})
	t.Run("existing word", func(t *testing.T) {
		word, def := "test", "this is a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "some new def")

		assertErrorEq(t, err, ErrWordExists)
		assertDef(t, dictionary, word, def)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		// Should succeed
		word, def := "test", "just a test"
		newDef := "new def str"
		dictionary := Dictionary{word: def}
		err := dictionary.Update(word, newDef)

		assertErrorEq(t, err, nil)
		assertDef(t, dictionary, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		// Should fail
		word, def := "test", "just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, def)

		assertErrorEq(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "some val"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertErrorEq(t, err, ErrNotFound)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrorEq(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDef(t testing.TB, d Dictionary, word, def string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word : ", got)
	}
	if got != def {
		t.Errorf("got %q  want %q", got, def)
	}
}
