

# 登录 MySQL 控制台
# docker exec -it mysql sh
# mysql -uroot -p
# my-secret-p

# 登录 Redis 控制台
# docker exec -it redis sh
# redis-cli


.PHONY: storage cache


storage:
	docker run \
	-d \
	-p 3306:3306 \
	--name mysql \
	-e MYSQL_ROOT_PASSWORD=my-secret-pw \
	-e MYSQL_DATABASE=xyz_vote \
	mysql:8.0.29
	

cache:
	docker run \
	-d \
	-p 6379:6379 \
	--name redis \
	redis:7-alpine