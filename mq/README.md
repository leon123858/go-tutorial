# rabbit mq 練習題

## 題目

題目:
使用 Go 語言和 RabbitMQ 實現一個簡單的訊息佇列系統。該系統包含一個生產者 (Producer) 和一個消費者 (Consumer)。生產者會產生隨機的整數,並將其發送到 RabbitMQ 的佇列中。消費者會從佇列中獲取訊息,並計算所有接收到的整數的總和。
要求:

使用 Go 語言的 RabbitMQ 客戶端庫 (如 amqp)。
生產者每秒產生一個隨機整數,範圍在 1 到 100 之間。
消費者從佇列中獲取訊息,並實時計算總和。
當消費者接收到的整數總和超過 1000 時,顯示總和並退出程式。
使用 Go 的 goroutine 和 channel 實現生產者和消費者的並發通訊。

提示:

使用 amqp.Dial 建立與 RabbitMQ 的連線。
使用 channel.QueueDeclare 宣告一個佇列。
生產者使用 channel.Publish 發送訊息到佇列。
消費者使用 channel.Consume 從佇列中獲取訊息。
使用 json.Marshal 和 json.Unmarshal 對訊息進行序列化和反序列化。

## 安裝指令

can view web on [rabbitmq](http://localhost:15672/)

```shell
docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```