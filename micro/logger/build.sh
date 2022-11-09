GEN_DIR=pb
INC_DIR=$(find ~/go/ -name empty.proto | xargs -l dirname)
protoc -I=${INC_DIR} --go_opt=paths=source_relative --go_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. ./${GEN_DIR}/logger.proto 