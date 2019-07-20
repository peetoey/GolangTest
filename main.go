package main

import (
"fmt"
// "io/ioutil"
// "log"
// "net/http"

"github.com/gin-gonic/gin"
)

type userGeo struct {
Lat string
Lng string
}

type userAddress struct {
Street string
Suit string
City string
Zipcode string
Geo userGeo
}

type userCompany struct {
Name string
CatchPhrase string
Bs string
}

type user struct {
ID int
Name string
Username string
Email string
Address userAddress
Phone string
Website string
Company userCompany
}

type addressBook struct {
Firstname string
Lastname string
Code string
Phone string
}

var router = gin.Default()

func main() {

handleRequest()
router.Run()
}

func handleRequest() {

router.GET("/", home)
router.GET("/users", getUser)
router.GET("/users/:id", get)
// router.GET("/users/:id/:id2", getSecondUser)

}

func home(c *gin.Context) {
fmt.Fprintf(c.Writer, "Saksit Jantaraplin")
}

func getUser(c *gin.Context) {

// userLists := []user{}
// userDetail1 := user{
// Id: 1,
// Name: "Saksit Jantaraplin",
// Username: "Toey",
// Email: "saksit.ja@kbtg.tech",
// Address: userAddress{
// Street: "fwfasdsdw",
// Suit: "dsfdsgds",
// City: "afsafsa",
// Zipcode: "aadsa",
// Geo: userGeo{
// Lat: -123,
// Lng: 456,
// },
// },
// Phone: "345435435",
// Website: "ofokslkdsf",
// Company: userCompany{
// Name: "a;;lsakd",
// CatchPhrase: "lkaklflksdfj",
// Bs: "ksjfshgid",
// },
// }

// userDetail2 := user{
// Id: 2,
// Name: "Teerapat Voradanunt",
// Username: "Boat",
// Email: "teerapat.vo@kbtg.tech",
// Address: userAddress{
// Street: "fwfasdsdw",
// Suit: "dsfdsgds",
// City: "afsafsa",
// Zipcode: "aadsa",
// Geo: userGeo{
// Lat: -123,
// Lng: 456,
// },
// },
// Phone: "345435435",
// Website: "ofokslkdsf",
// Company: userCompany{
// Name: "a;;lsakd",
// CatchPhrase: "lkaklflksdfj",
// Bs: "ksjfshgid",
// },
// }

// userLists = append(userLists, userDetail1)
// userLists = append(userLists, userDetail2)

// userID := r.URL.Path[len("/users/"):]
// if userID == "" {

// userListsJSON, err := json.Marshal(userLists)

// if err != nil {
// fmt.Println(err)
// }

// fmt.Fprintf(w, string(userListsJSON))

// } else {

// for index := 0; index < len(userLists); index++ {

// if userID == strconv.Itoa(userLists[index].Id) {
// json.NewEncoder(w).Encode(userLists[index])
// }

// }
// }

userID := c.Param("id")
userID2 := c.Param("id2")
fmt.Fprintf(c.Writer, getJSON(userID, userID2))
}

func getSecondUser(c *gin.Context) {
userID1 := c.Param("id")
userID2 := c.Param("id2")
fmt.Fprintf(c.Writer, getJSON(userID1, userID2))
}

func getJSON(userID1, userID2 string) string {

return userID1 + " - " + userID2

// if userID2 == "" {
// req, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID2)
// if err != nil {
// fmt.Println(err)
// }
// defer req.Body.Close()

// response, err := ioutil.ReadAll(req.Body)
// if err != nil {
// fmt.Println(err)
// }

// return string(response)

// }

// else {

// // // getUser1
// // req1, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID1)
// // if err != nil {
// // fmt.Println(err)
// // }
// // defer req1.Body.Close()

// // response1, err := ioutil.ReadAll(req1.Body)
// // if err != nil {
// // fmt.Println(err)
// // }
// // log.Println(response1)

// //getUser2
// req2, err := http.Get("https://jsonplaceholder.typicode.com/users/" + userID2)
// if err != nil {
// fmt.Println(err)
// }
// defer req2.Body.Close()

// response2, err := ioutil.ReadAll(req2.Body)
// if err != nil {
// fmt.Println(err)
// }

// return string(response2)
// }

}