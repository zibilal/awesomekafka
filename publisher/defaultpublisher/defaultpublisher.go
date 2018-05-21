package defaultpublisher

import (
	"github.com/Shopify/sarama"
	"fmt"
	"log"
	"errors"
)

type DefaultPublisher struct {
	topic string
	brokers []string
	config *sarama.Config
}

func NewDefaultPublisher(configure func() interface{},
	topic string, servers []string) (*DefaultPublisher, error) {

	p := new(DefaultPublisher)
	var ok bool
	p.config, ok = configure().(*sarama.Config)

	if !ok {
		return nil, errors.New("unexpected type, expected type is *sarama.Config")
	}
	p.topic = topic
	p.brokers = servers

	return p, nil
}

func (p *DefaultPublisher) Publish (key interface{}, data interface{}) error {

	bkey, okKey := key.([]byte)
	if !okKey {
		return fmt.Errorf("unexpected data type of key %T", key)
	}
	bdata, okData := key.([]byte)
	if !okData {
		return fmt.Errorf("unexpected data type of data %T", data)
	}
	var err error
	var producer sarama.SyncProducer

	if p.config != nil {
		producer, err = sarama.NewSyncProducer(p.brokers, p.config)
	} else {
		producer, err = sarama.NewSyncProducer(p.brokers, nil)
	}

	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key: sarama.ByteEncoder(bkey),
		Value: sarama.ByteEncoder(bdata),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	return nil
}


