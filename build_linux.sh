# /bin/sh
cd $(dirname $0)
cd ./frontend
npm install
npm run build
cd ../
GOOS=linux GOARCH=amd64 go build start.go