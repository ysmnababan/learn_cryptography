package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type VaultCreds struct {
	Data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"data"`
}

func main() {
	vaultAddr := "http://127.0.0.1:8200"
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN not set")
	}

	// 1. Fetch creds from Vault
	req, _ := http.NewRequest("GET", vaultAddr+"/v1/database/creds/app-role", nil)
	req.Header.Add("X-Vault-Token", vaultToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var creds VaultCreds
	if err := json.NewDecoder(resp.Body).Decode(&creds); err != nil {
		log.Fatal(err)
	}

	username := creds.Data.Username
	password := creds.Data.Password
	fmt.Printf("Got dynamic creds: %s / %s\n", username, password)

	// 2. Open connection to DB
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta", username, password, "mydb")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	name := "somename" + uuid.NewString()
	res := db.Create(&User{
		Username: name,
		Email:    uuid.NewString(),
	})
	if res.Error != nil {
		log.Println(res.Error)
		return
	}

	users := []User{}
	db.Find(&users)
	fmt.Println(users)
}
