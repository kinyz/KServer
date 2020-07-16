package imongo

import "gopkg.in/mgo.v2"

type IMongo interface {
	Start()
	GetSession() *mgo.Session
	GetCollection(collName string) *mgo.Collection
}
