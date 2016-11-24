package main

import (
  "fmt"
  "log"
  "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s",msg,err)
    panic(fmt.Sprintf("%s: %s",msg,err))
  }
}


func main() {
  //ahora conectamos con amqp server 
   conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
   failOnError(err, "Failed to connect to RabbitMQ")
   defer conn.Close()

  //una vez conectado al servidor se debe de abrir un canal por el cual viajaran los datos
   ch,err := conn.Channel()
   failOnError(err, "Failed to open a channel")
   defer ch.Close()

  //ya que tenemos el canal abierto debemos de crear la cola en la cual se reportaran los mensajes
   q, err := ch.QueueDeclare(
    "hello", // name
     false,   // durable
     false,   // delete when unused
     false,   // exclusive
     false,   // no-wait
     nil,     // arguments
   )
   failOnError(err, "Failed to declare a queue")

   body := "este es el cuerpo del mensaje a mandar"

   err = ch.Publish(
     "",     // exchange
     q.Name, // routing key
     false,  // mandatory
     false,  // immediate
     amqp.Publishing {
     ContentType: "text/plain",
     Body:        []byte(body),
  })

  failOnError(err, "Failed to publish a message")
  
}

