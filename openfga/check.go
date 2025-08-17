// queries

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/smithy-go/document"
	openfga "github.com/openfga/go-sdk"
)

func CheckAccess(ctx context.Context, client *openfga.SdkClient, storeID string){
	checkResp,_,err:=client.Check(ctx,storeID).
	Body(openfga.CheckRequest{
		User: "user:1",
		Relation: "reader",
		Object: "document:doc1",
	}).
	Execute()
	if err!=nil{
		log.Fatal("Error checking access:", err)
	}
	fmt.Printf("Can user:1 read document:doc1? %v\n", checkResp.Allowed)
}