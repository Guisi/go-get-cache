package go_get_cache

import (
	"bytes"
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {
	firstUrl := "http://www.blankwebsite.com"
	secondUrl := "https://www.blank.org"
	expiration := 2 * time.Second

	urlGetter := NewUrlGetter(expiration)

	resp, _, cached := urlGetter.Get(firstUrl)

	//first time cache should be empty, then is expected cached to be false
	if cached {
		t.Errorf("UrlGetter got %t, want %t", cached, false)
	}

	otherUrlResp, _, cached := urlGetter.Get(secondUrl)

	//when using a different URL, cached should be false
	if cached {
		t.Errorf("UrlGetter got %t, want %t", cached, true)
	}

	//and responses should be different
	if bytes.Compare(resp, otherUrlResp) == 0 {
		t.Error("FirstUrl response and secondUrl response should be different")
	}

	time.Sleep(1 * time.Second)

	cachedResp, _, cached := urlGetter.Get(firstUrl)

	//one second later, cache for the same URL should be valid, then is expected cached to be true
	if !cached {
		t.Errorf("UrlGetter got %t, want %t", cached, true)
	}

	//and response should be equal
	if bytes.Compare(resp, cachedResp) != 0 {
		t.Error("FirstUrl response and cached response should be equal")
	}

	time.Sleep(2 * time.Second)

	_, _, cached = urlGetter.Get(firstUrl)

	//two seconds later, cache for the same URL should have expired, then is expected cached to be false
	if cached {
		t.Errorf("UrlGetter got %t, want %t", cached, true)
	}

}
