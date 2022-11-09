GEN_DIR=pb
INC_DIR=$(find ~/go/ -name empty.proto | xargs -n1 dirname)
protoc -I=${INC_DIR} --go_opt=paths=source_relative --go_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. ./${GEN_DIR}/logger.proto 
