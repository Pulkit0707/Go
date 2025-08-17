package main

import (
	"context"
	"fmt"
	"log"
)

func main(){
	ctx:=context.Background()

	// create openfga client
	client,err:= NewFGAClient()
	if err!=nil{
		log.Fatal("Error creating client",err)
	}

	// create a store
	storeId:=CreateStore(ctx,client)
	fmt.Println("Created store", storeId)

	// write model
	WriteMoedl(ctx,client,storeId)

	// add relationship
	WriteRelationship(ctx,client,storeId)

	// check relatioship
	CheckAccess(ctx,client,storeId)
}