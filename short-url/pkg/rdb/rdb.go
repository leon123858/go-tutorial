package rdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"short-url/pkg/mongo"
	"time"
)

type Client struct {
	rdb *goredislib.Client
	rs  *redsync.Redsync
	mdb *mongo.Client
	ctx context.Context
}

func NewClient(addr string, mdb *mongo.Client) *Client {
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	return &Client{
		rdb: client,
		rs:  rs,
		mdb: mdb,
		ctx: context.Background(),
	}
}

func (c *Client) Close() error {
	return c.rdb.Close()
}

func (c *Client) SetUrl(shortUrl string, url Url) error {
	// convert url to string
	jsonData, err := json.Marshal(url)
	if err != nil {
		return err
	}
	err = c.rdb.Set(c.ctx, shortUrl, jsonData, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUrl(shortUrl string) (Url, error) {
	val, err := c.rdb.Get(c.ctx, shortUrl).Result()
	if errors.Is(err, goredislib.Nil) {
		// key does not exist
		mutex := c.rs.NewMutex(shortUrl, redsync.WithExpiry(5*time.Second), redsync.WithTries(1))
		// try to acquire lock
		err := mutex.Lock()
		switch {
		case err == nil:
			{
				// lock acquired, do something here
				defer func(mutex *redsync.Mutex) {
					_, err := mutex.Unlock()
					if err != nil {
						// handle error
						println("lock can not unlock", err.Error())
					}
				}(mutex)
				// get data from mongo
				url, err := c.mdb.FindUrl(shortUrl)
				if err != nil {
					return Url{}, err
				}
				// set data to redis
				final := Url{
					Password: url.Password,
					LongURL:  url.LongURL,
				}
				err = c.SetUrl(shortUrl, final)
				if err != nil {
					return Url{}, err
				}
				return final, nil
			}
		case errors.Is(err, redsync.ErrFailed):
			// lock not acquired
			time.Sleep(3 * time.Second)
			// try again
			var innerErr error
			val, innerErr = c.rdb.Get(c.ctx, shortUrl).Result()
			if innerErr != nil {
				return Url{}, innerErr
			}
		}
		return Url{}, err
	} else if err != nil {
		return Url{}, err
	}
	var url Url
	err = json.Unmarshal([]byte(val), &url)
	if err != nil {
		return Url{}, err
	}
	return url, nil
}
