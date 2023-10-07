package main

import "github.com/jchv/go-webview2"

func main() {

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Window:        nil,
		Debug:         true,
		DataPath:      "",
		AutoFocus:     false,
		WindowOptions: webview2.WindowOptions{Title: "Demo View", Width: 1000, Height: 800},
	})
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("https://google.com")
	defer w.Destroy()

	w.Run()
}
