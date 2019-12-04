# run & build
app-run:
	go run cmd/main.go

app-build:
	go build cmd/main.go

# Google App Engin
gae-deploy:
	gcloud app deploy

gae-browse:
	gcloud app browse
