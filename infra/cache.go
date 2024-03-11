package infra

import (
	"app/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
)

type Cache struct {
	c      rueidis.Client
	adpter rueidiscompat.Cmdable
}

func (c *Cache) C() rueidis.Client {
	return c.c
}

func (c *Cache) Proxy() rueidiscompat.Cmdable {
	return c.adpter
}

func (c *Cache) Set(ctx *gin.Context, key string, val string) error {
	cmd := c.c.B().Set().Key(key).Value(val).Build()
	return c.c.Do(ctx, cmd).Error()
}

func (c *Cache) SetNx(ctx *gin.Context, key string, val string) error {
	cmd := c.c.B().Set().Key(key).Value(val).Build()
	return c.c.Do(ctx, cmd).Error()
}

func NewCache() *Cache {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{config.Conf.Cache.Server}})
	if err != nil {
		panic(err)
	}
	return &Cache{
		c:      client,
		adpter: rueidiscompat.NewAdapter(client),
	}
}
