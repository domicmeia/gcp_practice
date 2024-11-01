package translation

import (
	"context"
	"fmt"

	"github.com/domicmeia/gcp_practice/config"
	"github.com/domicmeia/gcp_practice/handler/rest"
	"github.com/redis/go-redis/v9"
)

var _ rest.Translator = &Database{}

type Database struct {
	conn *redis.Client
}

func NewDatabaseServic(cfg config.Configuration) *Database {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DatabaseURL, cfg.DatabasePort),
		Password: "",
		DB:       0,
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Close() error {
	return s.conn.Close()
}

func (s *Database) Translate(word string, language string) string {
	out := s.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language))
	return out.Val()
}
