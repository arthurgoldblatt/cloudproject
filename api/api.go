package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"errors"
)


//Declare a global array of Credentials
//See credentials.go

/*YOUR CODE HERE*/
creds := []Credentials{}


func RegisterRoutes(router *mux.Router) error {

	/*

	Fill out the appropriate get methods for each of the requests, based on the nature of the request.

	Think about whether you're reading, writing, or updating for each request


	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)

	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	access_token := ""
	cookie, err := request.Cookie("access_token")
	if err = nil {
		access_token = cookie.Value
	}
	fmt.fprintf(response, access_token)

	return
}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	userID := request.URL.Query.Get("userID")
	fmt.fprintf(response, userID)

	return
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.request

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	err := JSON.NewDecoder(request.Body).Decode(&credInst)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.fprintf(response, credInst.Username + "\n")
	fmt.fprintf(response, credInst.Password)

	return
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	err := JSON.NewDecoder(request.Body).Decode(&credInst)
	for _, cred := range creds {
		if cred.Username == credInst.Username {
			err = 1
		}
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	creds = creds.append(creds, credInst)

	return
}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array

		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	err := JSON.NewDecoder(request.Body).Decode(&credInst)
	index := -1
	for i := 0; i < len(creds); i++ {
		if creds[i].Username == credInst.Username {
			index = i
			break
		}
	}
	if index == -1 {
		err = -1
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	fmt.fprintf(response, index)

	return
}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	err := JSON.NewDecoder(request.Body).Decode(&credInst)
	password := ""
	for i := 0; i < len(creds); i++ {
		if creds[i].Username == credInst.Username {
			password = creds[i].Password
			break
		}
	}
	if password == "" {
		err = -1
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	fmt.fprintf(response, password)

	return
}



func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	found := false
	for i := 0; i < len(creds); i++ {
		if creds[i].Username == credInst.Username {
			creds[i].Password = credInst.Password
			found = true
			break
		}
	}
	if err || !found {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	return
}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credInst := Credentials{}
	err := JSON.NewDecoder(request.Body).Decode(&credInst)
	index := -1
	for i := 0; i < len(creds); i++ {
		if creds[i].Username == credInst.Username {
			creds = Remove(creds, i)
			index = i
			break
		}
	}
	if index == -1 {
		err = -1
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	return
}
