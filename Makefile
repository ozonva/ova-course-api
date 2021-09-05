DBSTRING:="postgres://$(DB_USERNAME):$(DB_PASSWORD)@localhost:5430/$(DB_NAME)?sslmode=disable"

.PHONY: build
build: .install-go-deps .vendor-proto .proto-generate migration-up .build

.PHONY: .install-go-deps
.install-go-deps:
	go get -u github.com/rs/zerolog/log
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/protobuf/reflect/protoreflect@v1.27.1
	go get -u google.golang.org/protobuf/runtime/protoimpl
	go get -u google.golang.org/protobuf/types/known/emptypb
	go get -u google.golang.org/protobuf/types/known/timestamppb
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u github.com/jmoiron/sqlx
	go get -u github.com/joho/godotenv
	go get -u github.com/pressly/goose/v3/cmd/goose
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/Masterminds/squirrel
	go get -u github.com/opentracing/opentracing-go
	go get -u github.com/segmentio/kafka-go



.PHONY: .vendor-proto
.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/ova-course-api
	cp api/ova-course-api/ova-course-api.proto vendor.protogen/api/ova-course-api/ova-course-api.proto
	@if [ ! -d vendor.protogen/google ]; then \
	  git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
	  mkdir -p  vendor.protogen/google/ &&\
	  mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
	  rm -rf vendor.protogen/googleapis ;\
	fi

.PHONY: proto-generate
 proto-generate:
	protoc -I vendor.protogen \
			--go_out=pkg/ova-course-api --go_opt=paths=import \
			--go-grpc_out=pkg/ova-course-api --go-grpc_opt=paths=import \
			api/ova-course-api/ova-course-api.proto
	mv pkg/ova-course-api/github.com/ozonva/ova-course-api/pkg/ova-course-api/* pkg/ova-course-api/
		rm -rf pkg/ova-course-api/github.com
		mkdir -p cmd/ova-course-api


.PHONY: .build
.build:
	go build -o ./bin/main ./cmd/ova-course-api

 .PHONY: migration-up
 migration-up:
	goose -dir=migrations postgres "user=$(DB_USER) dbname=$(DB_NAME) password=$(DB_PASSWORD) sslmode=disable" up




