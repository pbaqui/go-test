# go-test

## Create topic
`gcloud pubsub topics create jsonData`

## Api
POST Request that publishes a PubSubMessage
`gcloud functions deploy api --entry-point Sender --runtime go111 --trigger-http --env-vars-file .env.dev.yaml`

## Consumer
Receive a PubSubMessage and store to the table `data`
`gcloud functions deploy consumer --entry-point Receive --runtime go111 --trigger-topic=jsonData`

## Logs
`gcloud functions logs read consumer`

## Clean
* Delete topic: `gcloud pubsub topics delete jsonData`
* Delete function api: `gcloud functions delete api`
* Delete function consumer:  `gcloud functions delete consumer`