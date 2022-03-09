package sample

import (
	"fmt"

	"github.com/chenqichen/ksyun-ks3-go-sdk/ks3"
)

// BucketCORSSample shows how to get or set the bucket CORS.
func BucketCORSSample() {
	// New client
	client, err := ks3.New(endpoint, accessID, accessKey)
	if err != nil {
		HandleError(err)
	}

	// Create the bucket with default parameters
	err = client.CreateBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	rule1 := ks3.CORSRule{
		AllowedOrigin: []string{"*"},
		AllowedMethod: []string{"PUT", "GET", "POST"},
		AllowedHeader: []string{},
		ExposeHeader:  []string{},
		MaxAgeSeconds: 100,
	}

	rule2 := ks3.CORSRule{
		AllowedOrigin: []string{"http://www.a.com", "http://www.b.com"},
		AllowedMethod: []string{"GET"},
		AllowedHeader: []string{"Authorization"},
		ExposeHeader:  []string{"x-ks3-test", "x-ks3-test1"},
		MaxAgeSeconds: 100,
	}

	// Case 1: Set the bucket CORS rules
	err = client.SetBucketCORS(bucketName, []ks3.CORSRule{rule1})
	if err != nil {
		HandleError(err)
	}

	// Case 2: Set the bucket CORS rules. if CORS rules exist, they will be overwritten.
	err = client.SetBucketCORS(bucketName, []ks3.CORSRule{rule1, rule2})
	if err != nil {
		HandleError(err)
	}

	// Get the bucket's CORS
	gbl, err := client.GetBucketCORS(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Bucket CORS:", gbl.CORSRules)

	// Delete bucket's CORS
	err = client.DeleteBucketCORS(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Delete bucket
	err = client.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Println("BucketCORSSample completed")
}