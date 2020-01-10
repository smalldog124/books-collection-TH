version?=latest
build_crawler:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/crawler ./main.go

build_image :
	docker build -t smalldog124/book-crawler:${version} .

push_image :
	docker push smalldog124/book-crawler:${version}

deploy :
	scp ./bin/crawler root@47.88.155.215:/root/book-crawler
	sed 's/TAG/${version}/g' docker-compose.template.yml >docker-compose.yml
	scp ./docker-compose.yml root@47.88.155.215:/root/book-crawler
	ssh root@47.88.155.215 "cd /root/book-crawler; docker-compose pull; docker-compose up -d"
	ssh root@47.88.155.215 "cd /root/book-crawler; ./crawler"