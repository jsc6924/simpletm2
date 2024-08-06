#!/bin/bash
cd "$(dirname "$0")"

scp -r swarm.yml Makefile 'jsc@jscrosoft.com:/home/jsc/simpletm2/'
# scp -r backend/migrate 'jsc@jscrosoft.com:/home/jsc/simpletm2/backend/'