package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	defaultExpirationTime = time.Hour
)

// Client used to make requests to redis
type Client struct {
	*redis.Client
	ttl       time.Duration
	namespace string
}

// Param is an optional param for redis client.
type Param func(*Client)

var redisClient *Client

// NewClient is a client constructor.
func NewClient(connectionURL, password, namespace string, params ...Param) *Client {

	c := redis.NewClient(&redis.Options{
		Addr:        connectionURL,
		Password:    "", // no password set
		DB:          0,
		DialTimeout: 15 * time.Second,
		MaxRetries:  10, // use default DB
	})

	// Test redis connection
	if _, err := c.Ping().Result(); err != nil {
		log.Panicf("unable to connect to redis: %s", err)
	}

	client := &Client{
		Client:    c,
		ttl:       defaultExpirationTime,
		namespace: namespace,
	}

	for _, applyParam := range params {
		applyParam(client)
	}

	setRedisClient(client)

	return client
}

func setRedisClient(client *Client) {
	redisClient = client
}

func RedisClient() *Client {
	return redisClient
}

func (c *Client) Ping() error {
	_, err := c.Client.Ping().Result()
	return err
}

func (c *Client) Set(key string, value interface{}, duration time.Duration) error {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Set(key, value, duration).Err()
}

func (c *Client) Get(key string) (interface{}, error) {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Get(key).Result()
}

func (c *Client) Delete(key string) (int64, error) {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Del(key).Result()
}

func (c *Client) Exists(key string) (bool, error) {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	i, err := c.Client.Exists(key).Result()
	return i >= 1, err
}

func (c *Client) SAdd(key string, value interface{}) error {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	_, err := c.Client.SAdd(key, value).Result()
	return err
}

func (c *Client) SDelete(key string) error {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	_, err := c.Delete(key)
	return err
}

func (c *Client) SRemove(key string, member interface{}) error {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	_, err := c.Client.SRem(key, member).Result()
	return err
}

func (c *Client) SMembers(key string) ([]string, error) {
	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.SMembers(key).Result()
}
