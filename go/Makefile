# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: minh-ngu <minh-ngu@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2022/09/28 10:56:09 by minh-ngu          #+#    #+#              #
#    Updated: 2024/09/16 22:59:19 by ngoc             ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

all:
	air -c air.toml
build:
	cd build-tool && go build -o ../build .
run:
	cd run-tool && go build -o ../run .
.PHONY: all clean fclean re build run
