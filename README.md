# store_queryservice

```zsh
go get -u github.com/akira-saneyoshi/store_pb@v1.0.0
go get -u github.com/go-sql-driver/mysql
go get -u github.com/onsi/ginkgo/v2
go get -u github.com/onsi/gomega
go get -u go.uber.org/fx
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

```zsh
grpcurl -plaintext localhost:8083 list
grpcurl -plaintext localhost:8083 list proto.CategoryQuery
grpcurl -plaintext localhost:8083 list proto.ProductQuery

grpcurl -plaintext localhost:8083 proto.CategoryQuery.List
grpcurl -plaintext -d '{"id" : "762bd1ea-9700-4bab-a28d-6cbebf20ddc2" }' localhost:8083 proto.CategoryQuery.ById
grpcurl -plaintext -d '{"keyword" : "ペン" }' localhost:8083 proto.ProductQuery.ByKeyword
grpcurl -plaintext localhost:8083 proto.ProductQuery.List

grpcurl -plaintext -d '{"keyword" : "メン" }' localhost:8083 proto.ProductQuery.ByKeyword

grpcurl -cacert presen/prepare/queryservice.pem queryservice:8083 list 
grpcurl -cacert ./queryservice.pem -d '{"keyword" : "ペン" }' queryservice:8083 proto.ProductQuery.ByKeyword
```
