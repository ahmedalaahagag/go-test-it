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
		got := repeat("a")
		want := "aaaaa"
		assertCorrectMessage(got, want)
	})
}
