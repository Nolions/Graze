# run & build
app-run:
	go run cmd/main.go

app-build:
	go build cmd/main.go

# Google App Engin
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

