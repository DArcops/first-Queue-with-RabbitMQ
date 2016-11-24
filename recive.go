package main

import (
  "fmt"
  "log"

  "github.com/streadway/amqp"

)


func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}


func main() {
  //al igual que en el sender debemos de conectar al amqp server:
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  //ch juega el papel del canal , al igual que un socket donde los datos viajan
  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()
  
  q, err := ch.QueueDeclare(
    "hello", // name
    false,   // durable
    false,   // delete when usused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")
  

  //Aqui se consume lo que haya dentro del canal con ch.Consume
  msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
  )
  failOnError(err, "Failed to register a consumer")

  forever := make(chan bool)

  go func() {
    for d := range msgs {
      log.Printf("Received a message: %s", d.Body)
    }
  }()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}
