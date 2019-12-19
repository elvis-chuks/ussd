package main

import (
	"fmt"
	"log"
	"net/http"
)

func ussd_callback(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	session_id := r.URL.Query()["sessionId"][0]
	service_code := r.URL.Query()["serviceCode"][0]
	phone_number := r.URL.Query()["phoneNumber"][0]
	text := r.URL.Query()["text"][0]

	_ = fmt.Sprintf("%s,%s,%s,%s",session_id,service_code,phone_number,text)

	fmt.Println(text)

	if len(text) == 0{

		w.Write([]byte("CON Gplang Ussd \n1. Send Money\n2. Buy Airtime"))

	}else{

		switch text{

		case "1":
			w.Write([]byte("CON Select Bank \n1. FBN\n2. UBA"))

		case "2":
			w.Write([]byte("CON Select ISP \n1. MTN \n2. Airtel \n3. Glo"))

		case "1*1":
			w.Write([]byte("END Success"))

		case "1*2":
			w.Write([]byte("END Airtime purchase succesful"))

		default:
			w.Write([]byte("END Invalid input"))

		}
	}
	

}

func main(){
	fmt.Println("this is a ussd application")

	http.HandleFunc("/",ussd_callback)

	log.Fatal(http.ListenAndServe(":5000",nil))
}