package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./.config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {

	//CREATE DIAL
	// conConfig := app.ConnectionConfig{
	// 	Network: "tcp",
	// 	Address: viper.GetString("broker1.ip") + ":" + strconv.Itoa(viper.GetInt("broker1.port")),
	// }
	// con, err := app.NewConnection(conConfig)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	//DELETE TOPIC
	// if err := con.DeleteTopic([]string{"deneme", "test23"}); err != nil {
	// 	log.Fatal(err.Error())
	// }

	//CREATE TOPIC
	// createTopicConfig := app.TopicConfig{
	// 	Name:              "test",
	// 	Partitions:        6,
	// 	ReplicationFactor: 3,
	// }
	// if err := con.CreateTopic(createTopicConfig); err != nil {
	// 	log.Fatal(err.Error())
	// }

	//BROKERS
	// brokers := []string{
	// 	fmt.Sprintf("%s:%d", viper.GetString("broker1.ip"), viper.GetInt("broker1.port")),
	// 	fmt.Sprintf("%s:%d", viper.GetString("broker2.ip"), viper.GetInt("broker2.port")),
	// 	fmt.Sprintf("%s:%d", viper.GetString("broker3.ip"), viper.GetInt("broker3.port")),
	// }

	//PRODUCER
	//var ctx context.Context = context.Background()
	// config := app.Config{
	// 	Brokers: brokers,
	// 	Topic:   "Golangdan_Selamlar",
	// }
	// producerConfig := app.ProducerConfig{
	// 	CompressionCodec: kafka.Gzip,
	// }
	// producer := app.NewProducer(config, producerConfig)
	// record := kafka.Message{
	// 	Key:   []byte("Go Key 1"),
	// 	Value: []byte("Go Value 1"),
	// }
	// if err := producer.Produce(ctx, record); err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println("Veri gönderildi.")

	//CONSUMER
	// var ctx context.Context = context.Background()
	// config := app.Config{
	// 	Brokers: brokers,
	// 	Topic:   "fakedata",
	// }

	// consumerConfig := app.ConsumerConfig{
	// 	GroupID: "fakedata_group",
	// }

	// consumer, err := app.NewConsumer(config, consumerConfig)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// for {
	// 	msg, err := consumer.Consume(ctx)
	// 	if err != nil {
	// 		fmt.Println("Mesaj okunamadı")
	// 	}
	// 	fmt.Printf("Partition:%d,Key:%s,Value:%s\n", msg.Partition, msg.Key, msg.Value)
	// }

}
