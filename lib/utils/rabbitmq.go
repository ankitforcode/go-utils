package utils

import (
	"github.com/streadway/amqp"
)

func PubToExchange(conn *amqp.Connection, qName string, rkey string, b []byte, client []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	if err := ch.ExchangeDeclare(qName, "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.Publish(qName, rkey, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			AppId:        string(client),
			ContentType:  "application/json",
			Body:         b,
		}); err != nil {
		return err
	}
	return nil
}

func SubToExchange(conn *amqp.Connection, qName string, rkey string, f fn) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	if err := ch.ExchangeDeclare(qName, "topic", true, false, false, false, nil); err != nil {
		return err
	}
	q, err := ch.QueueDeclare(qName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := ch.QueueBind(q.Name, rkey, qName, false, nil); err != nil {
		return err
	}
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			if err := f(d.Body,
				[]byte(d.AppId)); err == nil {
				d.Ack(true)
			}
		}
	}()
	<-forever
	return nil
}

type fn func([]byte, []byte) error
