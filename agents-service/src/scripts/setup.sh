# shellcheck disable=SC1113
#/!bin/bash

HOME="/Users/anhlai"

PATH_TO_FOLDER=$HOME"/learnenough/go/go-kit/go-microservices-go-kit/"

sudo docker-compose down

./clean-up.sh

cd $PATH_TO_FOLDER"agents-service/src/api/"
go build -a -o ../../main main.go

cd $PATH_TO_FOLDER"agents-service/src/scripts"
sudo docker-compose up -d --build