stack = simpletm
envfile = ./.env

include $(envfile)
export

up:
	docker stack deploy -c swarm.yml --resolve-image always $(stack)

down:
	docker stack rm $(stack)

config:
	docker compose -f swarm.yml config