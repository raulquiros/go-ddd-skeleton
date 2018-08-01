package Service

type PostResquest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	State   int    `json:"status"`
}

func Post(request PostResquest) {

}
