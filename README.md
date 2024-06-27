## Generating documentation

```
go generate ./...
```

## Linting

```
golangci-lint run
```

## Running tests

```
go test -v terraform-provider-string-functions/internal/provider
```

## Running acceptance tests
```
cd internal/provider
$env:TF_ACC=1; go test -count=1 -run='TestAccChunkStrings' 
cd ../..
```

## Building for macOS (Apple Silicon)

```
$env:GOOS = "darwin"; $env:GOARCH = "arm64"; go build -o terraform-provider-string-functions_darwin_arm64
```

## Building for Linux (x86_64)

```bash
GOOS=linux GOARCH=amd64 go build -o terraform-provider-string-functions_linux_amd64
```