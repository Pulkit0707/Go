// connect to OpenFGA

package main

import(
	openfga "github.com/openfga/go-sdk"
)

func NewFGAClient()(*openfga.SdkClient, error){
	return openfga.NewSdkClient(&openfga.ClientConfigurations{
		ApiScheme: "http",
		ApiHost: "localhost:8080",
	})
}