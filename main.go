package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
    client, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/" + os.Getenv("INFURA_ID"))

    if err != nil {
		log.Fatalf("Oops! There was a problem", err)
    } else {
		fmt.Println("Sucess! you are connected to the Ethereum Network")
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		fmt.Println(header.Number.String())
	}
}