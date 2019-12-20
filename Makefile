ProjectID= web-todo-list
Path=${PWD}

# run API Server
app-run:
	go run cmd/api/main.go

# Google App Engin

# run data store emulators
gae-datastore:
	gcloud beta emulators datastore  start --data-dir=${Path}

# Set ProjectID
gae-set-ProjectID:
	gcloud config set project ${ProjectID}

# deploy to gcp gae
gae-deploy:
	gcloud app deploy