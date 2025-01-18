package main

import (
	"fyne.io/fyne/v2/dialog" // Ubah ke versi v2

	"dapodik_sync/services"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Struktur data login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var dataDapodik string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Login")

	// Input field
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username or Email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	// Button Login
	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		if username == "" || password == "" {
			dialog.ShowInformation("Error", "Username dan Password tidak boleh kosong", myWindow)
			return
		}
		// Proses Login
		if services.PerformLogin(username, password) {
			dialog.ShowInformation("Success", "Login Berhasil", myWindow)
			services.OpenMainWindow(myApp)
			myWindow.Close()
		} else {
			dialog.ShowInformation("Error", "Login Gagal", myWindow)
		}
	})

	// Layout Login
	loginForm := container.NewVBox(
		widget.NewLabel("Sign In"),
		usernameEntry,
		passwordEntry,
		layout.NewSpacer(),
		loginButton,
	)
	myWindow.SetContent(loginForm)
	myWindow.ShowAndRun()
}
