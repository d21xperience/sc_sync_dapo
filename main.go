package main

// func main() {
// 	// URL tujuan
// 	apiURL := "http://localhost:5774/roleperan"

// 	// Data form yang akan dikirimkan
// 	formData := url.Values{}
// 	formData.Add("username", "d21xperience@gmail.com")
// 	formData.Add("password", "Pasja123*")
// 	formData.Add("semester_id", "20222")

// 	// Membuat request baru dengan metode POST
// 	req, err := http.NewRequest("POST", apiURL, strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	// Menambahkan header Content-Type
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	// Membuat HTTP client dan mengirim request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Membaca respons
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	// Menampilkan respons
// 	fmt.Println("Response:", string(body))
// 	fmt.Printf("URL extracted: %s\n", url)

// 	// Lakukan request ke URL yang diekstrak
// 	response, err := http.Get("http://localhost:5774/login?data=eyJzZWtvbGFoX2lkIjoiOGE1YmQ5NTctNjZiYy00MDk2LTlmZjEtZmVlMDk2Yjg3YmEwIiwicGVyYW5faWQiOjUzLCJwZXJhbl9pZF9yZWFsIjo1M30=&params=eyJ1c2VybmFtZSI6ImQyMXhwZXJpZW5jZUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IlBhc2phMTIzKiIsInNlbWVzdGVyX2lkIjoiMjAyMjIiLCJyZW1lbWJlcm1lIjowfQ==")
// 	if err != nil {
// 		log.Fatalf("Failed to fetch extracted URL: %v", err)
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		log.Fatalf("Failed to fetch extracted URL: status code %d", response.StatusCode)
// 	}

// 	// Proses respons dari URL yang diekstrak
// 	fmt.Println("Successfully fetched the extracted URL.")

// 	baseURL := "http://localhost:5774/rest/PesertaDidik"
// 	params := url.Values{}
// 	params.Add("sekolah_id", "8a5bd957-66bc-4096-9ff1-fee096b87ba0")
// 	params.Add("pd_module", "pdterdaftar")
// 	params.Add("semester_id", "20221")

// 	urlWithParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())

// 	// resp, err := http.Get(urlWithParams)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Response:", string(body))

// }
