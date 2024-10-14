package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"jazida-api/internal/entity"
	"jazida-api/internal/infra/db"
	"log"
	"net/http"
	"os"
)

type LoadHandler struct {
	db *db.Queries
}

func NewLoadHandler(db *db.Queries) *LoadHandler {
	return &LoadHandler{
		db: db,
	}
}

func (l *LoadHandler) GetLoads(w http.ResponseWriter, r *http.Request) {
	loads, err := l.db.GetLoads(r.Context())
	if err != nil {
		log.Fatal("Error getting loads", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(loads)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (l *LoadHandler) SaveLoad(w http.ResponseWriter, r *http.Request) {

	var newLoad entity.Load
	if err := json.NewDecoder(r.Body).Decode(&newLoad); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	l.db.CreateLoad(r.Context(), db.CreateLoadParams{
		Client:        newLoad.Client,
		Plate:         newLoad.Plate,
		Material:      newLoad.Material,
		Quantity:      newLoad.Quantity,
		Paymentmethod: newLoad.PaymentMethod,
		Signature:     newLoad.Signature,
	})
}

func (l *LoadHandler) SaveSignature(w http.ResponseWriter, r *http.Request) {

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

	url := entity.Signature{
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
