#! /bin/bash
# build-image.sh

cd accountservice
dep ensure -v
export GOOS=linux
export CGO_ENABLED=0
# go build -a -installsuffix nocgo -o /app .
go build -o ../bin/account
echo build `pwd`
cd ..

echo "Start build docker image ...."
docker build -t ping/account .