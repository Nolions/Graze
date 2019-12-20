## runServer

# run API Server
app-run:
	go run cmd/api/main.go

# run Cache Server
cache-run:
	go run cmd/cache/main.go

ProjectID= web-todo-list
Path=${PWD}

## Google App Engin

# Set ProjectID
gae-set-ProjectID:
	gcloud config set project ${ProjectID}

# run datastore emulators
gae-datastore:
	gcloud beta emulators datastore  start --data-dir=${Path}

# deploy to gcp gae
gae-deploy:
	gcloud app deploy