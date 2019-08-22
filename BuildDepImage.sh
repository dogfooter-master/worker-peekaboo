#!/bin/sh
EXE=docker
unameOut="$(uname -s)"
if [ 0 -ne 0 ]; then
	EXE=${EXE}.exe
fi

${EXE} build --no-cache -t dermaster/worker-peekaboo:latest .
docker push dermaster/worker-peekaboo:latest
