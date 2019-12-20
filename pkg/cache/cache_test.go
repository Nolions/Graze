package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Get(t *testing.T) {
	c := New()

	k, v := "test", "HelloWord"

	c.Set(k, []byte(v))
	tmp, err := c.Get(k)

	assert.Nil(t, err)
	assert.Equal(t, "HelloWord", string(tmp))
}

func TestSimpleCache_Del(t *testing.T) {
	c := New()

	k, v := "test", "HelloWord"

	c.Set(k, []byte(v))

	err := c.Del(k)

	tmp, err := c.Get(k)

	assert.Nil(t, tmp, err)
}