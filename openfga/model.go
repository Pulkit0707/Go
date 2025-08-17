// define and write the auth model

package main

import (
	"context"
	"log"

	openfga "github.com/openfga/go-sdk"
)

func WriteModel(ctx context.Context, client *openfga.SdkClient, storeID string){
	model:=openfga.AuthorizationModel{
		SchemaVersion: "1.1",
		TypeDefinitions: []openfga.TypeDefinition{
			{
				Type: "document",
				Relations: &map[string]openfga.Userset{
					"reader": {This: &map[string]interface{}{}},
					"writer": {This: &map[string]interface{}{}},
				},
			},
			{
				Type: "user",
			},
		},
		_,_,err:=client.WriteAuthorizationModel(ctx,StoreID).
		Body(openfga.WriteAuthorizationModelRequest{
			SchemaVersion: model.SchemaVersion,
			TypeDefinitions: model.TypeDefinitions,
		}).
		Execute()
		if err!=nil{
			log.Fatal("Error writing model:", err)
		}
	}
}