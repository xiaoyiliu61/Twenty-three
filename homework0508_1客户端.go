package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client:=http.Client{}

	request,err:=http.NewRequest("GET","http://127.0.0.1:9000/userlogin",nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	response,err:=client.Do(request)
	defer response.Body.Close()
	if response.StatusCode==200 {
		body,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}
