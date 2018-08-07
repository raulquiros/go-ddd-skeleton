package Model

import "go-ddd-structure/Domain/Types"

type PostRepository interface {
	Create(post Types.Post) (bool, error)

	Update(conditions Types.Post) (bool, error)

	Delete(conditions Types.Post) (bool, error)
}
