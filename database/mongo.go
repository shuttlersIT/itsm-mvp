package database

import (
	"fmt"

	"github.com/shuttlersIT/itsm-mvp/structs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDBConnection Encapsulates a connection to a database.
type MongoDBConnection struct {
	session *mgo.Session
}

// SaveUser register a user so we know that we saw that user already.
func (mdb MongoDBConnection) SaveUser(u *structs.User) error {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	if _, err := mdb.LoadUser(u.Email); err == nil {
		return fmt.Errorf("user already exists")
	}
	c := mdb.session.DB("portalsession").C("users")
	err := c.Insert(u)
	return err
}

// LoadUser get data from a user.
func (mdb MongoDBConnection) LoadUser(Email string) (result structs.User, err error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("portalsession").C("users")
	err = c.Find(bson.M{"email": Email}).One(&result)
	return result, err
}

// GetSession return a new session if there is no previous one.
func (mdb *MongoDBConnection) GetSession() *mgo.Session {
	if mdb.session != nil {
		return mdb.session.Copy()
	}
	session, err := mgo.Dial("mongodb://mongo:27017/portalsession")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
