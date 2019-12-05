ProjectID=
Path=${PWD}


# run & build
app-run:
	go run cmd/main.go

app-build:
	go build cmd/main.go

# Google App Engin
# emulators
# run data store emulators
gae-emulators-datastore:
	gcloud beta emulators datastore start --data-dir=${Path}

# Set ProjectID
ProjectID=
gae-set-ProjectID:
	gcloud config set project ${ProjectID}

# deploy
gae-deploy:
	gcloud app deploy

# browse
gae-browse:
	gcloud app browse


