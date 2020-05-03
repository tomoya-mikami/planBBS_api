// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// [END import]

type Event struct {
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	URL         string `firestore:"url"`
	Date        int64  `firestore:"date"`
}

type Plan struct {
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	Events      *[]Event
}

func addPlanHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "planbbs")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	switch r.Method {
	case http.MethodPost:
		events := []Event{
			{Title: "空港", Description: "羽田空港にむかいます", URL: "https://tokyo-haneda.com/index.html", Date: 1588491615665},
			{Title: "ねこ", Description: "ねこです", URL: "", Date: 1588491620000},
		}
		plan := Plan{Title: "旅行プラン", Description: "赤坂の旅行プランです", Events: &events}

		_, _, err = client.Collection("Plans").Add(ctx, plan)
		if err != nil {
			log.Printf("データ書き込みエラー　Error:%T message: %v", err, err)
			return
		}
	case http.MethodGet:
		plans := make([]Plan, 0)

		iter := client.Collection("Plans").Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			var plan Plan
			doc.DataTo(&plan)
			plans = append(plans, plan)
		}

		res, _ := json.Marshal(plans)

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	default:
		http.NotFound(w, r)
		return
	}
}

// [START main_func]

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/plan", addPlanHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// [END indexHandler]
// [END gae_go111_app]
