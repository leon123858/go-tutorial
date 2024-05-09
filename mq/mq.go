package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// 建立與 RabbitMQ 的連線
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %s", err)
		}
	}(conn)

	// 建立一個 Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Fatalf("Failed to close channel: %s", err)
		}
	}(ch)

	// 宣告一個佇列
	q, err := ch.QueueDeclare(
		"sum_queue", // 佇列名稱
		false,       // 是否持久化
		false,       // 是否自動刪除
		false,       // 是否具有排他性
		false,       // 是否阻塞
		nil,         // 額外屬性
	)
	failOnError(err, "Failed to declare a queue")

	// 生產者
	go func() {
		for {
			// 產生隨機整數
			number := rand.Intn(100) + 1
			body, _ := json.Marshal(number)

			// 發送訊息到佇列
			err = ch.Publish(
				"",     // exchange
				q.Name, // 佇列名稱
				false,  // 是否強制
				false,  // 是否立即
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			failOnError(err, "Failed to publish a message")

			time.Sleep(1 * time.Second)
		}
	}()

	// 消費者
	msgs, err := ch.Consume(
		q.Name, // 佇列名稱
		"",     // 消費者
		true,   // 是否自動確認
		false,  // 是否具有排他性
		false,  // 是否不等待回覆
		false,  // 是否阻塞
		nil,    // 額外屬性
	)
	failOnError(err, "Failed to register a consumer")

	sum := 0
	for msg := range msgs {
		// 解析訊息內容
		var number int
		err := json.Unmarshal(msg.Body, &number)
		if err != nil {
			log.Printf("Failed to parse message: %s", err)
			continue
		}

		// 計算總和
		sum += number
		fmt.Printf("Received number: %d, Sum: %d\n", number, sum)

		// 總和超過 1000 時退出
		if sum > 1000 {
			break
		}
	}

	fmt.Printf("Total sum: %d\n", sum)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
