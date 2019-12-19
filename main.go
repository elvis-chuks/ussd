package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ussd_callback(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	fmt.Println(r.URL.Query())
	

	// session_id := r.URL.Query()["sessionId"][0]
	// service_code := r.URL.Query()["serviceCode"][0]
	// phone_number := r.URL.Query()["phoneNumber"][0]
	texts,ok := r.URL.Query()["text"]
	if !ok || len(texts[0]) < 1 {
		w.Write([]byte("invalid request"))
	}
	text := r.URL.Query()["text"][0]

	// _ = fmt.Sprintf("%s,%s,%s,%s",session_id,service_code,phone_number,text)



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

func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	w.Write([]byte("This App Works In Production"))
}

func main(){

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}
	fmt.Println("this is a ussd application")

	http.HandleFunc("/",ussd_callback)
	http.HandleFunc("/test",test)

	log.Fatal(http.ListenAndServe(":"+port,nil))
}