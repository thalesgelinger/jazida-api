package services

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"jazida-api/db"
	"jazida-api/entities"
	"log"
	"net/http"
	"os"
)

func GetLoads(w http.ResponseWriter, r *http.Request) {
	loads, err := db.GetLoads()
	if err != nil {
		log.Fatal("Error getting loads")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(loads)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SaveLoad(w http.ResponseWriter, r *http.Request) {

	var newLoad entities.Load
	if err := json.NewDecoder(r.Body).Decode(&newLoad); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	db.SaveLoad(&newLoad)
}

func SaveSignature(w http.ResponseWriter, r *http.Request) {

	log.Println("Receiving image...")

	w.Header().Set("Content-Type", "image/jpeg")

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName, err := generateRandomName(10)
	if err != nil {
		http.Error(w, "Error creating file name", http.StatusInternalServerError)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working dir")
		return
	}

	filePath := fmt.Sprintf("%s/signatures/%s.jpg", cwd, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file", http.StatusInternalServerError)
		return
	}

	log.Println(w, "File uploaded successfully")

	domain := "http://localhost:8080/api"

	url := entities.Signature{
		Url: fmt.Sprintf("%s/signatures/%s.jpg", domain, fileName),
	}

	json.NewEncoder(w).Encode(url)
}

func generateRandomName(length int) (string, error) {
	numBytes := length / 2

	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomName := hex.EncodeToString(randomBytes)

	if len(randomName) > length {
		randomName = randomName[:length]
	}

	return randomName, nil
}
