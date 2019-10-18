package collectinfo

/*
- path: "/api/ex-golang/rest-api/"
	  GET: req: {"name": "name"}, res: {"name": "name", "description": "test", "datas": [{post-data-1}, {post-data-2}] }
	  POST: req: {"name": "name", "description": "test", "data": {"id": 1, "name": "test1" }}, res: {"name": "name", "message": "ok"}
*/

//CreateCollectInfoRequest request of post method
type CreateCollectInfoRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Data        CollectInfo `json:"data"`
}

//CreateCollectInfoResponse response of post method
type CreateCollectInfoResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

//CollectInfo request data
type CollectInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//GetCollectInfoRequest request of get method
type GetCollectInfoRequest struct {
	Name string `json:"name"`
}

//GetCollectInfoResponse response of get method
type GetCollectInfoResponse struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Datas       []CollectInfo `json:"datas"`
}
