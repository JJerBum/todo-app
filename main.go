package main

import (
	"fmt"

	"github.com/JJerBum/todo-app/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 새로운 앱 생성
	a := app.New()

	// TODO Application이 제목인 새로운 창을 만듭니다.
	w := a.NewWindow("TODO Application")

	// window 크기를 300 by 400으로 조정
	w.Resize(fyne.NewSize(300, 400))

	t := models.NewTodo("Show thie on the window")

	// window에 content와 label을 지정해 줍니다.
	w.SetContent(
		container.NewBorder(
			nil, // 컨테이너 상단

			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
				widget.NewButton("Add", func() { fmt.Println("Add was clicked!") }),
				// take the rest of the space
				widget.NewEntry(),
			),

			nil, // Right
			nil, // Left

			// 나머지는 나머지 공간을 모두 차지합니다.
			container.NewCenter(
				widget.NewLabel(t.String()),
			),
		),
	)

	// 보여줍니다.
	w.ShowAndRun()

}
