package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionnary := Dictionnary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionnary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionnary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
