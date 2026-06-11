package initialize

import (
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/ve-weiyi/ve-blog-gin/config"
)

// RabbitMQ 创建 RabbitMQ 连接，带自动重连
func RabbitMQ(cfg *config.Config) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.Username, cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host, cfg.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("connect rabbitmq failed: %v", err)
	}

	rmq := &rabbitMQConn{
		conn:    conn,
		url:     url,
		closeCh: make(chan struct{}),
	}

	// 启动断线重连监听
	go rmq.watchAndReconnect()

	return conn
}

type rabbitMQConn struct {
	conn    *amqp.Connection
	url     string
	mu      sync.RWMutex
	closeCh chan struct{}
	closed  bool
}

// watchAndReconnect 监听连接关闭并自动重连
func (r *rabbitMQConn) watchAndReconnect() {
	for {
		r.mu.RLock()
		if r.conn == nil {
			r.mu.RUnlock()
			return
		}
		notifyClose := r.conn.NotifyClose(make(chan *amqp.Error, 1))
		r.mu.RUnlock()

		select {
		case err := <-notifyClose:
			if err != nil {
				log.Printf("[RabbitMQ] connection closed: %v, reconnecting...", err)
				r.tryReconnect()
			}
		case <-r.closeCh:
			return
		}
	}
}

// tryReconnect 尝试重连
func (r *rabbitMQConn) tryReconnect() {
	for {
		r.mu.RLock()
		closed := r.closed
		r.mu.RUnlock()
		if closed {
			return
		}

		conn, err := amqp.Dial(r.url)
		if err != nil {
			log.Printf("[RabbitMQ] reconnect failed: %v, retry in 5s", err)
			time.Sleep(5 * time.Second)
			continue
		}

		r.mu.Lock()
		r.conn = conn
		r.mu.Unlock()
		log.Printf("[RabbitMQ] reconnected successfully")
		return
	}
}
