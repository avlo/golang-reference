package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/trinary"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// Sending address
const seed = trinary.Trytes("SIEGWQTEKSXYKGBVKSFTNVJEQZHSUZTJFOZYSEVFPAKTQLPAMGLTUKFNYLYKPDNSH9NXIJJS9BWQDOZZ9")

// Receiving address
const address = trinary.Trytes("WQBXUEJEVQERVOPNSGMASKBICMDNHDQLXETOPVULISQLIBUAMFOGTGHEJALDPGHQWX9ZTQCZFZACNHZTWA9KUWQDAD")

const minimumWeightMagnitude = 9
const depth = 3

var wg sync.WaitGroup

var topic string = "go-mqtt/sample"

func main() {
	var broker = "localhost"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
	opts.SetClientID("005_mqtt_event")
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	wg.Add(1)
	// go sub(client)
	go publish(client, broker)
	wg.Wait()
}

func publish(client mqtt.Client, args string) {
	argss := "http://" + args + ":14265"
	apiRef, err := api.ComposeAPI(api.HTTPClientSettings{URI: argss})
	must(err)

	num := 1000
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish(topic, 0, false, text)
		sendBundle(*apiRef)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sendBundle(apiRef api.API) {
	// Define an input transaction object
	// that sends 1 i to your new address
	transfers := bundle.Transfers{
		{
			Address: address,
			Value:   1,
		},
	}

	fmt.Println("Sending 1 i to " + address)

	trytes, err := apiRef.PrepareTransfers(seed, transfers, api.PrepareTransfersOptions{})
	must(err)

	myBundle, err := apiRef.SendTrytes(trytes, depth, minimumWeightMagnitude)
	must(err)

	fmt.Println(myBundle)
}

func sub(client mqtt.Client) {
	// defer wg.Done()
	token := client.Subscribe(topic, 1, nil)

	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
