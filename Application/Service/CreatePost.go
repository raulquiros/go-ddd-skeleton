package Service

import (
	"go-ddd-structure/Domain/Model"
	"go-ddd-structure/Domain/Types"
)

type PostResquest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	State   int    `json:"status"`
}

func CreatePost(request PostResquest, generateIdService Model.GenerateIdInterface, postRepository Model.PostRepository) (bool, error) {

	var post Types.Post
	post.ID, _ = generateIdService.Generate()
	post.Title = request.Title
	post.Content = request.Content
	post.State = request.State
	post.State = request.State

	created, err := postRepository.Create(post)

	return created, err
}
