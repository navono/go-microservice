#! /bin/bash
# build-image.sh

cd ./accountservice
go build -o ../bin/account

cd ../
docker build -t ping/account .