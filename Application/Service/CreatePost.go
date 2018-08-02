package Service

import (
	"github.com/satori/go.uuid"
	"go-ddd-structure/Domain/Types"
)

type PostResquest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	State   int    `json:"status"`
}

func CreatePost(request PostResquest) {

	var post Types.Post
	post.ID, _ = uuid.NewV4()
	post.Title = request.Title
	post.Content = request.Content
	post.State = request.State
	post.State = request.State

}
