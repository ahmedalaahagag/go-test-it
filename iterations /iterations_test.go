package iterations

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Repeat a char 5 times", func(t *testing.T) {
		got := repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(got, want)
	})

	t.Run("Repeat a char 0 times", func(t *testing.T) {
		got := repeat("a", 0)
		want := ""
		assertCorrectMessage(got, want)
	})
}
