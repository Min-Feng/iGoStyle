# tutorial

## go run

```bash
CFG_NAME=dev CGO_ENABLED=0 go run ./cmd/server/main.go
```

## go test

integration test

```bash
CFG_NAME=dev go test -tags=integration ./... -count=1
```

執行特定 unit test 的 subtest, 需要在該 test 檔案的路徑底下執行命令才有效果

```bash
CFG_NAME=dev go test -test.run TestLessonRepo -testify.m ^Test_sqlByTutorIDGroup$
```

## config file

[config/dev.yaml](./config/dev.yaml)
