package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type key_value_struct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var key_value_map map[int]string

func main() {
	mux1 := httprouter.New()
	mux2 := httprouter.New()
	mux3 := httprouter.New()
	key_value_map = map[int]string{
		1:  "a",
		2:  "b",
		3:  "c",
		4:  "d",
		5:  "e",
		6:  "f",
		7:  "g",
		8:  "h",
		9:  "i",
		10: "j",
	}

	mux1.GET("/keys/:key_id", getKeyById)
	mux1.GET("/keys", getKeys)
	mux1.PUT("/keys/:key_id/:value", putKeys)

	mux2.GET("/keys/:key_id", getKeyById)
	mux2.GET("/keys", getKeys)
	mux2.PUT("/keys/:key_id/:value", putKeys)

	mux3.GET("/keys/:key_id", getKeyById)
	mux3.GET("/keys", getKeys)
	mux3.PUT("/keys/:key_id/:value", putKeys)
	//mux.DELETE("/locations/:location_id", deleteLocations)

	go func() {
		server1 := http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: mux1,
		}
		server1.ListenAndServe()
	}()
	go func() {

		server2 := http.Server{
			Addr:    "0.0.0.0:3001",
			Handler: mux2,
		}
		server2.ListenAndServe()
	}()

	server3 := http.Server{
		Addr:    "0.0.0.0:3002",
		Handler: mux3,
	}
	server3.ListenAndServe()

}

func putKeys(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("hello")
	key_id := p.ByName("key_id")
	value := p.ByName("value")

	fmt.Println("map", key_value_map)
	key_id_int, err := strconv.Atoi(key_id)
	if err != nil {
		fmt.Println(err)
	}
	key_value_map[key_id_int] = value
	fmt.Println("map", key_value_map)
}

func getKeyById(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key_id := p.ByName("key_id")
	key_id_int, err := strconv.Atoi(key_id)
	if err != nil {
		log.Fatal(err)
	}

	key_value := key_value_map[key_id_int]

	key_value_json := key_value_struct{
		Key:   key_id_int,
		Value: key_value,
	}

	b, _ := json.Marshal(key_value_json)
	fmt.Println(string(b))
	fmt.Fprintf(rw, "\n************************\n")
	fmt.Fprintf(rw, string(b))
	fmt.Fprintf(rw, "\n************************\n")

}

func getKeys(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	i := 0
	var key_value_struct_array []key_value_struct
	key_value_struct_array = make([]key_value_struct, len(key_value_map))
	for k, v := range key_value_map {
		key_value_struct_array[i] = key_value_struct{
			Key:   k,
			Value: v,
		}
		i++
	}

	b, _ := json.Marshal(key_value_struct_array)
	fmt.Println(string(b))
	fmt.Fprintf(rw, "\n************************\n")
	fmt.Fprintf(rw, string(b))
	fmt.Fprintf(rw, "\n************************\n")

}
