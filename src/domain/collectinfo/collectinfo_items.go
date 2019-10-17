package collectinfo

/*
- path: "/example-golang-rest-api/"
	  GET: req: {"name": "name"}, res: {"name": "name", "description": "test", "datas": [{post-data-1}, {post-data-2}] }
	  POST: req: {"name": "name", "description": "test", "data": {"id": 1, "name": "test1" }}, res: {"name": "name", "message": "ok"}
*/

type CreateCollectInfoRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Data        CollectInfo `json:"data"`
}

type CreateCollectInfoResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type CollectInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetCollectInfoRequest struct {
	Name string `json:"name"`
}

type GetCollectInfoResponse struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Datas       []CollectInfo `json:"datas"`
}
