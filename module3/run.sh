#!/bin/bash

IMAGE_NAME="httpv2"
CONTAINER_NAME="http_container"

docker run --name $CONTAINER_NAME -p 8080:80 -d $IMAGE_NAME
PID=`docker inspect -f '{{.State.Pid}}' $CONTAINER_NAME`
nsenter -t $PID -n ip a
