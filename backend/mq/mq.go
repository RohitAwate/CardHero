package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type Task struct {
	ID uuid.UUID

	Op   uint16
	Args []string

	EnqueueTime    time.Time
	DequeTime      time.Time
	CompletionTime time.Time

	Success       bool
	StatusMessage string
}

var rdb *redis.Client
var ctx context.Context

var QueueName = "TasksQueue"

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	log.Println("Connected to Redis")
	ctx = context.Background()
}

type DispatchFunc func(task Task)

type MessageQueue struct {
	subscriptions map[uint16]DispatchFunc
}

func (mq *MessageQueue) Enqueue(task Task) error {
	jsonTask, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not marshal task while enqueueing")
	}

	err = rdb.Publish(ctx, QueueName, jsonTask).Err()
	return err
}

func (mq *MessageQueue) Subscribe(op uint16, df DispatchFunc) {
	mq.subscriptions[op] = df
}

func (mq *MessageQueue) StartListenAndDispatch() {
	sub := rdb.Subscribe(ctx, QueueName)
	defer func(sub *redis.PubSub) {
		err := sub.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(sub)

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(msg)
	}
}