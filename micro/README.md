docker build -t merger .

docker-compose build --no-cache

go test -run TestServer

grpcurl -plaintext -d @ -import-path pb -proto merger.proto \
    localhost:7777 pb.MergeLists/Merge << EOM
{
    "list1": {"v": [1]},
    "list2": {"v": [2]}
}
EOM
