package helper

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

type Emp []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SetCache(key string, emp interface{}) bool {
	Cache.Set(key, emp, cache.NoExpiration)
	return true
}

func GetCache(key string) (Emp, bool) {
	var emp Emp
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(Emp)
	}
	return emp, found
}
