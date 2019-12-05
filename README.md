# ToDo List API Service

## setting environment variable
cope .env.example to .env
    cp .env.example .env
    
### Run on local

Run
    make app-run

Build 
    make app-build
    
## Run on Google APP Engine

Set Project
    make gae-set-ProjectID ProjectID=<ProjectID>
    
Deploy
    make gae-deploy
    
Browser
    make gae-browse