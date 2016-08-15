package mgodb

import (
  mg "gopkg.in/mgo.v2"
  "strings"
 )

// ServerExists checks the mongoDB server if it's up
// URL: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
// examples: localhost, mongodb://myuser:mypass@localhost:40001
// default port is 27017 
func ServerExists(url string) (*mg.Session, error) {
 session, err := mg.Dial(url)
 /*if err != nil {
  panic(err)
 }*/
 return session, err
}

// DbExists Checks the existance of a database 
func DbExists(s *mg.Session, db string) (bool) {
 names, err := s.DatabaseNames()
 if err != nil {
  return false
 }
 for _, value := range names {
  if (value == db){
   return true
  }
 }
 return false;
}

// CollectionExists Checks the existance of a collection
func CollectionExists(s *mg.Session, db string, clt string) (int){
 //get the database structure from the session
 dbase := s.DB(db)
 cltnames, err := dbase.CollectionNames()
 if err != nil {
  return -1
 }
 /* Here is the structure Collection. Use it to count the total number of documents
 type Collection struct {
    Database *Database
    Name     string // "collection"
    FullName string // "db.collection"
 }
 */
 fullName := strings.Join([]string{db,clt},".")
 for _, value := range cltnames {
  if (value == clt){
   myClt := mg.Collection{dbase, clt, fullName} //init the structure 
   sizeClt, _ := myClt.Count() //count the total number of documents
   return sizeClt
  }
 }
 return -1
}

