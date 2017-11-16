package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"flag"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/julienschmidt/httprouter"
)

/*

	I would like to have an application that can get consul addres from command line, environment variable or use default
	and send specified message into application. Send message can be done as HTTP or Azure Service Bus message.

	First I can start with HTTP messages only. Interesting what would the best way to specify messages.

	It would be better to have some kind of UI application that can test message. Let's do this in the similar fashion as I did
	in reactive applications course. Hm... I can do routing on different endpoint. Something like http://localhost:4000/api/websocket

	For development purpose I probably can setup some lightweight reverse proxy and split UI served by webpack-dev-server and api served by go application.

	What can I use as the proxy here. NGINX probably would be the good choice. But I don't want to start this locally. Intersting how to do that from docker container.

	But I do want to use webpack-dev-server as development service tool.

	Ok, what feature I would like to start from:
	1. Check message pattern registration
	2. Get service IP addresses and state by name

*/

var tpl *template.Template

func main() {

	checkMessageCommand := flag.NewFlagSet("check-message", flag.ExitOnError)
	webCommand := flag.NewFlagSet("web", flag.ExitOnError)

	checkMessageText := checkMessageCommand.String("message", "", "Message json to check. (Required)")
	checMessageConsulAddress := checkMessageCommand.String("consul", "http://127.0.0.1:8500", "specify consul address with port")

	webCommandPort := webCommand.Int("port", 4000, "specify port to listen at")
	webCommandHost := webCommand.String("host", "localhost", "specify host to listen at")
	//webCommandConsulAddress := webCommand.String("consul", "http://127.0.0.1:8500", "specify consul address with port")

	if len(os.Args) < 2 {
		fmt.Println("web or subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "web":
		webCommand.Parse(os.Args[2:])
	case "check-message":
		checkMessageCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if webCommand.Parsed() {
		runWebListener(*webCommandHost, *webCommandPort)
	}

	if checkMessageCommand.Parsed() {
		if *checkMessageText == "" {
			checkMessageCommand.PrintDefaults()
			os.Exit(1)
		}
		checkMessage(*checMessageConsulAddress, *checkMessageText)
	}

}
func checkMessage(consulAddress string, messageText string) {
	log.Println("Checking message", messageText)

	messageChecker, err := NewMessageChecker(consulAddress)
	if err != nil {
		log.Fatalln(err)
	}

	info, err := messageChecker.GetMessageInfo(messageText)

	log.Printf("The message will be routed to: %+v; Service: %+v\r\n",
		info.MessageRegistration.ServiceName,
		strings.Join(mapSeviceInfosToStrings(info.ServiceInfos), ","))
}

func mapSeviceInfosToStrings(i []ServiceInfo) []string {
	a := make([]string, 0)
	for _, i := range i {
		a = append(a, i.String())
	}
	return a
}

func runWebListener(host string, port int) {
	templateContent, err := Asset("pub/index.gohtml")
	if err != nil {
		log.Fatalln("Error reading index.gohtml from assets", err)
	}

	tpl = template.Must(template.New("index.gohtml").Parse(string(templateContent)))

	mux := httprouter.New()

	mux.ServeFiles("/pub/*filepath", &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "pub",
	})
	mux.GET("/", index)
	mux.POST("/checkMessageRegistration", checkMessageRegistration)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("Listening on", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func checkMessageRegistration(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// check incomming message with consul registrations
	fmt.Fprintf(w, "This is the request, %+v!", req.FormValue("message"))
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
