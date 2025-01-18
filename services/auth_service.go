package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Fungsi untuk mengirim request login
func PerformLogin(username, password string) bool {
	url := "http://localhost:8081/api/v1/login"
	loginData := LoginRequest{
		Username: username,
		Password: password,
	}

	// Konversi ke JSON
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return false
	}

	// Kirim request POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	// Periksa response
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// Window Utama setelah Login Berhasil
func OpenMainWindow(a fyne.App) {
	mainWindow := a.NewWindow("Main Window")
	mainWindow.Resize(fyne.NewSize(400, 400))

	// Tombol Get Data
	getDataButton := widget.NewButton("Get Data", func() {
		data := getDataFromDapodik()
		dialog.ShowInformation("Data Retrieved", fmt.Sprintf("Data: %s", data), mainWindow)
	})

	// Tombol Sync
	syncButton := widget.NewButton("Sync", func() {
		success := syncDataToServer()
		if success {
			dialog.ShowInformation("Success", "Data berhasil dikirim!", mainWindow)
		} else {
			dialog.ShowInformation("Error", "Gagal mengirim data", mainWindow)
		}
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Welcome!"),
		getDataButton,
		syncButton,
	)
	mainWindow.SetContent(content)
	mainWindow.Show()
}

// Fungsi untuk Get Data dari aplikasi Dapodik
func getDataFromDapodik() string {
	// Simulasi pengambilan data
	return "Sample Data dari Aplikasi Dapodik"
}

// Fungsi untuk Sync Data ke server
func syncDataToServer() bool {
	url := "http://localhost:8082/api/v1/data-dapodik"
	payload := map[string]string{"data": "Sample Data dari Dapodik"}

	// Konversi ke JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return false
	}

	// Kirim request POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	resp.StatusCode = http.StatusOK
	// if resp.StatusCode == http.StatusOK {
	// 	return true
	// }
	return false
}
