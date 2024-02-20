# printmessage

A simple sample of a `go` project following [this](https://medium.com/gitconnected/learning-go-part-one-basic-concepts-64b50f11e19a) article.

## Go version

```
go version
```

## Initial Setup

```
mkdir printmessage && cd printmessage
go mod init printmessage
```

## Build

```
go build
```

## Run

```
go run main.go
```

## Unit Test

```
go test
```

## Managing imports


### Instaling tool
```
go install golang.org/x/tools/cmd/goimports@latest
```

### Inserting imports

```
goimports -w main.go
```