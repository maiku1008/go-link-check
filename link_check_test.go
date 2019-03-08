package main

import "testing"

// Incomplete
func TestLinkCheck(t *testing.T) {

    got := FollowUrl()
    want := "Let's follow some URLS!"

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}

func TestReadFromList(t *testing.T) {
    urls := []string{"google.com", "memesfunny.org", "sk1u.com"}

    got := ReadCSV("url_list.csv")
    want := urls[1]

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}
