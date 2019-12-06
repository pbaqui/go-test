package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

// ProjectID is the env var of project id
var ProjectID = "verdant-bond-260917"

// PubSubMessage is the payload of a pubsub event
type PubSubMessage struct {
	Data []byte `json:"data"`
}

//Input is a struct
type Input struct {
	OwnerID string `json:"ownerId" bigquery:"ownerid"`
	Source  string `json:"source" bigquery:"source"`
	EventID string `json:"eventId" bigquery:"eventid"`
}

// Receive func logs an event payload
func Receive(ctx context.Context, m PubSubMessage) error {
	var input Input
	if err := json.Unmarshal(m.Data, &input); err != nil {
		log.Fatalf("Unable to unmarshal message %v with error %v", string(m.Data), err)
	}

	fmt.Println(input)

	log.Printf("Data received: %s", string(m.Data))

	// Store data
	client, err := bigquery.NewClient(ctx, ProjectID)
	if err != nil {
		return nil
	}
	datasetID := "dataset"
	tableID := "data"
	u := client.Dataset(datasetID).Table(tableID).Inserter()

	if err := u.Put(ctx, input); err != nil {
		return err
	}

	return nil
}
