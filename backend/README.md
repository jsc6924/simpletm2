## development

### with live reload

```
air
```

### normal

```
go build -o main .
./main.exe
```

### docker compose

notice that live reload does not work with docker compose on windows

```
docker compose up --build
```

## build image

```
cd backend
docker build -f prod.Dockerfile -t jsc723/simpletm2-backend:0.0.15 .
docker push jsc723/simpletm2-backend:0.0.15
update swarm.yml
../push.sh
#on remote
make up
```

## db migration

```
sqlite3 data/SimpleTM.db < migrate/1.sql
```

## generate orm code using sqlboiler

```
sqlboiler sqlite3
```
