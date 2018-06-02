package dependency

import (
	"github.com/Shopify/sarama"
	//"io/ioutil"
	"log"
)

var producer sarama.SyncProducer

var (
	//	brokers = []string{"localhost:9093", "localhost:9094"}
	brokers = []string{"localhost:9092"}
)

func InitKafka() {
	var err error
	var config = sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(brokers, config)
	if(err != nil) {
		log.Printf("Error initializing kafka %v", err)
	}
	return
}

func setKafkaProducer(p sarama.SyncProducer) {
	producer = p
}

func handleRecoverableError(err error, str string) {
	if err != nil {
		log.Printf("Sentry Event id %s", str)
		log.Printf("%s %#v %s", str, err, err.Error())
	}
}

// SendMessage Given a topic string and en event send it to the client
/*func sendMessage(topic string, data string) {
	// bData, _ := ioutil.ReadFile("/home/ubuntu/testgo/mytemp.json")
	bData, _ := ioutil.ReadFile("/Users/bandi.kishore/test/Test.json")
	// bData, _ := json.Marshal(mydata)
	// fmt.Print("Enter text: %s",bData)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bData),
	}
	writeToKafka(msg, producer)
}*/

func writeToKafka(msg *sarama.ProducerMessage, producer sarama.SyncProducer) error {
	defer func() {
        if err := recover(); err != nil {
	        	log.Printf("Recovering from Panic!! %s",err)
	        	handleKafkaMsgOnException(msg);
        }
    }()
	_,_,err := producer.SendMessage(msg);
	if(err != nil) {
		log.Print("Error while pushing to kafka")
		handleRecoverableError(err, "Error in sending message")
		handleKafkaMsgOnException(msg);
		return nil
	}
	log.Print("Message sent successfully to Kafka!!")
	/*select {
	case producer.Input() <- msg:
	case err := <-producer.Errors():
		handleRecoverableError(err, "Error in sending message")
	}*/
	return nil
}

func handleKafkaMsgOnException(msg *sarama.ProducerMessage) {
	b, err := msg.Value.Encode();
		if(err != nil) {
			log.Printf("Error occured while tring to get message from encode %v", err)
		} else {
			sendByteMessage(b);
		}
}

func closeKafka() error {
	return producer.Close()
}
