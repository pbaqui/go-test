package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

//Output is a struct
type Output struct {
	OwnerID string `json:"ownerId"`
	Source  string `json:"source"`
	EventID string `json:"eventId"`
}

// ProjectID is the env var of project id
var ProjectID = os.Getenv("PROJECTID")

// PubSubTopic is the topic name
var PubSubTopic = os.Getenv("PSOUTPUT")

// Sender generates a static json and sends it to Cloud Pub/Sub
func Sender(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, ProjectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var output Output
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter the json object")
	}

	json.Unmarshal(reqBody, &output)

	topic := client.Topic(PubSubTopic)

	// let's pub it
	outputJSON, _ := json.Marshal(output)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: outputJSON,
	})

	id, err := result.Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Pubbed record as message id %v: %v", id, string(outputJSON))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
