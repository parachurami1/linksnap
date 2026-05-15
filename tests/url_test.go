package tests

import (
	"linksnap/service"
	"testing"
)

func TestSlugGenerator(t *testing.T) {
	want := 5
	got := len(service.Shorten("https://youtube.com"))
	if want != got {
		t.Errorf("want %v got %v \n", want, got)
	}
	print("\n\n")
}
