package main

import (
	"fmt"
	//"github.com/go-martini/martini"
	"net/http"
	"log"
	"io"
	"os"
	"encoding/json"
	//"bufio"
	"io/ioutil"
	//"bufio"
	//"bytes"
	"bytes"
)

//var m *martini.ClassicMartini


type TestData struct {
	Id int32 `json:"id"`
	Name string `json: "name"`
}

type TestProcessor struct {

}

func (h TestProcessor) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Called")
	fmt.Fprintf(w, "1 hello, you've hit %s\n", r.URL.Path)
}

func AddTest()  {

}

func logger(h http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main()  {
	fmt.Println("Test")
	//m = martini.Classic()
	//
	//r := martini.NewRouter()
	//r.Post(`/albums`, AddTest)
	//m.Action(r.Handle)
	//p := TestProcessor{}

	//h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "2 hello, you've hit %s\n", r.URL.Path)
	//})
	//
	//err1 := http.ListenAndServe(":9999", h)
	//fmt.Println(err1)
	//
	//
	//err2 := http.ListenAndServe(":9990", TestProcessor{})
	//fmt.Println(err2)

	//***************************

	//h := http.NewServeMux()
	//
	//h.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
	//	tmp := r.Method
	//	fmt.Println(tmp)
	//	fmt.Fprintln(w, "Hello, you hit foo!")
	//})
	//
	//h.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello, you hit bar!")
	//})
	//
	//h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(404)
	//	fmt.Fprintln(w, "You're lost, go home")
	//})
	//
	//hl := logger(h)
	//
	//err := http.ListenAndServe(":9999", hl)
	//log.Fatal(err)

	//************************

	//GetMethod()
	PostMethod()
}

func GetMethod()  {
	res, err := http.Get("http://localhost:50756/hello")

	if(err != nil){
		fmt.Println(err)
		return
	}
	responseData,err := ioutil.ReadAll(res.Body)
	responseString := string(responseData)
	fmt.Println(responseString)
	var yourStuff string

	data, err := io.Copy(os.Stdout, res.Body)
	json.NewDecoder(res.Body).Decode(&yourStuff)
	fmt.Println(res.Body)
	fmt.Println(data)
	fmt.Println(res)

	//unicode
	//response, err := http.Get("http://127.0.0.1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer response.Body.Close()
	//
	//scanner := bufio.NewScanner(response.Body)
	//scanner.Split(bufio.ScanRunes)
	//var buf bytes.Buffer
	//for scanner.Scan() {
	//	buf.WriteString(scanner.Text())
	//}
	//fmt.Println(buf.String())

	//reader := bufio.NewReader(res.Body)
	//for {
	//	line, _ := reader.ReadBytes('\n')
	//
	//
	//	log.Println(string(line))
	//}

}

func PostMethod()  {
var url = "http://localhost:50756/send"
	a := TestData{Id:1, Name:"Bob"}
	b, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	res := TestData{}
	json.Unmarshal(b, &res)
	//json.
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//http.Post("","application/json", bytes.NewBuffer())

}