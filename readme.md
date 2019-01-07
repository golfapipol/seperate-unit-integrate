# Seperate Unit Test and Integration Test

ใส่ // +build ลงไปใน test ที่เป็น integration

```golang
// +build integration

package fizzbuzz
```

## Run Unit Test

```bash
go test ./...
```

## Run Integration Test

```bash
go test -tags=integration ./...
```