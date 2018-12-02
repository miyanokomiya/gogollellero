#!/bin/bash

echo "docker-compose stop"
docker-compose stop
echo "rm -rf vendor"
rm -rf vendor
echo "docker-compose up"
docker-compose up
