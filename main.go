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
	
	



	// if len(text) == 0{

	// 	w.Write([]byte("CON Gplang Ussd \n1. Send Money\n2. Buy Airtime"))

	// }
	texts,ok := r.URL.Query()["text"]
	if !ok || len(texts[0]) < 1 {
		w.Write([]byte("CON Gplang Ussd \n1. Send Money\n2. Buy Airtime"))
		return
	}
	text := texts[0]

		switch text{

		case "1":
			w.Write([]byte("CON Select Bank \n1. FBN\n2. UBA"))
			return

		case "2":
			w.Write([]byte("CON Select ISP \n1. MTN \n2. Airtel \n3. Glo"))
			return

		case "1*1":
			w.Write([]byte("END Success"))
			return

		case "1*2":
			w.Write([]byte("END Airtime purchase succesful"))
			return

		default:
			w.Write([]byte("END Invalid input"))
			return

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