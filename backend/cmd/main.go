// yang package main bakal di execute first time
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

// bikin var buat db
var db *sql.DB

// var jwt
var jwtKey = []byte("8ae13976e986991472e41829b2a2b3f1ca2d48b40f93a1cd4ac06d77be4fa1d175876c3f594d313a2ab8c0b1411200d55305dc9c4a1db10031cc490e4e34fda58f1cc771df64e5aa488b46061db4a4948cab1d4a4d5f5652868739c17f547738063d6f52dc74da12bcf33b917032f5568101cf406e44a55d8186690138560cd086f609b485ead02871d0fa5c278d84fe567dfb92df7a0bf2e81b8018af633ec85ed39e1224f30b9e8245f22790b889f46191e672af78bb20f0e6720f1033bf0f5d8a8b5a272266d43b88cd69eabbdc09f75de363eb1833619055d155db6723d56405d942a4203818deb567385c5f11ee5cf65f5713879f6a0cc248350c27a0a9")

// tambahan type user struct
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// pas bikin login handler, pertama bikin struct dulu buat loginnya
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type loginResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

func main() {
	dbInit()
	routerInit()
	// defer dbnya dipindah ke sini daripada harus open close di function terpisah yg manggil db. Kira2 itu penyebabnya kenapa klo ditaruh di dbInit itu db connectionnya selalu close
	defer db.Close()
}

func dbInit() {
	// declare & initialize variable pake shorthand := dibanding pake var type = value
	// terus load env loadnya error krn beda directory
	// solusinya dikasih fullpath lokasi env filenya
	err := godotenv.Load("../.env")
	//jujur masih bingung sama err != nill, perlu ditelaah lagi
	if err != nil {
		log.Print("Error loading .env file ", err)
	}

	//connect to db, ternyata sql.open() itu nge-return 2 values db n error. lucunya pas tanda : diilangin di db, err = sql.open jadi bisa, tapi defer closenya juga dicomment
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	// Ensure the database connection is alive
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Database connection failed:", err)
	// }

	//klo ini ga dicomment, bakal error Failed to insert user into database
	// defer db.Close()

	// // // Periodically check database connection
	// go func() {
	// 	for {
	// 		err := db.Ping()
	// 		if err != nil {
	// 			log.Println("Database connection lost:", err)
	// 		} else {
	// 			log.Println("Database connection is alive")
	// 		}
	// 		time.Sleep(10 * time.Second) // check every 10 seconds
	// 	}
	// }()
}

func routerInit() {
	//Init routernya dulu
	router := mux.NewRouter()

	//terus define http routes nya
	router.HandleFunc("/api/register", registerHandler).Methods("POST")
	router.HandleFunc("/api/login", loginHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	//start http server
	log.Println("server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

// fungsi buat ngehash password pake bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// function handler buat register
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	//tambahin handler buat missing form values
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
	}

	// hashed password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Unable to hash password", http.StatusInternalServerError)
	}
	user.Password = hashedPassword

	// Log the query and parameters
	// log.Printf("Executing query: INSERT INTO users (username, email, password) VALUES (%s, %s, %s)\n", user.Name, user.Email, user.Password)

	// //insert value ke db, pertama querynya langsung masuk di exec, sekrang dipisah, dibikin var sendiri
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err = db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered succesfully",
		"user":    user,
	})
}

// Function handler buat login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	//parse request body
	var loginreq loginRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginreq); err != nil {
		log.Printf("Error decoding login request: ", err)
		http.Error(w, "Invalid reques payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//query user dari db
	var dbPassword string
	var name string
	query := "SELECT password, name FROM users WHERE email = $1"
	row := db.QueryRow(query, loginreq.Email)
	if err := row.Scan(&dbPassword, &name); err != nil {
		log.Printf("Error querying user ", err)
		http.Error(w, "Invalid Credentials ", http.StatusUnauthorized)
		return
	}

	// bandingin password dari db sama yang diinput user
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginreq.Password)); err != nil {
		log.Printf("Password does not match ", err)
		http.Error(w, "Invalid Credentials ", http.StatusUnauthorized)
		return
	}

	// generate jwt token
	token, err := generateJWT(loginreq.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// respon token
	response := loginResponse{Token: token, Name: name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// function generate token
func generateJWT(email string) (string, error) {
	expTime := time.Now().Add(24 * time.Hour)
	claims := &jwtClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
