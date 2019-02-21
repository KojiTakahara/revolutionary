# revolutionary

## setup

- `$ go get github.com/labstack/echo`
- `$ go get github.com/dgrijalva/jwt-go`
- `$ cd static`
- `$ npm install`

## local development environment

### start
- `$ dev_appserver.py app.yaml`
  - `dev_appserver.py --clear_datastore=yes app.yaml`
- `$ cd static`
- `$ npm run build -- --watch`
-> http://localhost:8080/

### master data

- http://localhost:8080/api/v1/race

### tournament data
- http://localhost:8080/cron/v1/tournamentHistory/{{tournamentId}}
  - tournamentId = http://dmvault.ath.cx/duel/tournament_history.php?tournamentId=XXX

## deploy

- `$ gcloud app deploy --project {PROJECT_ID} --version {VERSION}`

#### Uploading cron jobs
- `$ gcloud app deploy cron.yaml`

#### Uploading datastore　index
- `$ gcloud app deploy index.yaml`
