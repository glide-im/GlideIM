package mq_nsq

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type NSQConsumerConfig struct {
	Topic       string
	Channel     string
	NsqLookupds []string
}

type NSQConsumer struct {
	c    *nsq.Consumer
	conf *NSQConsumerConfig
}

func NewConsumer(c *NSQConsumerConfig) (*NSQConsumer, error) {
	config := nsq.NewConfig()
	config.LookupdPollTimeout = 15 * time.Second
	consumer, err := nsq.NewConsumer(c.Topic, c.Channel, config)
	if err != nil {
		return nil, err
	}
	n := &NSQConsumer{
		c:    consumer,
		conf: c,
	}
	return n, nil
}

func (c *NSQConsumer) Connect() error {
	return c.c.ConnectToNSQLookupds(c.conf.NsqLookupds)
}

func (c *NSQConsumer) AddHandler(handler nsq.Handler) {
	c.c.AddHandler(handler)
}
