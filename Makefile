OPT_SEED ?=
ifneq ($(seed),)
OPT_SEED := -seed :$(seed)
endif

port ?= 8080

run:
	docker-compose up -d 

logs:
	docker-compose logs --tail=100 -f

reset:
	docker-compose up -d --force-recreate

stop:
	docker-compose stop

restart:
	docker-compose restart dex

clean:
	- docker-compose down

exec: 
	docker-compose exec dex bash

ps:
	docker-compose ps

node:
	go run cmd/node/node.go -c genesis/nodes/node-$(node) -genesis genesis/genesis.gob -port $(port) -rpc-addr :$(rpc-addr) $(OPT_SEED)

wallet:
	go run cmd/wallet/wallet.go -c credentials/node-$(node) $(cmd)

gen_credentials:
	go run cmd/gen_credentials/gen_credentials.go -N $(N)

gen_genesis:
	go run cmd/gen_genesis/gen_genesis.go -N $(N) -t $(t) -g $(g) -tokens tokens.txt -distribute-to credentials -dir genesis

credential_info:
	go run cmd/credential_info/credential_info.go -c credentials/node-$(node)

watch-api:
	PORT=$(port) ./bin/air -c api/air.conf	

.PHONY: %
.DEFAULT: exec