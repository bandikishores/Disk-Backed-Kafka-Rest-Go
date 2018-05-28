package dependency

import (
	"github.com/beeker1121/goque"
	"io/ioutil"
	"log"
	"time"
	"github.com/Shopify/sarama"
)

var q *goque.Queue

const time_in_ms = 1000

const dirQueuePath = "/Users/bandi.kishore/test/diskqueue/"

func InitDiskQueue() {
	var err error;
	q, err = goque.OpenQueue(dirQueuePath)
	if(err != nil) {
		log.Printf("Error occured while creating disk backed Queue at location %s with err %s", dirQueuePath, err)
		return
	}
	go startDeQueue()
}
	
func startDeQueue() {
	ticker := time.NewTicker(time.Millisecond*time_in_ms)
	for {
		select {
			case <-ticker.C:
				if(q.Length() == 0) {
					log.Print("No topic Found")
				} else {
					topic, err := q.Dequeue()
		            if(err != nil) {
			            log.Printf("%s Error occured while reading topic from file", err)
		            }
					if(q.Length() == 0) {
						log.Print("No Data Found for Topic %s", topic)
						return
					}
					
		            data, err := q.Dequeue()
		            if(err != nil) {
			            log.Printf("%s Error occured while reading data from file", err)
		            }
		            if(topic != nil && data != nil) {
			            sendToKafka(topic.ToString(), data.Value)
		            } else {
			            	log.Print("Either topic or data read was nil")
		            }
					
				}
		}
	}
}

// SendMessage Given a topic string and en event send it to the client
func sendMessage(topic string, data string) {
	// bData, _ := ioutil.ReadFile("/home/ubuntu/testgo/mytemp.json")
	bData, _ := ioutil.ReadFile("/Users/bandi.kishore/test/Test.json")
	// bData, _ := json.Marshal(mydata)
	// fmt.Print("Enter text: %s",bData)
	q.EnqueueString(topic)
	q.Enqueue(bData)
}

func sendToKafka(topic string, bData []byte) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bData),
	}
	writeToKafka(msg, producer)
}

func closeDiskQueue() error {
	return q.Close()
}