package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"./controllers"
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
var usernames []string

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
	if !authenticate(w, req) {
		return
	}
	json.NewEncoder(w).Encode(contracts)
}

func getContractEndpoint(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
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
	if !authenticate(w, req) {
		return
	}
	controllers.EthereumClientConnect()
	var contract CarContract
	// _ = json.NewDecoder(req.Body).Decode(&contract)
	contract.ID = generateUUID()
	contracts = append(contracts, contract)
	json.NewEncoder(w).Encode(contract)
}

func createContractEndpointWithID(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
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

func authenticate(w http.ResponseWriter, req *http.Request) bool {
	var jwt = getJWT(w, req)
	if len(jwt) < 1 {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}
	if !isTokenValid(jwt) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}
	return true
}

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
	addUserIfNotExist(user.Username)
	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
	//json.NewEncoder(w).Encode(user)
}

func addUserIfNotExist(user string) {
	for _, item := range usernames {
		if item == user {
			return
		}
	}
	usernames = append(usernames, user)
}

func hasUserLoggedIn(user string) bool {
	for _, item := range usernames {
		if item == user {
			return true
		}
	}
	return false
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
		log.Println("valid token. Checking if logged in")
		if name, ok := claims["name"].(string); ok && hasUserLoggedIn(name) {
			return true
		}
		return false
	}
	log.Println(err)
	return false
}

func getJWT(w http.ResponseWriter, r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	s := strings.Split(authorizationHeader, " ")
	fmt.Printf("%v", s)
	if len(s) < 2 || s[0] != "Bearer" {
		return ""
	}
	return s[1]
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
