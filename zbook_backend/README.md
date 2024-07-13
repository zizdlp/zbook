# go

## 1. go module init

```shell
go mod init github.com/zizdlp/zbook #init
go mod tidy  #install deps
```

## database

> database 用于存储用户表、markdown表等，本repo选择postgres，可以使用sqlc等工具

### 安装依赖

> database 使用sqlc、golang migrate等

## image

## proxy

连不上github时使用代理

### wsl1

```shell
export https_proxy="http://127.0.0.1:1080"
export http_proxy="http://127.0.0.1:1080"
```

### wsl2

```shell
export hostip=$(cat /etc/resolv.conf | grep nameserver | awk '{ print $2 }')
alias www="https_proxy=\"http://${hostip}:1080\" http_proxy=\"http://${hostip}:1080\""
# or set it global
export hostip=$(cat /etc/resolv.conf | grep nameserver | awk '{ print $2 }')
export https_proxy="http://${hostip}:1080" 
export http_proxy="http://${hostip}:1080"
```

## testgo

## go module init

1. go mod init github.com/zdlpsina/testgo
2. go mod tidy

## go run

  go run main.go

## go build

  go build main.go

## go test

  go test -v -cover ./...

## go export cgo

go build -o golib.so -buildmode=c-shared main.go

## bazel& gazelle

### run

```shell
bazel run //:gazelle
```

### 添加依赖

```shell
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
bazel run //:gazelle
```

### test

```shell
bazel test //srcgo:srcgo_test
```

## POST MAN

### CREATE_USE

```shell
post: 0.0.0.0:8080/users 
json:-->body-->json
{
    "username":"guestwhat",
    "password":"guestwhat",
    "email":"guestwhat@zizdlp.com"
}
```

### LOGIN

```shell
post: 0.0.0.0:8080/users/login
json: -->body-->json
{
    "username":"guestwhat",
    "password":"guestwhat",
}
```

### CREATE_REPO

auth:-->bear_token

```shell
post: 0.0.0.0:8080/repos
json: -->body-->json
{
 "repo_name": "test"
}
```

docker build -t b_image -f baseDockerfile .
docker build -t alpine_image -f alpineDockerfile .

//
sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
  &&

&& bazel run //:gazelle \
&& bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies \
&& bazel run //:gazelle \
bazel run //:gazelle
bazel test ... && bazel build ...

## gazelle

bazel run //image:gazelle
bazel run //image:gazelle -- update-repos -from_file=image/go.mod -to_macro=deps.bzl%go_dependencies
bazel run //image:gazelle

## 使用github action secret

```shell
on: push
jobs:
  deployment:
    runs-on: ubuntu-latest
    steps:
    - name: Deploy Stage
      uses: fjogeleit/http-request-action@v1
      with:
        url: 'https://api.zizdlp.com/sync_repo'
        method: 'POST'
        customHeaders: '{"Content-Type": "application/json"}'
        data:  '{"repo_id":${{ secrets.API_ID }}, "username":  "${{ secrets.API_USER }}" ,"key": "${{ secrets.API_KEY }}"}'
```

## mockgen

export PATH=$PATH:~/go/bin

## grpc

brew install protobuf
protoc --version

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

### setting json in vscode

  "protoc": {
      "options": [
          "--proto_path=proto",
      ]
  }

brew tap ktr0731/evans
brew install evans

exit

evans  --host localhost --port 9090 -r repl

### gateway

go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

## goldmark

reader:
  source:所有字符
  segment.Stop, segment.Start:[Start,Stop)
  line, seg := reader.PeekLine()
 fmt.Println("===continue  second line is:", string(line[:len(line)-1]), seg.Start, seg.Stop)
  [33  10]    63
      start,stop

go get github.com/rakyll/statik
go intall github.com/rakyll/statik
