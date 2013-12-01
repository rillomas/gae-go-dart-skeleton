package xclamm

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"encoding/json"
	"io/ioutil"
)


const (
	visitorInfoKind string = "visitorInfo"
	visitorInfoKey string = "theVisitorInfoKey"
)

type VisitorInfo struct {
	Count int // total number of visitors
}

// setup handlers
func init() {
	http.HandleFunc("/api/1/visitorInfo", handleVisitorInfoRequest)
}

func handleVisitorInfoRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Debugf("Method: %s Url:%s ContentLength: %d\n", r.Method, r.URL, r.ContentLength)
	if (r.Method == "GET") {
		getVisitorInfo(c,w,r)
	} else if (r.Method == "POST") {
		setVisitorInfo(c,w,r)
	}
}


func getVisitorInfo(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	k := datastore.NewKey(c,visitorInfoKind, visitorInfoKey, 0, nil)
	info := new(VisitorInfo)
	err := datastore.Get(c,k, info)
	// ignore the NoSuchEntityError and return the default value if we don't have the entity stored yet
	if err != nil && err != datastore.ErrNoSuchEntity {
		c.Errorf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outBuf, err := json.Marshal(info)
	if err != nil {
		c.Errorf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	size, err := w.Write(outBuf)
	if err != nil {
		c.Errorf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Debugf("Wrote %d bytes as response\n", size)

}

func setVisitorInfo(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.Errorf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var info VisitorInfo
	json.Unmarshal(buf, &info)
	k := datastore.NewKey(c,visitorInfoKind, visitorInfoKey, 0, nil)

	_,err = datastore.Put(c,k,&info)
	if err != nil {
		c.Errorf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
