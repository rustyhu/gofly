package someframes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Act struct {
	Action  string
	Payload string
}

type ActAdd struct {
	Action  string
	Payload Info
}

type Info struct {
	Id   string `json:"ID"`
	Name string `json:"Name"`
	Area int32  `json:"Area"`
}

var store = map[string]Info{}

// test cases
var addtxt = `{
	"action": "add",
	"payload": {
		"ID": "mm886",
		"Name": "WestL",
		"Area": 3300
	}
}`

// one liner { "action": "add", "payload": { "ID": "mm886", "Name": "WestL", "Area": 3300 } }

var shorttxt = `{
	"action": "get",
	"payload": "mm886"
}`

func hpost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n##### hpost #####")

	reqjson := ActAdd{}
	var err error
	// c-style once do-while{} error process to avoid `goto`
	for i := 0; i < 1; i++ {
		reqbyte, err := io.ReadAll(r.Body)
		if err != nil {
			break
		}

		err = json.Unmarshal(reqbyte, &reqjson)
		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Println(reqjson)
		pl := &reqjson.Payload
		fmt.Printf("Get payload: %#v\n", *pl)
		store[pl.Id] = *pl
	}
}

func hget(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n##### hget #####")

	reqbyte, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	reqjson := map[string]string{}
	json.Unmarshal(reqbyte, &reqjson)

	fmt.Println("Show request:", reqjson)
	id := reqjson["payload"]
	if lk, ok := store[id]; ok {
		info := fmt.Sprintf("hget id: %v, content: %#v\n", id, lk)
		fmt.Println(info)
		io.WriteString(w, info)
	} else {
		fmt.Printf("hget id: %v - no this element\n", id)
		http.Error(w, "", http.StatusBadRequest)
	}
}

func hdelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n##### hdelete #####")

	reqbyte, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	reqjson := map[string]string{}
	json.Unmarshal(reqbyte, &reqjson)

	fmt.Println("Show request:", reqjson)
	id := reqjson["payload"]
	fmt.Printf("Delete id: %v, content: %#v\n", id, store[id])
	delete(store, id)
}

func StartServer() {
	http.HandleFunc("/", hpost)
	http.HandleFunc("/g", hget)
	http.HandleFunc("/d", hdelete)
	http.ListenAndServe(":8080", nil)
}
