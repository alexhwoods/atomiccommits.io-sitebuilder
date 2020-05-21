package util

import "testing"

func TestInvertUrl(t *testing.T) {
	urls := []string{
		"https://www.atomiccommits.io/joining-streams/",
		"http://www.atomiccommits.io/inline-functions/",
		"https://vim-adventures.com",
		"https://mail.google.com/mail/u/0/#inbox",
	}

	invertedUrls := []string{
		"io.atomiccommits.www/joining-streams/",
		"io.atomiccommits.www/inline-functions/",
		"com.vim-adventures",
		"com.google.mail/mail/u/0/#inbox",
	}

	for i, url := range urls {
		if InvertUrl(url) != invertedUrls[i] {
			t.Errorf("Url " + url + " should invert to " + invertedUrls[i] + ", but it inverts to " + InvertUrl(url))
		}
	}
}
