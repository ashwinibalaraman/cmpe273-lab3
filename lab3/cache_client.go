package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var circle_map map[string]interface{}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func add(node interface{}) {
	if x, ok := node.(string); ok {
		//circle_map.add
	}
}
func main() {
	fmt.Println(hash("Sharded data"))
}

func shardKeys(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
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

func sharding(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
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
