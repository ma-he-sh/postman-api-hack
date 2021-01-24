### GITIGNORE CLI

### Server
- cli api server
- run manually by
```sh
go run server
```
- or use docker-compose
```sh
docker-compose up -d
```
- goto `localhost:8080/api/{request}`


### Parser 
- run a cron job to sync `https://github.com/github/gitignore` and parse the data to json file
#### Setup
```sh
python3 -m venv venv
source venv/bin/activate
python parser.py
```
