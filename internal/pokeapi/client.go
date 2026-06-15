package pokeapi

import (
	"net/http"
	"time"

	"github.com/AFEScalante/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}