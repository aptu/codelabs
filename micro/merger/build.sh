GEN_DIR=pb
protoc --go_opt=paths=source_relative --go_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=. ./${GEN_DIR}/merger.proto 

protoc --go_opt=paths=source_relative --go_out=${GEN_DIR} --proto_path=../logger/pb  --go-grpc_opt=paths=source_relative --go-grpc_out=${GEN_DIR} ../logger/${GEN_DIR}/logger.proto 