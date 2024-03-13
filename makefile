


# 拉取镜像
# docker pull bitnami/mysql:latest
# docker pull bitnami/redis:latest

# 登录控制台
# docker exec -it storage sh
# mysql -uroot -pmy-secret-pw
# docker exec -it cache sh
# redis-cli


mysql:
	docker run -itd \
		--name storage \
		-p 13306:3306 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		-e MYSQL_ROOT_PASSWORD=my-secret-pw \
		bitnami/mysql:latest
	
redis:
	docker run -itd \
		--name cache \
		-p 6379:6379 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		bitnami/redis:latest
	
test:
	go test -count=1 -v ./...



.PHONY: mysql redis test 
