## update and deploy

```
cd frontend
npm run build
docker build -f prod.Dockerfile -t jsc723/simpletm2-frontend:0.0.7 .
docker push jsc723/simpletm2-frontend:0.0.7
update swarm.yml
git add -u
git commit -m update
git push
git pull
make up
```
