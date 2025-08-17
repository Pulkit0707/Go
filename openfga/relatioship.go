// Add relatioships

package main

import(
	"context"
	"log"

	openfga "github.com/openfga/go-sdk"
)

func WriteRelationship(ctx context.Context, client *openfga.SdkClient, storeID string){
	_,_,err:=client.Write(ctx, storeID).
	Body(openfga.WriteRequest{
		Writes: &openfga.TupleKeys{
			TupleKeys: []openfga.TupleKey{
				{
					User: "user:1",
					Relation: "reader",
					Object: "document:doc1",
				},
			},
		},
	}).
	Execute()
	if err!=nil{
		log.Fatal("Error writing relatioship:", err)
	}
}