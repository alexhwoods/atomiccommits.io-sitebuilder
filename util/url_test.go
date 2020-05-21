package util

import (
	"strings"
	"testing"
)

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

func TestUninvertUrl(t *testing.T) {
	invertedUrls := []string{
		"io.atomiccommits.www/joining-streams/",
		"io.atomiccommits.www/inline-functions/",
		"com.vim-adventures",
		"com.google.mail/mail/u/0/#inbox",
	}

	uninvertedUrls := []string{
		"www.atomiccommits.io/joining-streams/",
		"www.atomiccommits.io/inline-functions/",
		"vim-adventures.com",
		"mail.google.com/mail/u/0/#inbox",
	}

	for i, url := range invertedUrls {
		if InvertUrl(url) != uninvertedUrls[i] {
			t.Errorf("Url " + url + " should invert to " + uninvertedUrls[i] + ", but it inverts to " + InvertUrl(url))
		}
	}
}

func TestReversabilityOfInvertingUrls(t *testing.T) {
	urlWithHttps := "https://www.blog.sub.domain.really.overdoing.this.atomiccommits.io/foo/bar/baz/etc"

	url := strings.ReplaceAll(urlWithHttps, "https://", "")

	if InvertUrl(InvertUrl(urlWithHttps)) != url {
		t.Errorf("Inverting a url is not reversable.")
	}
}
