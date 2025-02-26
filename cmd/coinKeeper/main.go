package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// Создание нового окна
		w := new(app.Window)
		w.Option(app.Title("Hello Gio"), app.Size(unit.Dp(400), unit.Dp(300)))

		// Запуск основного цикла
		if err := loop(w); err != nil {
			panic(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	// Создание темы Material Design
	th := material.NewTheme()

	// Основной цикл обработки событий
	for {
		e := w.NextEvent()
		switch e := e.(type) {
		case system.FrameEvent:
			// Создание контекста отрисовки
			gtx := layout.NewContext(new(op.Ops), e)
			// Отрисовка текста "Hello, Gio!"
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H1(th, "Hello, Gio!").Layout(gtx)
			})
			// Обновление окна
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			// Выход из цикла при закрытии окна
			return e.Err
		}
	}
}

/*

func main() {
	go func() {
		w := app.NewWindow(app.Title("Купи Скайрим"), app.Size(800, 600))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.UpdateEvent:
			w.Invalidate()
		case app.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			// Рисуем фон
			paint.Fill(gtx.Ops, color.NRGBA{R: 70, G: 206, B: 173, A: 255})

			// Пример текста
			label := material.H1(th, "- Hello there!")
			label.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}


import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	appTest := app.New()
	wind := appTest.NewWindow("Lets Go BROOO")

	label := widget.NewLabel("Lets Go BROOO\n И вот так могет!")

	wind.SetContent(container.NewVBox(
		label,
	))

	wind.ShowAndRun()

}
*/
