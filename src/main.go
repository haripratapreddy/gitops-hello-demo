/**
 * Copyright 2021 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START gke_hello_app]
// [START container_hello_app]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	environment := os.Getenv("ENVIRONMENT")
	build_id := os.Getenv("BUILD_ID")
	commit_id := os.Getenv("GIT_COMMIT_ID")
	fmt.Fprintf(w, "Hello from BrainUpgrade!\n")
	fmt.Fprintf(w, "Feature: Initial\n")
	fmt.Fprintf(w, "BUILD_ID: %s\n", build_id)
	fmt.Fprintf(w, "GIT_COMMIT_ID: %s\n", commit_id)
	fmt.Fprintf(w, "Hostname: %s\n", host)
	fmt.Fprintf(w, "Environment: %s\n", environment)

}
