## URL Getter with a configurable lifetime cache

### How-to

````
import "github.com/Guisi/go-get-cache"

expiration := 30 * time.Second

urlGetter := go_get_cache.NewUrlGetter(expiration)

resp, err, cached := urlGetter.Get("www.sampleurl.com")
````