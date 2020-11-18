package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//专门往卡夫卡写日志的模块
var (
	client sarama.SyncProducer//生命一个全局的连接kafka的生产者
)
//初始化、连接kafka
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//连接kafka
	client,err = sarama.NewSyncProducer(addrs,config)
	if err != nil{
		fmt.Println("producer close,err:",err)
		return
	}
	return
}
//发送消息到kafka
func Sendmessage(topic, data string) (err error){
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic =topic
	msg.Value = sarama.StringEncoder(data)
	//发送到kafka
	pid,offset,err := client.SendMessage(msg)
	if err != nil{
		fmt.Println("send message failed,",err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
	return
}
