package go4it

import "testing"

func TestFiles(t *testing.T) {
	got := PWD()
	want := "/home/elanticrypt0/Documentos/__projects/go4it"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Logf("got %q, wanted %q", got, want)
	}
}
