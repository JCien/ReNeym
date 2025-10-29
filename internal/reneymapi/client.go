package reneymapi

import (
	"time"

	"github.com/JesusC-XR/ReNeym/internal/reneymcache"
)

// Client
type Client struct {
	cache reneymcache.Cache
}

// NewClient
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: reneymcache.NewCache(cacheInterval),
	}
}
