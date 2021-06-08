package client

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/agalue/onms-sink-send/protobuf/sink"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type KafkaProducer struct {
	Brokers       string
	Topic         string
	MaxBufferSize int
	producer      sarama.SyncProducer
}

func (cli *KafkaProducer) Connect() error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	brokerList := strings.Split(cli.Brokers, ",")
	var err error
	cli.producer, err = sarama.NewSyncProducer(brokerList, config)
	return err
}

func (cli *KafkaProducer) Close() {
	if cli.producer != nil {
		cli.producer.Close()
	}
}

func (cli *KafkaProducer) Publish(data []byte) error {
	if cli.producer == nil {
		return fmt.Errorf("please connect to Kafka")
	}
	totalChunks := cli.getTotalChunks(data)
	var chunk int32
	id := uuid.New().String()
	for chunk = 0; chunk < totalChunks; chunk++ {
		bytes := cli.wrapMessageToSink(id, data, chunk, totalChunks)
		partition, offset, err := cli.producer.SendMessage(&sarama.ProducerMessage{
			Topic: cli.Topic,
			Key:   sarama.StringEncoder(id),
			Value: sarama.ByteEncoder(bytes),
		})
		if err != nil {
			return err
		}
		log.Printf("Message %d/%d with key %s sent to partition %d, offset %d on topic %s", chunk+1, totalChunks, id, partition, offset, cli.Topic)
	}
	return nil
}

func (cli *KafkaProducer) getTotalChunks(data []byte) int32 {
	if cli.MaxBufferSize == 0 {
		return int32(1)
	}
	chunks := int32(math.Ceil(float64(len(data) / cli.MaxBufferSize)))
	if len(data)%cli.MaxBufferSize > 0 {
		chunks++
	}
	return chunks
}

func (cli *KafkaProducer) getRemainingBufferSize(messageSize, chunk int32) int32 {
	if cli.MaxBufferSize > 0 && messageSize > int32(cli.MaxBufferSize) {
		remaining := messageSize - chunk*int32(cli.MaxBufferSize)
		if remaining > int32(cli.MaxBufferSize) {
			return int32(cli.MaxBufferSize)
		}
		return remaining
	}
	return messageSize
}

func (cli *KafkaProducer) wrapMessageToSink(id string, data []byte, chunk, totalChunks int32) []byte {
	bufferSize := cli.getRemainingBufferSize(int32(len(data)), chunk)
	offset := chunk * int32(cli.MaxBufferSize)
	msg := data[offset : offset+bufferSize]
	sinkMsg := &sink.SinkMessage{
		MessageId:          id,
		CurrentChunkNumber: chunk,
		TotalChunks:        totalChunks,
		Content:            msg,
	}
	bytes, err := proto.Marshal(sinkMsg)
	if err != nil {
		return []byte{}
	}
	return bytes
}
