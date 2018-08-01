package Types

import (
	"github.com/satori/go.uuid"
	"google.golang.org/genproto/googleapis/type/date"
)

type Post struct {
	ID        uuid.UUID `bson:"id"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	State     int       `bson:"status"`
	CreatedAt date.Date `bson:"created_at"`
	UpdatedAt date.Date `bson:"updated_at"`
}
