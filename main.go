
package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	b64 "encoding/base64"
)


func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
		//get username
		usernames, ok := r.URL.Query()["female"]
		if !ok || len(usernames[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		username := usernames[0]



			// Open database connection
		db, err := sql.Open("mysql", "root:@/mario")
		if err != nil {
			panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
		}
		defer db.Close()

		// Execute the query
		//var nuna = "1"
		//rows, err := db.Query("SELECT * FROM females where id=?",nuna)
		rows, err := db.Query("SELECT * FROM users where username=?",username)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Get column names
		columns, err := rows.Columns()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Make a slice for the values
		values := make([]sql.RawBytes, len(columns))

		// rows.Scan wants '[]interface{}' as an argument, so we must copy the
		// references into such a slice
		// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if value == nil{
			fmt.Println("user does not exist")
			fmt.Fprintln("user does not exist")
		}

		// Fetch rows
		for rows.Next() {
			// get RawBytes from data
			err = rows.Scan(scanArgs...)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			// Now do something with the data.
			// Here we just print each column as a string.
			var value string
			for i, col := range values {
				// Here we can check if the value is nil (NULL value)
				if col == nil {
					value = "NULL"
				} else {
					value = string(col)
				}
				fmt.Println(columns[i], ": ", value)
				fmt.Fprintln(w, string(value))
			}
			fmt.Println("-----------------------------------")
		}
		if err = rows.Err(); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//movie, err := dao.FindById(params["id"])
	fmt.Fprintln(w, params["id"])

	// Open database connection
	db, err := sql.Open("mysql", "root:@/mario")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	//var nuna = "1"
	//rows, err := db.Query("SELECT * FROM females where id=?",nuna)
	rows, err := db.Query("SELECT * FROM females where id=?",params["id"])
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
func TryThings(w http.ResponseWriter, r *http.Request) {
	//get username
    usernames, ok := r.URL.Query()["username"]
    if !ok || len(usernames[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
    username := usernames[0]
	log.Println("Url Param 'username' is: " + string(username))
	fmt.Fprintln(w, string(username))
	dara := b64.StdEncoding.EncodeToString([]byte(username))
	sDec, _ := b64.StdEncoding.DecodeString(dara)
	fmt.Fprintln(w, "dara " +string(dara))
	fmt.Fprintln(w, "sDec "+ string(sDec))



	//get passwprd
	passwords, ok := r.URL.Query()["password"]
    if !ok || len(passwords[0]) < 1 {
        log.Println("Url Param 'password' is missing")
        return
    }
    password := passwords[0]
	log.Println("Url Param 'key' is: " + string(password))
	fmt.Fprintln(w, string(password))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/females", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	r.HandleFunc("/logins", TryThings).Methods("POST")
	if err := http.ListenAndServe(":8400", r); err != nil {
		log.Fatal(err)
	}
}