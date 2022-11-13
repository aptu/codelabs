# Overview

# Building image

```
docker build -t merger .
```

## Testing local changes

```
sudo docker build --build-arg local=1 -t test .
```
Note: testing binary file must exist during image build.

docker-compose build --no-cache

# Testing

```
go test -run TestServer
```

```
grpcurl -plaintext -d @ -import-path pb -proto merger.proto \
    localhost:7777 pb.MergeLists/Merge << EOM
{
    "list1": {"v": [1]},
    "list2": {"v": [2]}
}
EOM
```


