package services

import (
	"context"
	"database/sql"
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"interfaces/services/game"
	"log"
	"net/http"
	_ "net/http"
	"strconv"
)
func closeInsert(insert *sql.Rows){
	defer insert.Close()
}
func GetSnippetHero( r *http.Request) game.Hero {
	var hero game.Hero
	var name,unittype string
	var health,mana,stamina,defence float64
	name = ""
	unittype = ""
	name = r.URL.Query().Get("name")
	health, _ = strconv.ParseFloat(r.URL.Query().Get("health"),32)
	mana, _ = strconv.ParseFloat(r.URL.Query().Get("mana"),32)
	stamina, _ =  strconv.ParseFloat(r.URL.Query().Get("stamina"),32)
	defence, _ =  strconv.ParseFloat(r.URL.Query().Get("defence"),32)
	unittype =  r.URL.Query().Get("unittype")
	if unittype == "Warrior"{
		warrior := &game.Warrior{}
		warrior.SetName(name)
		warrior.SetHealth(float32(health))
		warrior.SetMana(float32(mana))
		warrior.SetStamina(float32(stamina))
		warrior.SetDefence(float32(defence))
		warrior.SetUnittype(unittype)
		hero = warrior
	}else if unittype == "Mage"{
		mage := &game.Mage{}
		mage.SetName(name)
		mage.SetHealth(float32(health))
		mage.SetMana(float32(mana))
		mage.SetStamina(float32(stamina))
		mage.SetDefence(float32(defence))
		mage.SetUnittype(unittype)
		hero = mage
	}else if unittype == "Rogue"{
		rogue := &game.Rogue{}
		rogue.SetName(name)
		rogue.SetHealth(float32(health))
		rogue.SetMana(float32(mana))
		rogue.SetStamina(float32(stamina))
		rogue.SetDefence(float32(defence))
		rogue.SetUnittype(unittype)
		hero = rogue
	}else {
		return nil
	}
	return hero
}
func Push(db *sql.DB,str string) {
	insert,err := db.Query(str)
	closeInsert(insert)
	if err!=nil{
		panic(err)
	}
}
func ChangePointsWinner(collection *mongo.Collection,login string,points int){
	filter := bson.D{{"name", login}}
	update := bson.D{
		{"$inc", bson.D{
			{"points", points},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	print(updateResult)
}
func ChangePointLoser(collection *mongo.Collection,login string){
		filter := bson.D{{"name", login}}
		update := bson.D{
			{"$inc", bson.D{
				{"points", -1},
			}},
		}
		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		print(updateResult)
}
