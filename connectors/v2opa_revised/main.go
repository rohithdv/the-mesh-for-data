// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"
	"os"

	v2opa "github.com/mesh-for-data/mesh-for-data/connectors/v2opa_revised/openapiserver"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	log.Printf("Env. variable extracted: %s - %s\n", key, value)
	return value
}

func main() {
	// port := getEnvWithDefault("PORT_OPA_CONNECTOR", defaultPort)
	// opaServerURL = getEnv("OPA_SERVER_URL") // set global variable

	// log.Println("OPA_SERVER_URL env variable in OPAConnector: ", opaServerURL)
	// log.Println("Using port to start go opa connector : ", port)

	// log.Printf("Server starts listening on port %v", port)
	// lis, err := net.Listen("tcp", ":"+port)
	// if err != nil {
	// 	log.Fatalf("Error in listening: %v", err)
	// }
	// s := grpc.NewServer()
	// srv := &server{opaReader: opabl.NewOpaReader(opaServerURL)}
	// pb.RegisterPolicyManagerServiceServer(s, srv)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Error in service: %v", err)
	// }

	log.Printf("Server started")

	DefaultApiService := v2opa.NewDefaultApiService()
	DefaultApiController := v2opa.NewDefaultApiController(DefaultApiService)

	router := v2opa.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":50050", router))

}