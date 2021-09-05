package cache_test

import (
	"testing"

	"github.com/stakkato95/lru_cache/cache"
	"github.com/stretchr/testify/assert"
)

func TestSingleItem(t *testing.T) {
	c := cache.NewCache(2)
	c.Set("key", "value")

	value, ok := c.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "value", value)

	c.Clear()

	value, ok = c.Get("key")
	assert.False(t, ok)
	assert.Nil(t, value)
}

func TestMultipleItems(t *testing.T) {
	// testData := map[string]string{
	// 	"k1": "v1",
	// 	"k2": "v2",
	// 	"k3": "v3",
	// }

	// c := cache.NewCache(3)

	// var lastKey string
	// for key, val := range testData {
	// 	assert.False(t, c.Set(key, val))
	// 	lastKey = key
	// }

	// for key, expectedVal := range testData {
	// 	actualVal, ok := c.Get(key)
	// 	assert.True(t, ok)
	// 	assert.Equal(t, expectedVal, actualVal)
	// }

	// newKey := "k4"
	// newVal := "v4"
	// assert.False(t, c.Set(newKey, newVal))

	// evictedVal, ok := c.Get(lastKey)
	// assert.False(t, ok)
	// assert.Nil(t, evictedVal)
}
