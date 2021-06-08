package main

import (
	"encoding/xml"
	"flag"
	"log"
	"time"

	"github.com/agalue/onms-sink-send/client"
	"github.com/agalue/onms-sink-send/model"
)

func main() {
	cli := new(client.KafkaProducer)

	flag.StringVar(&cli.Brokers, "brokers", "127.0.0.1:9092", "The Kafka Brokers to connect to, as a comma separated list")
	flag.StringVar(&cli.Topic, "topic", "OpenNMS.Sink.Events", "The Sink API topic for OpenNMS Events")
	flag.IntVar(&cli.MaxBufferSize, "bufferSize", 1024, "Maximum Sink Buffer Size")
	flag.Parse()

	if err := cli.Connect(); err != nil {
		panic(err)
	}
	defer cli.Close()

	event := &model.Event{
		UEI:       "uei.opennms.org/traps/sample",
		EventTime: &model.Time{Time: time.Now()},
		NodeID:    10,
		Interface: "127.0.0.1",
		Source:    "External",
		Severity:  "Warning",
		Mask: &model.Mask{
			Elements: []model.MaskElement{
				{
					MEname:  "id",
					MEvalue: []string{".1.2.3.4.5.6.7"},
				},
				{
					MEname:  "generic",
					MEvalue: []string{"6"},
				},
				{
					MEname:  "specific",
					MEvalue: []string{"1"},
				},
			},
		},
		LogMsg: &model.LogMsg{
			Destination: "donotpersist",
			Content:     "This is a test",
		},
		Descr: "This is a test",
		Parameters: &model.Parms{
			Params: []model.Param{
				{
					Name:  "owner",
					Value: "agalue",
				},
			},
		},
		AlarmData: &model.AlarmData{
			ReductionKey: "uei.opennms.org/traps/sample::10",
			AlarmType:    3,
		},
	}

	bytes, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		panic(err)
	}

	log.Printf("Event Payload: %s", string(bytes))
	if err := cli.Publish(bytes); err != nil {
		panic(err)
	}
}
