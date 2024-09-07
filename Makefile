# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: minh-ngu <minh-ngu@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2022/09/28 10:56:09 by minh-ngu          #+#    #+#              #
#    Updated: 2024/09/07 10:03:03 by ngoc             ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

all:
	@docker compose -f ./docker-compose.yml up -d --build
	@docker compose -f ./docker-compose.yml logs -f

down:
	@docker compose -f ./docker-compose.yml down

up:
	@docker compose -f ./docker-compose.yml up -d --build
	@docker compose -f ./docker-compose.yml logs -f

stop:
	@docker stop $$(docker ps)

remove_con:
	@docker rm -f $$(docker ps -a -q)

remove_images:
	@docker image prune --all --force

re:
	-@make gitd M="provisoire"
	#@make clean
	@make down
	@make up

clean:
	-docker stop $$(docker ps -qa)
	-docker rm $$(docker ps -qa)
	-docker rmi -f $$(docker images -qa)
	-docker volume rm $$(docker volume ls -q)
	-docker network rm $$(docker network ls -q)
	gio trash -f next/.next
	gio trash -f database/*.err database/*.log database/*.pid
	sudo gio trash -f zap/build/.zig-cache

cleanNext:
	cd next && gio trash -f node_modules
	cd next && npm install

cleanAll:
	-docker system prune

gitd:
	git add -A -- :!*.o :!*.swp :!*.env :!*.crt :!*.key
	git commit -m "$(M)"
	git push

dock:
	docker build -t img -f "$(DF)" "$(F)"
	docker run --rm -ti img sh

CURRENT_DIR := $(shell pwd)
docksql:
	docker build -t img -f "./cmp-be/Dockerfile_mysql_alpine" "./cmp-be"
	docker run --rm -ti \
	--name mysql \
	-v $(CURRENT_DIR)/cmp-be/database:/var/lib/mysql/data \
    -p 3306:3306 \
	-e DB_DATABASE=cmp_db \
    -e DB_PORT=3306 \
    -e DB_USER=haipb \
    -e DB_PASSWORD=123456 \
    -e DB_ROOT_PASSWORD=123456  \
	img sh

dockphp:
	docker build -t img -f "./cmp-be/Dockerfile_phpfpm" "./cmp-be"
	docker run -d -p 9000:9000 img
	# docker run --rm -ti -p 9000:9000 img sh

dockadmin:
	docker build -t img -f "./cmp-be/Dockerfile_phpMyAdmin" "./cmp-be"
	docker run --rm -ti \
		--name nest \
		-p 8000:8000 \
		-v $(CURRENT_DIR)/cmp-be/phpMyAdmin:/var/www/localhost/htdocs/phpmyadmin \
		img sh

docknest:
	docker build -t img -f "./cmp-proxy-app/Dockerfile_nest" "./cmp-proxy-app"
	docker run --rm -ti \
		--name nest \
		-v $(CURRENT_DIR)/cmp-proxy-app:/app \
    	-p 3001:3000 \
		-e DB_DATABASE=cmp_db \
    	-e DB_PORT=8889 \
    	-e DB_USER=haipb \
    	-e DB_PASSWORD=123456 \
    	-e DB_ROOT_PASSWORD=123456  \
		img sh

dockzap:
	docker build -t img -f "./zap/Dockerfile" "./zap"
	docker run --rm -ti \
		--name zap \
    	-p 4000:4000 \
		img sh

dockbe:
	docker build -t img -f "./cmp-be/Dockerfile_phpfpm" "./cmp-be"
	docker run --rm -ti -p 9000:9000 img sh

dockrm:
	docker rmi img:latest 

inspect:
	docker network inspect comac_network

devbe:
	@docker compose -f ./cmp-be/docker-compose.yml up -d --build
	@docker compose -f ./cmp-be/docker-compose.yml logs -f

devfe:
	@docker compose -f ./cmp-fe/docker-compose.yml up -d --build
	@docker compose -f ./cmp-fe/docker-compose.yml logs -f

devproxy:
	@docker compose -f ./cmp-proxy-app/docker-compose.yml up -d --build
	@docker compose -f ./cmp-proxy-app/docker-compose.yml logs -f

zap:
	@docker compose build --no-cache zap
	@docker compose up -d zap

gin1:
	cd gin/src && go build -o gin main.go

gin:
	@docker compose build --no-cache gin
	@docker compose up -d gin

next:
	@docker compose build --no-cache next
	@docker compose up -d next

remix:
	@docker compose build --no-cache remix
	@docker compose up -d remix

mysql:
	@docker compose build --no-cache mysql
	@docker compose up -d mysql

# @echo "PATH = $$PATH"
zig:
	cd zap/build && zig build

test:
	cd zap/build/src && zig build-exe test.zig && ./test
.PHONY: all clean fclean re test zap zig gin mysql next remix
