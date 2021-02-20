package main

import (
	"github.com/gomodule/redigo/redis"
)

type RoutingClient struct {
	// 接続
	conn redis.Conn
}

func CreateRoutingClient(redisAddr string) *RoutingClient {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	return &RoutingClient{conn: conn}
}

func (r *RoutingClient) Close() error {
	return r.conn.Close()
}

func (r *RoutingClient) GetAddr(subdomain string) (string, error) {
	return redis.String(r.conn.Do("GET",subdomain))
}

func (r *RoutingClient) SetAddr(subdomain string, addr string) error {
	_ , err := r.conn.Do("SET",subdomain,addr)
	return err
}