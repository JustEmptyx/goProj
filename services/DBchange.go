package services

import (
	"database/sql"
	"fmt"
	"interfaces/services/game"

	//"github.com/JustEmptyx/goProj"
)
func Refactor(hero game.Hero) string{
	hero.GetName()
	str :="INSERT INTO units VALUES( 0,'"
	str += hero.GetName()+"',"
	str +=fmt.Sprint(hero.GetHP()) +","
	str +=fmt.Sprint(hero.GetMana()) +","
	str += fmt.Sprint(hero.GetStamina()) +","
	str += fmt.Sprint(hero.GetDefence()) +",'"
	str += fmt.Sprint(hero.GetUnitType()) +"','"
	str += fmt.Sprintf(hero.GetOwner()) +"')"
	return str
}
func ClearDBMySQL(db *sql.DB){
	str := "DELETE FROM units"
	_,err := db.Query(str)
	if err!=nil{
		panic(err)
	}

}