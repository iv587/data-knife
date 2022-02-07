# /bin/sh
cd $(dirname $0)
cd ./frontend
yarn install
yarn run build
cd ../
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build gredisw.go