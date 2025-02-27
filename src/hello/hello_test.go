package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func (t *testing.T, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func (t *testing.T)  {
		got := Hello("Lotiger", "")
		want := "Hello, Lotiger"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func (t *testing.T)  {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func (t *testing.T)  {
		got := Hello("Lotiger", "Spanish")
		want := "Hola, Lotiger"
		assertCorrectMessage(t, got, want)
	})
}