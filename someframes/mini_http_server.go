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

/* test cases
{
	"action": "add",
	"payload": {
		"ID": "mm886",
		"Name": "WestL",
		"Area": 3300
	}
}

{ "action": "add", "payload": { "ID": "lk21", "Name": "EastL", "Area": 400 } }

{
	"action": "get",
	"payload": "mm886"
}
*/

// reqUnmarshalSucceed is a error processing wrapper, reqjson must be a pointer
// return value indicates whether to stop early
func reqUnmarshalSucceed(w http.ResponseWriter, r *http.Request, reqjson any) bool {
	reqbyte, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	err = json.Unmarshal(reqbyte, reqjson)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}

func hpost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n##### hpost #####")

	reqjson := ActAdd{}
	if !reqUnmarshalSucceed(w, r, &reqjson) {
		return
	}

	fmt.Println(reqjson)
	pl := &reqjson.Payload
	fmt.Printf("Get payload: %#v\n", *pl)
	store[pl.Id] = *pl
}

func hget(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n##### hget #####")

	reqjson := map[string]string{}
	if !reqUnmarshalSucceed(w, r, &reqjson) {
		return
	}

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

	reqjson := map[string]string{}
	if !reqUnmarshalSucceed(w, r, &reqjson) {
		return
	}

	fmt.Println("Show request:", reqjson)
	id := reqjson["payload"]
	fmt.Printf("Delete id: %v, content: %#v\n", id, store[id])
	delete(store, id)
}

func StartServer() {
	http.HandleFunc("/", hpost)
	http.HandleFunc("/g", hget)
	http.HandleFunc("/d", hdelete)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
