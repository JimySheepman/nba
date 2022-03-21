#!/bin/bash
if [ ! "$(docker ps -q -f name=postgres)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=postgres)" ]; then
        # cleanup
        echo "clean image"
        docker rm postgres
    fi
    # run your container
    docker run --name postgres -e POSTGRES_PASSWORD=root -d -p 5432:5432 postgres
fi

if [ ! "$(docker ps -q -f name=nba)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=nba)" ]; then
        # cleanup
        echo "clean image"
        docker rm nba
    fi
    # buid your container
    docker build -t nba .
    # run your container
    docker run -dp 3000:3000 nba
fi
