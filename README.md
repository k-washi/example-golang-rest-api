# example-golang-rest-api

---


```yaml
- path: "/example-golang-rest-api/"
      GET: req: {"name": "name"}, res: {"name": "name", "description": "test", "datas": [{post-data-1}, {post-data-2}] }
	POST: req: {"name": "name", "description": "test", "data": {"id": 1, "name": "test1" }}, res: {"name": "name", "message": "create info and store database"}
- path: /health/
      GET: status:200 response and , {"health": 200}
- info:
      git: https://github.com/k-washi/example-golang-rest-api.git
      image:
      description: golangにより構築したREST API
```

## make

```bash

make help
#bin/%:             build binaries ex. make bin/myproj
#build:             build binary
#deps:              Install dependencies
#devel-deps:        Setup
#lint:              Lint
#test:              Run tests

```