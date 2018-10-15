package main

import (
	"github.com/34blast/gofantasyfootball/common"
	"fmt"
	//	"github.com/zemirco/couchdb"
	//	"net/url"
	"time"
	//	"github.com/couchbase/gocb"
	// "gopkg.in/couchbase/gocb.v1"
)

func main() {
	fmt.Println("populateplayers.main() : starting exectuion")
	fmt.Println()
	start := time.Now()

	/*
	password := ""
	cluster, _ := gocb.Connect("couchdb://127.0.0.1")
	bucket, _ := cluster.OpenBucket("fantasybucket", password)

	fmt.Println(bucket)
		u, err := url.Parse("http://127.0.0.1:5984/")
		if err != nil {
			panic(err)
		}

		// create a new client
		client, err := couchdb.NewClient(u)
		if err != nil {
			panic(err)
		}

		// get some information about your CouchDB
		info, err := client.Info()
		if err != nil {
			panic(err)
		}
		fmt.Println(info)

		// create a database
		if _, err = client.Create("player"); err != nil {
			panic(err)
		}
	*/
	var ajgreen common.Player
	ajgreen.SetAll(common.WR, "Cincinnati", 99)
	ajgreen.SetLastName("Green")
	ajgreen.SetFirstName("AJ")
	fmt.Println("ajgreen info =  ", ajgreen.String())
	fmt.Println()
	
	var tePerson common.Person
	tePerson.SetFirstName("Tyler")
	tePerson.SetLastName("Eiffert")
	teifert := common.CreatePlayer(common.TE, "Cincinnati", 79, tePerson)
	fmt.Println("teifert info =  ", teifert.String())
	fmt.Println()

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println("elapsed time to execute: ", elapsed)
	fmt.Println()
	fmt.Println("populateplayers.main() : ending exectuion")
}
