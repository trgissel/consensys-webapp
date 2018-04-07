package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// mock contract
type CarContract struct {
	ID                     string   `json:"id,omitempty"`
	Manufacture            string   `json:"manufacture,omitempty"`
	EmissionsCertification string   `json:"firstname,omitempty"`
	Owners                 []string `json:"lastname,omitempty"`
}

var contracts []CarContract

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func getContractsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(contracts)
}

func getContractEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range contracts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func createContractEndpoint(w http.ResponseWriter, req *http.Request) {
	var contract CarContract
	// _ = json.NewDecoder(req.Body).Decode(&contract)
	contract.ID = generateUUID()
	contracts = append(contracts, contract)
	json.NewEncoder(w).Encode(contract)
}

func createContractEndpointWithID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var contract CarContract
	// _ = json.NewDecoder(req.Body).Decode(&contract)
	contract.ID = params["id"]
	var foundDup = false
	for _, item := range contracts {
		if item.ID == contract.ID {
			foundDup = true
			break
		}
	}
	if foundDup {
		http.Error(w, "Conflict", http.StatusConflict)
	} else {
		contracts = append(contracts, contract)
		json.NewEncoder(w).Encode(contract)
	}
}

var mySigningKey = []byte("secret")

func login(w http.ResponseWriter, req *http.Request) {
	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)
	/* Set token claims */
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
	//json.NewEncoder(w).Encode(user)
}

func isTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_, ok := claims["name"]
		if ok {
			return true
		}
	} else {
		log.Println(err)
	}
	return false
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/contract", createContractEndpoint).Methods("POST")
	router.HandleFunc("/contract", getContractsEndpoint).Methods("GET")
	router.HandleFunc("/contract/{id}", createContractEndpointWithID).Methods("POST")
	router.HandleFunc("/contract/{id}", getContractEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
