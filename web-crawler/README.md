# Web Crawlers

## Description

利用 go 編寫爬蟲, 抓取 ptt 網站的文章標題

作為 golang 高併發設計的範例

## install

```shell
go get -u github.com/PuerkitoBio/goquery
go mod tidy
```

## 目錄

- 用物件導向的方式撰寫爬蟲 (OOP)
- 用函數導向的方式撰寫爬蟲 (FP)

## 使用方式

可以到主程式修改關鍵字

```shell
go run cmd/oop.go
go run cmd/fp.go
```