#!/bin/bash

docker exec -it gogollellero_web goose -env test up
if [ $# -eq 0 ]; then
  docker exec -it gogollellero_web bash -c "GO_ENV=test go test github.com/miyanokomiya/gogollellero/app/server/..."
else
  docker exec -it gogollellero_web bash -c "GO_ENV=test go test github.com/miyanokomiya/gogollellero/app/server/$1/..."
fi
