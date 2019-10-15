package main

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/awa/go-iap/appstore"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	client := appstore.New()
	client.ProductionURL = appstore.ProductionURL
	//client.SandboxURL = appstore.SandboxURL

	f, err := os.Open("./latestReceipts/latest")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 10000000)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	str := string(line)

	req := appstore.IAPRequest{
		ReceiptData: str,
		Password:    "",
	}

	var result map[string]interface{}
	err = client.Verify(context.Background(), req, &result)
	log.Println("error===============")
	log.Printf("%#v\n", err)

	log.Println("result===============")
	b, _ := json.Marshal(result)
	ioutil.WriteFile("./latestReceipts/latest.json", b, os.ModePerm)
}
