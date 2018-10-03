package main

import (
	"fmt"
//	"github.com/couchbase/gocb"
//	"gopkg.in/couchbase/gocb.v1"
)

var (
	cbConnStr  = "couchbase://localhost"
	cbBucket   = "example"
	cbPassword = ""
	jwtSecret  = []byte("UNSECURE_SECRET_TOKEN")
)

func main() {
	fmt.Println("createbuckets.main() Starting execution ")
/*

	bucketSettings := gocb.BucketSettings{
		FlushEnabled:  true,
		IndexReplicas: true,
		Name:          "fantasybucket",
		Password:      "",
		Quota:         120,
		Replicas:      1,
		Type:          gocb.Couchbase}

	myCluster, connectError := gocb.Connect(cbConnStr)
	if connectError != nil {
		panic(connectError)
	}
	fmt.Println("createbuckets.main() connected sucessfully ")

	clusterManager := myCluster.Manager("admin", "admin")

	error := clusterManager.InsertBucket(&bucketSettings)
	if error != nil {
		panic(error)
	}

	fmt.Println("createbuckets.main() inserted a bucket okay ")
	
	
	buckets, error := clusterManager.GetBuckets()
	if error != nil {
		panic("ERROR getting buckets")
	}
	for _, bucketData := range buckets {
		fmt.Println("BucketName = ", bucketData.Name)
//		bucketData.Manager("", "").CreatePrimaryIndex("", true, false)
	}
	
	*/

	fmt.Println("")
	fmt.Println("createbuckets.main() ending  execution ")

}
