#!/bin/bash

docker exec -it gogollellero_web goose up
docker exec -it gogollellero_web goose -env test up
