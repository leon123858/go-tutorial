package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Todo struct {
	Text string
	Done bool
}

type TodoList struct {
	Todos []Todo
}

func main() {
	a := app.New()
	w := a.NewWindow("Todo List")

	// set window size
	w.Resize(fyne.NewSize(400, 600))

	todoList := &TodoList{}
	input := widget.NewEntry()
	input.SetPlaceHolder("Add a todo")

	list := widget.NewList(
		func() int {
			return len(todoList.Todos)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(todoList.Todos[id].Text)
		},
	)

	addButton := widget.NewButton("Add", func() {
		todoList.Todos = append(todoList.Todos, Todo{Text: input.Text})
		input.SetText("")
		refreshList(list)
	})
	deleteAllButton := widget.NewButton("Delete All", func() {
		todoList.Todos = nil
		refreshList(list)
	})

	refreshList(list)

	title := widget.NewLabel("My Todo List")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// 在標題下面添加兩個空白行
	space1 := widget.NewLabel("")
	space2 := widget.NewLabel("")

	// now clock
	now := time.Now()
	nowLabel := widget.NewLabel(now.Format("2006-01-02 15:04:05"))
	go func() {
		for {
			time.Sleep(time.Second)
			now = time.Now()
			nowLabel.SetText(now.Format("2006-01-02 15:04:05"))
		}
	}()

	// counter
	//counter := 0
	counterWidget := widget.NewLabel("0")
	//mtx := &sync.Mutex{}
	//go func() {
	//	for i := 0; i < 100000; i++ {
	//		mtx.Lock()
	//		counter++
	//		mtx.Unlock()
	//	}
	//	strCounter := strconv.Itoa(counter)
	//	counterWidget.SetText(strCounter)
	//}()
	//go func() {
	//	for i := 0; i < 100000; i++ {
	//		mtx.Lock()
	//		counter++
	//		mtx.Unlock()
	//	}
	//	strCounter := strconv.Itoa(counter)
	//	counterWidget.SetText(strCounter)
	//}()

	content := container.NewBorder(
		container.NewVBox(title, space1, space2, input, addButton, deleteAllButton),
		nowLabel, counterWidget, nil,
		list,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

func refreshList(list *widget.List) {
	list.Refresh()
}
