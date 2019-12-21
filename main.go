package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ussd_callback(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	

	session_id := r.FormValue("sessionId")
	service_code := r.FormValue("serviceCode")
	phone_number := r.FormValue("phoneNumber")

	_ = fmt.Sprintf("%s,%s,%s",session_id,service_code,phone_number)
	text := r.FormValue("text")


	if len(text) == 0{
		w.Write([]byte("CON What would you want to check \n1. My Account \n2. My Phone Number"))
		return
	}else{
		switch text{

		case "1":
			w.Write([]byte("CON Choose account information you want to view \n1. Account Number\n2. Account Balance"))
			return

		case "2":
			w.Write([]byte(fmt.Sprintf("END Your Phone Number is %s",phone_number)))
			return

		case "1*1":
			w.Write([]byte("END Your Account Number is ACC1001"))
			return

		case "1*2":
			w.Write([]byte("END Your Balance is NGN 20,000"))
			return

		default:
			w.Write([]byte("END Invalid input"))
			return

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
	fmt.Println("this is a ussd application hosted using africastalking ussd platform, hosted on github")

	http.HandleFunc("/",ussd_callback)
	http.HandleFunc("/test",test)

	log.Fatal(http.ListenAndServe(":"+port,nil))
}