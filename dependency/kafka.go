package dependency

import (
	"github.com/Shopify/sarama"
	//"io/ioutil"
	"log"
)

var producer sarama.AsyncProducer

var (
	//	brokers = []string{"localhost:9093", "localhost:9094"}
	brokers = []string{"localhost:9092"}
)

func InitKafka() {
	var config = sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, _ = sarama.NewAsyncProducer(brokers, config)
	return
}

func setKafkaProducer(p sarama.AsyncProducer) {
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

func writeToKafka(msg *sarama.ProducerMessage, producer sarama.AsyncProducer) error {
	select {
	case producer.Input() <- msg:
	case err := <-producer.Errors():
		handleRecoverableError(err, "Error in sending message")
	}
	return nil
}

func closeKafka() error {
	return producer.Close()
}
