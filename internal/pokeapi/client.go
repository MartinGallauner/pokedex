package pokeapi

import (
	"github.com/martingallauner/pokedex/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(5 * time.Second),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
