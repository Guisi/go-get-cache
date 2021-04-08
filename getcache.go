package go_get_cache

import (
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"time"
)

/*
Change the implementation below to introduce a cache with a configurable lifetime so that;

- When there is a cache miss the Get receiver will fetch and cache the data.
- When there is a cache hit that data is returned

Notes

- No public receivers should be added to UrlGetter
- A constructor function may be added for a UrlGetter and fields may be added to UrlGetter
- Additional interfaces and types, public or private, may be added.
- Submission must include unit tests.
- External Packages may be imported

*/

type UrlGetter interface {
	Get(url string) ([]byte, error, bool)
}

type UrlGetterImpl struct {
	Storage *cache.Cache
}

func NewUrlGetter(expiration time.Duration) UrlGetter {
	return &UrlGetterImpl{
		Storage: cache.New(expiration, cache.NoExpiration),
	}
}

func (urlGetter *UrlGetterImpl) Get(url string) ([]byte, error, bool) {
	cached, found := urlGetter.Storage.Get(url)
	if found {
		return cached.([]byte), nil, found
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, err, false
	}

	response, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err, false
	}

	urlGetter.Storage.Set(url, response, cache.DefaultExpiration)

	return response, err, false
}
