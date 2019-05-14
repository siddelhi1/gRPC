package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	//	http.HandleFunc("/todo", addToDo)
	//	http.HandleFunc("/updatetodo/{id}", UpdateToDo)
	//	http.HandleFunc("/deleteall", DeleteAll)

	log.Fatal(http.ListenAndServe(":6060", nil))
}

/*
func addToDo(w http.ResponseWriter, r *http.Request) {
	address := flag.String("server", "http://localhost:8080", "HTTTP gateway url, e.g. http://localhost:8080")

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	var body string

	resp, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(`
	{
		"api":"v1",
		"toDo": {
			"title":"title (%s)",
			"description":"description (%s)",
			"reminder":"%s"
		}
	}
`, pfx, pfx, pfx)))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

}

//UpdateToDo Function will update the TODO list
func UpdateToDo(w http.ResponseWriter, r *http.Request) {
	address := flag.String("server", "http://localhost:8080", "HTTTP gateway url, e.g. http://localhost:8080")
	var body string
	fmt.Println("GET params were:", r.URL.Query())

	// if only one expected
	id := r.URL.Query().Get("id")
	resp, err := http.Get(fmt.Sprintf("%s%s/%s", *address, "/v1/todo/{}", id))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Read response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	address := flag.String("server", "http://localhost:8080", "HTTTP gateway url, e.g. http://localhost:8080")
	var body string

	req, err := http.NewRequest("DELETEALL", fmt.Sprintf("%s%s/%s", *address, "/v1/todo/all"), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call DeleteAll method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Delete response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Delete response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

}
*/
