package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var (
	streamName = "copper_deliver_stream"
	s = session.New(&aws.Config{Region: aws.String("us-east-2")})
	kc = firehose.New(s)
)

type RespStruct struct {
	RequestId int `json:"requestId"`
	Status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	resp := &RespStruct{ RequestId: 1, Status: "success" }
    resp2, err := json.Marshal(resp)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(resp2))

	if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("content-type", "application/json")
	w.Write(b)

	sendDataToKinesis([]byte(b))
}

func sendDataToKinesis(payloadData []byte) {
	req, resp := kc.PutRecordRequest(&firehose.PutRecordInput{
		DeliveryStreamName: &streamName,
		Record: &firehose.Record{ Data: payloadData },
	})
	err := req.Send()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", resp)
}

func main() {
	fmt.Println("Handling http://locahost:8080/")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
	

	// flag.Parse()
	// kc := kinesis.New(s)


	// startDestination := "1"
	// limit := int64(1)

	// streams, err1 := kc.DescribeDeliveryStream(&firehose.DescribeDeliveryStreamInput{
	// 	DeliveryStreamName: &streamName,
	// 	ExclusiveStartDestinationId: &startDestination,
	// 	Limit: &limit})
	// if err1 != nil {
	// 	panic(err1)
	// }
	// fmt.Printf("%v\n", streams)

	
}
