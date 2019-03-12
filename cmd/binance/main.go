package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"cryptocurrency/internal/app/binance"
	"cryptocurrency/pkg/util"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"sync"
)

var unsignedValue util.UnsignedString
var wg sync.WaitGroup

func init() {
	// load environment
	err := godotenv.Load("../../.env")
	util.CheckError(err)

	unsignedValue = util.UnsignedString{
		ApiKey: os.Getenv("binace_apiKey"),
		HashHandler: hmac.New(sha256.New, []byte(os.Getenv("binance_apiSecretKey"))),
	}
	binance.Ping()
}

func command() {
	//io.WriteString(os.Stdout, "please input any command: ")

	scanner := bufio.NewScanner(os.Stdin)
FINISH:
	for scanner.Scan() {
		switch scanner.Text() {
		case "server time":
			go binance.ServerTime()
		case "ping": go binance.Ping()
		case "account info": go binance.AccountInfo(unsignedValue)
		case "signed test": go binance.SignedTest()
		case "quit":
			break FINISH
		case "help":
			fmt.Println("commands: ping, server time, account info, signed test, quit")
			io.WriteString(os.Stdout, "please input any command: ")
		default:
			io.WriteString(os.Stdout, "unknown command!\n")
			io.WriteString(os.Stdout, "please input any command: ")
		}
	}
}

func main() {
	command()
	log.Println("Finished")
}
