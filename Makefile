ProjectID= web-todo-list
Path=${PWD}
DstastoreHostPort=8432


# run & build
app-run:
	go run cmd/main.go

app-build:
	go build cmd/main.go

# Google App Engin
# emulators
# run data store emulators
gae-datastore:
	gcloud beta emulators datastore  start --data-dir=${Path}

# Set ProjectID
gae-set-ProjectID:
	gcloud config set project ${ProjectID}

# deploy
gae-deploy:
	gcloud app deploy

# browse
gae-browse:
	gcloud app browse


