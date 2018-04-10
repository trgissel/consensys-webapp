package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"./contracts"

	"./controllers"
	"./internal"
	"./models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var contractList []models.DeployContractTransaction
var usernames []string
var ethereumConfig internal.EthereumConfig

var addressToCar = make(map[string]*contracts.Car)

// MilesWrapper for parsing body
type MilesWrapper struct {
	miles uint32 `json:"miles"`
}

func getContractsEndpoint(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
	json.NewEncoder(w).Encode(contractList)
}

func getContractEndpoint(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
	params := mux.Vars(req)
	for _, item := range contractList {
		if item.Address == params["id"] {
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
	address, hash, car, err := controllers.CreateCarContract(ethereumConfig.URL, ethereumConfig.KeyStorePath, ethereumConfig.Passphrase)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	var contract models.DeployContractTransaction
	contract.Address = address
	contract.TransactionID = hash
	contractList = append(contractList, contract)
	addressToCar[address] = car
	json.NewEncoder(w).Encode(contract)
}

func createContractEndpointWithID(w http.ResponseWriter, req *http.Request) {
	// do not allow POSTing for tx
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func getMileage(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
	params := mux.Vars(req)
	addresss := params["id"]
	car, ok := addressToCar[addresss]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	mileage, err := controllers.GetCarMileage(car)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	mileageI, err := strconv.ParseInt(mileage, 10, 64)
	if err != nil {
		http.Error(w, "Server Error converting mileage to int", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(mileageI)
}

func addMileage(w http.ResponseWriter, req *http.Request) {
	if !authenticate(w, req) {
		return
	}
	params := mux.Vars(req)
	addresss := params["id"]

	decoder := json.NewDecoder(req.Body)
	var milesWrapper MilesWrapper
	err := decoder.Decode(&milesWrapper)
	if err != nil {
		http.Error(w, "please provide the miles to add in the body", http.StatusBadRequest)
		return
	}

	txID, err := controllers.AddCarMiles(addresss, milesWrapper.miles)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(txID)
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
	ethereumConfig = internal.LoadEthereumConfig("./ethereumConfig.json")
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/contract", createContractEndpoint).Methods("POST")
	router.HandleFunc("/contract", getContractsEndpoint).Methods("GET")
	router.HandleFunc("/contract/{id}", createContractEndpointWithID).Methods("POST")
	router.HandleFunc("/contract/{id}", getContractEndpoint).Methods("GET")
	router.HandleFunc("/contract/{id}/mileage", getMileage).Methods("GET")
	router.HandleFunc("/contract/{id}/mileage", addMileage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
