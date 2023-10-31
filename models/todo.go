package models

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type Todo struct {
	Description string
	Done        bool
}

// NewTodo 함수는 Todo 구조체의 생성자입니다.
// Done은 기본적으로 false로 처리 됩니다.
func NewTodo(description string) Todo {
	return Todo{description, false}
}

// String() 메서드는 Stringer 인터페이스를 구현한 메서드 입니다.
// 이 메서드는 Todo구조체 값의 Description - Done형태로 값을 반환합니다.
func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}

func NewTodoFromDataItem(di binding.DataItem) Todo {
	v, _ := di.(binding.Untyped).Get()
	return v.(Todo)
}
