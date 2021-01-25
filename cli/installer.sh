!#/bin/bash
## DEVCLI installer

go build devcli.go

sudo cp devcli /usr/local/bin
cp .devcli ~/.devcli
