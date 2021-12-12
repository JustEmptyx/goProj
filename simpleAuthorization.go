package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"interfaces/services/session"

	"log"
	"net/http"
	"strconv"
	"time"
)

type LocalUser struct {
	Name string
	Password string
	Points float64
	Date string
}
const (
	COOKIE_NAME = "sessionId"
)

func SetCurrentSession(cs *session.Session){
	CurrentSession = cs
}

func SetCookieSession(lcookie *http.Cookie,w http.ResponseWriter){
	cookie = lcookie
	writer = w
}
func GetCurrentSession() (cs *session.Session){
	return CurrentSession
}

func GetCookieSession() (lcookie *http.Cookie,w http.ResponseWriter){
	return cookie,writer
}

var writer http.ResponseWriter
var cookie *http.Cookie
var CurrentSession *session.Session

func HandleRequestLogin(collection* mongo.Collection,){
	//login := ""
	//isLogged := false
	//points := 0.
	http.HandleFunc("/login",func(w http.ResponseWriter,req* http.Request){
		CurrentSession = session.NewSession()
		user := getSnipperUser(req)
		userFromDB := LocalUser{}
		filter := bson.D{{"name", user.Name}}
		err := collection.FindOne(context.TODO(), filter).Decode(&userFromDB)
		if err != nil {
			fmt.Fprintf(w,"No such login in the DataBase. Please register using /register?\"login\" = \"your_login\"&\"password\" = your_password")
		}
		if(userFromDB.Password == user.Password && user.Password != ""){
			if(userFromDB.Date != user.Date){
				user.Points =+ 1
				insertResult, _ := collection.InsertOne(context.TODO(), user)
				fmt.Println("player was updated : ",insertResult.InsertedID)
			}
			login := user.Name
			points := userFromDB.Points
			sessionId := CurrentSession.Init(login,int(points))
			cookie := &http.Cookie{
				Name:    COOKIE_NAME,
				Value:   sessionId,
				Expires: time.Now().Add(5*time.Minute),
				Path: "/",
			}
			SetCurrentSession(CurrentSession)
			SetCookieSession(cookie,w)
			http.SetCookie(writer,cookie)
			http.Redirect(w,req,"/",302)
			points = user.Points
		}
	})
	return
}
func HandleRequestRegister(collection* mongo.Collection){
	http.HandleFunc("/register",func(w http.ResponseWriter,req* http.Request){
		user := getSnipperUser(req)
		if(user.Name ==""){
			fmt.Fprintf(w,"You can't use empty login")
			return
		}
		userFromDB := LocalUser{}
		filter := bson.D{{"name", user.Name}}
		flt := collection.FindOne(context.TODO(), filter).Decode(&userFromDB)
		if flt != nil {
			user.Points = 10
			insertResult, err := collection.InsertOne(context.TODO(), user)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w,"Welcome, " + user.Name +". You are successfully registered.")
			fmt.Println("new player : ",insertResult.InsertedID)

		} else{
			fmt.Fprintf(w,"Such player already exist, please log in")
		}

	})
}
func getSnipperUser(r *http.Request) *LocalUser {
	user := LocalUser{}
	user.Name = r.URL.Query().Get("login")
	user.Password = r.URL.Query().Get("password")
	y,m,d := time.Now().Date()
	user.Date = strconv.Itoa(y) + "" + strconv.Itoa(int(m)) + "" + strconv.Itoa(d)
	return &user
}

func HandleRequestPersonalArea(login string,points float64,isSomeoneLogged bool){
	http.HandleFunc("/personalArea",func(w http.ResponseWriter,req* http.Request){
		if(!isSomeoneLogged){
			fmt.Fprint(w,"You are not logged in ! ")
		} else{
			fmt.Fprintf(w, "Hello, " + login )
			fmt.Fprintf(w, "You have : , " +  fmt.Sprint(points))
		}
	})
}