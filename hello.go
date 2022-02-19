package main

import (
	"bufio"
	"caihp/hellogo/wcfhelloservice"
	"caihp/hellogo/zaplog"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	slog "log"
	"os"
	"time"
)

var logger *zaplog.Logger = nil

func main() {
	fmt.Println("hello go")
	slog.Println("hello go")

	/* desc: color log function */
	// color.Green.Printf("teststring\r\n")
	// color.Bold.Printf("teststring\r\n")

	/* desc: zap log function */
	zaplog.Init()
	logger = zaplog.GetInitLogger()
	if logger == nil {
		logger.Infof("Failed to init log")
		return
	}
	logger.Infof("Succeed to init log")
	logger.Debugf("Succeed to init log")
	//logger.Warnf("Succeed to init log")
	//logger.Errorf("Succeed to init log")go hel

	/* desc: color plus zap log function */
	// logger.Infof(color.Green.Sprintf("Succeed to init log"))

	/* desc: wcf client function */
	// WcfExampleBasicUsage()
	WcfExampleWithOptions()

	bufio.NewReader(os.Stdin).ReadByte()
}

func WcfExampleBasicUsage() {
	client := soap.NewClient("http://localhost:8080/HelloToYou/Svc")
	service := wcfhelloservice.NewIWCFService(client)

	str := "czy1"
	reply, err := service.SayHello(&wcfhelloservice.SayHello{xml.Name{"http://tempuri.org/", "SayHello"}, &str})
	if err != nil {
		logger.Errorf("could't get reply: %v", err)
	}
	logger.Infof("result = %s", *reply.SayHelloResult)
}

func WcfExampleWithOptions() {
	client := soap.NewClient("http://localhost:8080/HelloToYou/Svc",
		soap.WithTimeout(time.Second * 5),
		soap.WithBasicAuth("usr", "psw"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
	service := wcfhelloservice.NewIWCFService(client)

	for {
		str := getInputString ()
		// displayInfo(input)
		if str == "exit" {
			break
		}
		callwcfservice(service, &str)
	}
}

func getInputString () string {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if len(input) > 1 && input[len(input) - 2] == '\r' {
		input = input[:len(input) - 2]
	} else {
		input = input[:len(input) - 1]
	}
	return input
}

func callwcfservice (service wcfhelloservice.IWCFService, paramter *string) {
	reply, err := service.SayHello(&wcfhelloservice.SayHello{xml.Name{"http://tempuri.org/", "SayHello"}, paramter})
	if err != nil {
		logger.Errorf("could't get reply: %v", err)
	}
	logger.Infof("result = %s", *reply.SayHelloResult)
}

func displayInfo(s string) {
	fmt.Printf("----------------------------\r\n")
	fmt.Printf("Input command: %s\r\n", s)
	fmt.Printf("Input command: %q\r\n", s)
	fmt.Printf("Input command: %q\r\n", []rune(s))
	fmt.Printf("Input command: %x\r\n", []rune(s))
	fmt.Printf("Input command: %q\r\n", []byte(s))
	fmt.Printf("----------------------------\r\n")
}