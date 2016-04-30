package storage

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	visitorInfoKind string = "visitorInfo"
	visitorInfoKey  string = "theVisitorInfoKey"
)

type VisitorInfo struct {
	Count int // total number of visitors
}

func GetVisitorInfoKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, visitorInfoKind, visitorInfoKey, 0, nil)
}

func GetVisitorInfo(c context.Context, k *datastore.Key) (*VisitorInfo, error) {
	info := new(VisitorInfo)
	err := datastore.Get(c, k, info)
	return info, err
}

func SetVisitorInfo(c context.Context, k *datastore.Key, info *VisitorInfo) (*datastore.Key, error) {
	return datastore.Put(c, k, info)
}
