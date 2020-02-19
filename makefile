version?=latest
build : build_crawler build_image build_ui

build_crawler:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/crawler ./main.go

build_backend :
	docker build -t smalldog124/book-crawler:${version} .

build_ui :
	docker build -t smalldog124/book-ui:${version} ./ui

push_image :
	docker push smalldog124/book-crawler:${version}
	docker push smalldog124/book-ui:${version}

deploy :
	# scp ./bin/crawler root@47.88.155.215:/root/book-crawler
	scp -i lek.pem ./initdb.sql root@47.88.155.215:/root/book-crawler
	sed 's/TAG/${version}/g' docker-compose.template.yml > docker-compose.yml
	scp -i lek.pem ./docker-compose.yml root@47.88.155.215:/root/book-crawler
	ssh -i lek.pem root@47.88.155.215 "cd /root/book-crawler; docker-compose pull; docker-compose up -d"
	newman run atdd/add_book_collection.json -e atdd/alibabacloud_environment.json -d atdd/data/add_book.json
	ssh -i lek.pem root@47.88.155.215 "cd /root/book-crawler; docker-compose down -v; docker-compose up -d"

crawler :
	ssh root@47.88.155.215 "cd /root/book-crawler; ./crawler"

test :
	go test ./cmd/... -cover
	golangci-lint run ./cmd/... ./internal/...
	docker-compose up -d books-db
	go test ./internal/... -cover
	make build_backend
	make build_ui
	sed 's/TAG/${version}/g' docker-compose.template.yml > docker-compose.yml
	docker-compose up -d
	newman run atdd/add_book_collection.json -e atdd/book_local_environment.json -d atdd/data/add_book.json 
	newman run atdd/add_book_wishlist.json -e atdd/book_local_environment.json -d atdd/data/add_wishlist.json 
	docker-compose down -v
	# docker rmi $(docker image ls --filter "dangling=true" -q)

integration_test:
	docker-compose up -d books-db
	go test ./internal/books
	docker-compose down -v

test_local:
	docker-compose up -d
	newman run atdd/add_book_collection.json -e atdd/book_local_environment.json -d atdd/data/add_book.json 
	newman run atdd/add_book_wishlist.json -e atdd/book_local_environment.json -d atdd/data/add_wishlist.json 
	docker-compose down -v