#!/bin/bash

docker exec -it gogollellero_web go-assets-builder -v Configs -p assets configs/ > app/server/assets/configs.go
