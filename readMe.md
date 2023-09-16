yarn init -y

yarn add --dev @grpc/grpc-js @grpc/proto-loader typescript ts-node

yarn proto-loader-gen-types --grpcLib=@grpc/grpc-js --outDir=proto/ proto/*.proto

protoc --go_out=. --go-grpc_out=. proto/random.proto

for running go server
go get github.com/githubnemo/CompileDaemon
CompileDaemon -command="./gots"

for running typescript
npm run client
npm run server

for running go