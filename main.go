package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"kubetui/driver"
	"kubetui/gateway"
	"kubetui/usecase"
	"os/exec"
)

func main() {

	podUsecese := usecase.PodUsecase{
		PodPort: gateway.PodGateway{
			PodDriver: driver.K8sPodDriver{},
		},
	}

	app := tview.NewApplication()
	list := tview.NewList()
	pages := tview.NewPages()

	pages.AddAndSwitchToPage("pod list", list, true)

	selectList := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r'}

	for i, out := range podUsecese.GetPodList("") {
		list.AddItem("List item "+string(i), out, selectList[i], func() {
			des, _ := exec.Command("kubectl", "-n", "", "describe", "pod", out).Output()
			pages.AddAndSwitchToPage("described pod", createTextView(string(des), pages), true)
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(pages, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func createTextView(text string, pages *tview.Pages) *tview.TextView {

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(true)

	textView.SetText(text)

	textView.SetDoneFunc(func(key tcell.Key) {
		if tcell.KeyEnter == key {
			pages.RemovePage("described pod")
		}
	})
	textView.InputHandler()

	return textView
}
