# example-golang-rest-api

---


```yaml
- path: "/rest-api/:8080"
      GET: req: /*?name="something", res: {"name": "name", "description": "test", "datas": [{post-data-1}, {post-data-2}] }
      POST: req: {"name": "name", "description": "test", id": 1, "data": "test1"}, res: {"name": "name", "message": "create info and store database"}
- path: /health/
      GET: status:200 response and , res:{"health": 200}
- info:
      git: https://github.com/k-washi/example-golang-rest-api.git
      image: kwashizaki/example-golang-rest-api
      description: golangにより構築したREST API
```

## make

```bash
# For linux
make help
#bin/%:             build binaries ex. make bin/app
#build:             build binary
#deps:              Install dependencies
#devel-deps:        Setup
#lint:              Lint
#test:              Run tests

```

## Docker 

コンテナのbuildと実行

```bash
docker build -t kwashizaki/example-golang-rest-api:v1.0.0 .
docker run -it -p 8080:8080 --rm --name example-golang-rest-api kwashizaki/example-golang-rest-api:v1.0.0

#no build
docker build -t kwashizaki/example-vue-cli:v1.0.0 -f ./DockerfileNoBuild .
```

Docker Hubへpush

```bash
docker push kwashizaki/example-golang-rest-api:v1.0.0
```