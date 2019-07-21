package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	// "io/ioutil"
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

var router = gin.Default()

func main() {

	handleRequest()
	router.Run()
}

func handleRequest() {

	router.GET("/", home)
	router.GET("/users", getUser)
	router.GET("/users/:id", getOneUser)
	// router.GET("/users/:id1/:id2", getTwoUser)
}

func home(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Saksit Jantaraplin")
}

func getUser(c *gin.Context) {
	req, err := http.Get("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Writer, string(response))
}

func getOneUser(c *gin.Context) {

	userID := c.Param("id")
	req, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Writer, string(response))
}

// func getTwoUser(c *gin.Context) {

// 	userID1 := c.Param("id1")
// 	req1, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer req1.Body.Close()
// 	response1, err := ioutil.ReadAll(req1.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	userID2 := c.Param("id2")
// 	req2, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer req2.Body.Close()
// 	response2, err := ioutil.ReadAll(req2.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

	// fmt.Fprintf(c.Writer, string(response))
// }

// func getJSON(userID string) string {

// 	req, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer req.Body.Close()

// 	response, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return string(response)
// }