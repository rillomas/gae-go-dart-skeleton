package gae_go_dart_skeleton

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
	"storage" // path to local package
)

const (
	// rootDirectory string = "web/build/web" // root directory when deploying
	rootDirectory          string = "web/web" // root directory when developing
	leftTemplateDelimiter  string = "{{{"
	rightTemplateDelimiter string = "}}}"
	indexFileName          string = "index.html"
)

// Create a template with a custom delimiter
// because the default delimiter interferes with polymer's templating
var xt *template.Template = template.Must(
	template.New("skeleton").
		Delims(leftTemplateDelimiter, rightTemplateDelimiter).
		ParseFiles(
		fmt.Sprintf("%s/%s", rootDirectory, indexFileName)))

// setup handlers
func init() {
	http.HandleFunc("/", handleIndexPageRequest)
	http.HandleFunc("/api/1/visitorInfo", handleVisitorInfoRequest)
}

func handleIndexPageRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// set the initial visitor info to the web site by templating
	k := storage.GetVisitorInfoKey(c)
	info, err := storage.GetVisitorInfo(c, k)
	// ignore the NoSuchEntityError and return the default value if we don't have the entity stored yet
	if err != nil && err != datastore.ErrNoSuchEntity {
		log.Errorf(c, "Failed getting visitor info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outBuf, err := json.Marshal(info)
	if err != nil {
		log.Errorf(c, "Error while marshaling visitor info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	param := struct {
		VisitorInfo string
	}{
		string(outBuf),
	}
	err = xt.ExecuteTemplate(w, indexFileName, &param)
	if err != nil {
		log.Errorf(c, "Failed executing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleVisitorInfoRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if r.Method == "GET" {
		getVisitorInfo(c, w, r)
	} else if r.Method == "POST" {
		setVisitorInfo(c, w, r)
	}
}

func getVisitorInfo(c context.Context, w http.ResponseWriter, r *http.Request) {
	k := storage.GetVisitorInfoKey(c)
	info, err := storage.GetVisitorInfo(c, k)
	// ignore the NoSuchEntityError and return the default value if we don't have the entity stored yet
	if err != nil && err != datastore.ErrNoSuchEntity {
		log.Errorf(c, "Failed getting visitor info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(info)
	if err != nil {
		log.Errorf(c, "Failed encoding json response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setVisitorInfo(c context.Context, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var info storage.VisitorInfo
	err := dec.Decode(&info)
	if err != nil {
		log.Errorf(c, "Failed decoding user request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	k := storage.GetVisitorInfoKey(c)
	_, err = storage.SetVisitorInfo(c, k, &info)
	if err != nil {
		log.Errorf(c, "Failed setting visitor info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
