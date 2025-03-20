package mongodoc

import (
	"time"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Todo struct {
	_ID       bson.ObjectID `bson:"_id,omitempty"`
	ID        todo.ID       `bson:"id"`
	Project   project.ID    `bson:"project"`
	Name      string        `bson:"name"`
	Done      bool          `bson:"done"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func (a *Todo) Into() (*todo.Todo, error) {
	return todo.New().
		ID(a.ID).
		Project(a.Project).
		Name(a.Name).
		Done(a.Done).
		CreatedAt(a._ID.Timestamp()).
		UpdatedAt(a.UpdatedAt).
		Build()
}

func New(a *todo.Todo) (*Todo, error) {
	return &Todo{
		ID:      a.ID(),
		Project: a.Project(),
		Name:    a.Name(),
	}, nil
}

type List []Todo

func (l List) Into() (todo.List, error) {
	list := make(todo.List, 0, len(l))
	for _, t := range l {
		todo, err := t.Into()
		if err != nil {
			return nil, err
		}
		list = append(list, todo)
	}
	return list, nil
}

func NewList(todos todo.List) (List, error) {
	list := make(List, 0, len(todos))
	for _, a := range todos {
		doc, err := New(a)
		if err != nil {
			return nil, err
		}
		list = append(list, *doc)
	}
	return list, nil
}
