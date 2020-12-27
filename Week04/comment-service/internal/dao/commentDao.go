package dao

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week04/comment/internal/data"

	"github.com/Shopify/sarama"
)

type CommentDao struct {
	comment data.Comment
}

func (c *CommentDao) sendComment(com *data.Comment) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true

	msg := sarama.ProducerMessage{Topic: "Comment"}
	msg.Partition = int32(com.Platform)
	//msg.Key = sarama.StringEncoder(string(com.Platform))
	msg.Value = sarama.StringEncoder(com.Message)
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("Producer closed,err:", err)
		return err
	}
	defer producer.Close()

	pid, offset, err := producer.SendMessage(&msg)
	if err != nil {
		fmt.Println("Send massage field！err: ", err)
		return err
	}

	fmt.Printf("Partition=%d,Offset=%d,err:%v\n", pid, offset, err)
	return nil
}
