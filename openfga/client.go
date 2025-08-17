// Connect to openfga

package main

import (
	openfga "github.com/openfga/go-sdk"
)

func NewFGAClient() (*openfga.SdkClient, error) {
	return openfga.NewSdkClient(&openfga.ClientConfiguration{
		ApiScheme: "http",
		ApiHost:   "localhost:8080", // change if running on another host
	})
}
