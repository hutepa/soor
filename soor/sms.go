package soor

import (
	"net/http"
//	"fmt"
//	"log"
//	"os"
//	"reflect"
)

func SendSMS(phone string,pin string) (resp *http.Response, err error) {
//recipients := [2]string{"96566096195","96560040171"}
//r := reflect.ValueOf(recipients)
	InitLogger()
	params := map[string]string{"username": "bwireless",
				     "password": "bwireless123",
				     "customerId": "998",
				     "senderText": "B.Wireless",
				     "messageBody": pin,
				     "recipientNumbers": phone,//r.String(),
				     "defdate": "",
				     "isBlink": "false",
				     "isFlash": "false" }

	req, err := http.NewRequest("GET", "https://www.smsbox.com/SMSGateway/Services/Messaging.asmx/Http_SendSMS", nil)
	if err != nil {
		//log.Print(err)
		//os.Exit(1)
		Error.Printf("%v\n",err)
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{

	}
	resp, err = client.Do(req)
	if err != nil {
		Error.Printf("%v\n",err)
	} else {
		Trace.Println(resp)
	}
	return resp, err

}