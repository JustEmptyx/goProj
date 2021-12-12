
package server

import (
	"context"
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"interfaces/services"
	"interfaces/services/game"
	"interfaces/services/session"
	"log"
	"net/http"
	"strconv"
	_ "strconv"
)


func initDBMySQL()*sql.DB{
	db,err :=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/mydbtest")
	if err!=nil{
		panic(err)
	}
	return db
}
func initDBMongo()*mongo.Client{
	client, errmdb := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if errmdb != nil {
		log.Fatal(errmdb)
	}
	errmdb = client.Connect(context.TODO())
	if errmdb != nil {
		log.Fatal(errmdb)
	}
	errmdb = client.Ping(context.TODO(), nil)
	if errmdb != nil {
		log.Fatal(errmdb)
	}
	return client
}

func closeDB(db *sql.DB) bool{
	 defer db.Close()
	 return true
}
func notLoggedMsg(w http.ResponseWriter){
	fmt.Fprintf(w,"You are not logged in")
}

func handleRequestAddRandomUnit(db *sql.DB, collection *mongo.Collection){
	http.HandleFunc("/addRandomUnit",func(w http.ResponseWriter,req *http.Request){
		SetCurrentSession()
		var login string
		cookie,err1 := req.Cookie(services.COOKIE_NAME)
		if cookie!=nil {
			login, _ = currentSession.Get(cookie.Value)
		}
		if(err1 == nil && login != ""){
			userFromDB := services.LocalUser{}
			filter := bson.D{{"name", login}}
			_ = collection.FindOne(context.TODO(), filter).Decode(&userFromDB)
			if(userFromDB.Points > 0) {
				hero := game.CreateRandomOne(login)
				services.ChangePointLoser(collection, login)
				str := services.Refactor(hero)
				services.Push(db, str)
				fmt.Fprintf(w, "generated")
				hero.PrintFullInfo(w)
			} else{
				fmt.Fprintf(w," not enough points for creating unit")
			}
		} else{
			notLoggedMsg(w)
		}
	})
}
func handleRequestClearMySQL(db *sql.DB){
	http.HandleFunc("/clear",func(w http.ResponseWriter,req* http.Request){
		SetCurrentSession()
		cookie,err:=req.Cookie(services.COOKIE_NAME)
		if(err == nil) {
			login, points := currentSession.Get(cookie.Value)
			if (login != "" || points != 0) {
				services.ClearDBMySQL(db)
			}
		} else {
				notLoggedMsg(w)
			}

	})
}

var currentSession *session.Session
func SetCurrentSession(){
	currentSession = services.GetCurrentSession()
}
func isPow2(num uint) bool {
	switch true {
	case num == 1:
		return true
	case num == 0:
		return false
	}

	if num % 2 == 1 {
		return false
	}

	return isPow2(num / 2)
}

func handleRequestHelp(){
	http.HandleFunc("/help",func(w http.ResponseWriter,req* http.Request){
		fmt.Fprintf(w,"/addRandomUnit - create random unit")
		fmt.Fprintf(w,"/clear - clear unit database")
		fmt.Fprintf(w,"/addSpecialUnit?...&.. - add unit with specialized params: name,health,mana,stamina,defence,unittype")
		fmt.Fprintf(w,"/currentunits - check the units in database")
		fmt.Fprintf(w,"/battle")
		fmt.Fprintf(w,"/register - create new account")
		fmt.Fprintf(w,"/login - log into an existing account")

	})
}

func handleRequestAddSpecialUnit(db *sql.DB){
	http.HandleFunc("/addSpecialUnit",func(w http.ResponseWriter,req *http.Request){
		SetCurrentSession()
		cookie,err:=req.Cookie(services.COOKIE_NAME)
		if(err == nil) {
			login, _ := currentSession.Get(cookie.Value)
			if(login != "") {
				hero := services.GetSnippetHero(req)
				str := services.Refactor(hero)
				services.Push(db, str)
				fmt.Fprintf(w, "generated")
				hero.PrintFullInfo(w)
			}
		} else{
			notLoggedMsg(w)
		}
	})
}
func handleRequestCurrentUnits(db*sql.DB){
	http.HandleFunc("/currentUnits",func(w http.ResponseWriter,req *http.Request){
		game.Heroes = make([]game.Hero,0)
		game.Heroes = services.ParseMySQL(db)
		if(len(game.Heroes) != 0) {
			for i := range (game.Heroes) {
				if(game.Heroes[i]!=nil) {
					game.Heroes[i].PrintFullInfo(w)
				}
			}
		}
	})
}

func handleRequestStats(){
	http.HandleFunc("/stats",func(w http.ResponseWriter,req *http.Request){
		SetCurrentSession()
		cookie,err:=req.Cookie(services.COOKIE_NAME)
		if(err == nil) {
			if(cookie!=nil) {
				login, points := currentSession.Get(cookie.Value)
				fmt.Fprintf(w,login + " you have " + strconv.Itoa(points) + " points \n")
			}else{
				notLoggedMsg(w)
			}

		} else{
			notLoggedMsg(w)
		}
	})
}
func Main(){
	var winner game.Hero
	var multiplyer int
	var name string

	db := initDBMySQL()
	client := initDBMongo()
	collection := client.Database("users").Collection("users")
	handleRequestHelp()
	services.HandleRequestLogin(collection)
	services.HandleRequestRegister(collection)
	handleRequestAddRandomUnit(db,collection)
	handleRequestClearMySQL(db)
	handleRequestAddSpecialUnit(db)
	handleRequestStats()

	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "main page")
		SetCurrentSession()
	})
	http.HandleFunc("/battle",func(w http.ResponseWriter,req *http.Request){
		heroes := services.ParseMySQL(db)
		if(isPow2(uint(len(heroes)))) {
			winner, multiplyer = game.StartBattleDB(heroes)
			name = winner.GetName()
			fmt.Fprintf(w, "winner : "+winner.GetOwner()+" with unit : "+winner.GetName()+" earn "+strconv.Itoa(multiplyer)+" points ")
			services.ChangePointsWinner(collection, winner.GetOwner(), multiplyer)
			services.ClearDBMySQL(db)
		} else {
			fmt.Fprintf(w,"Bad count of heroes, use power of 2 nums")
		}
	})
	http.HandleFunc("/battleSandbox",func(w http.ResponseWriter,req *http.Request){
		winner = game.StartBattle(128)
		name =winner.GetName()
		fmt.Fprintf(w, "winner : " + winner.GetOwner() + " with unit : " + winner.GetName() + " earn " + strconv.Itoa(multiplyer) + " points ")
	})
	handleRequestCurrentUnits(db)


	err := http.ListenAndServe(
		":8080",
		nil,
	)
	if err != nil {
		return 
	}
	closeDB(db)
}


//переписать главную функцию для боя для любого числа людей
//разобраться с очками и способом их начисления