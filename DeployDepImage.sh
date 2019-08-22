#!/bin/sh
docker service rm dermaster_worker_peekaboo 2>/dev/null
docker stack deploy -c docker-stack.deploy.yml dermaster
