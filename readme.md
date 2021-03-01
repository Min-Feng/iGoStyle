# tutorial

## 需求敘述

請實作兩隻API，分別是教師列表以及教師個人資訊API並透過快取回傳資料
- API 1 教師列表 路徑： `/api/tutors/{language.slug}`
- API 2 教師個人資訊 路徑： `/api/tutor/{tutor.slug}`

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
