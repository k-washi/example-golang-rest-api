# example-golang-rest-api

---


```yaml
- path: "/example-golang-rest-api/"
      - GET: {name: "name", description: "test", data: [{}, {}] }
      - POST: {name: "name", description: "test", data: {id: 1, name: "test1" }}
      - /health/: status:200 response and , {"health": 200}
    git: https://github.com/k-washi/example-golang-rest-api.git
    image:
    description: golangにより構築したREST API
```