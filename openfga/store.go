// Create a store

package main

import(
	"context"
	"log"

	openfga "github.com/openfga/go-sdk"
)

func CreateStore(ctx context.Context, client *openfga.SdkClient){
	storeResp,_,err:= client.CreateStore(ctx).Body(openfga.CreateStoreRequest{
		Name: "eg-store",
	}).Execute()
	if err!=nil{
		log.Fatal("Error creating store:", err)
	}
	return storeResp.id
}