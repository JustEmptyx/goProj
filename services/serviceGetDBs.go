package services

import (
	"database/sql"
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"interfaces/services/game"
	_ "net/http"
)
type units struct {
	Name     string  `json:"name"`
	Health   float32 `json:"health"`
	Mana     float32 `json:"mana"`
	Stamina  float32 `json:"stamina"`
	Defence  float32 `json:"defence"`
	Unittype string  `json:"unittype"`
	Owner string `json:"owner"`
}

func ParseMySQL(db *sql.DB) []game.Hero {
	res,err := db.Query("SELECT  `name`, `health`, `mana`, `stamina`, `defence`, `unittype`,`owner` FROM `units`")
	unitsArr := make([]units,0)
	Heroes := make([]game.Hero,1)
	for res.Next(){
		unit := units{}
		err = res.Scan(&unit.Name,&unit.Health,&unit.Mana,&unit.Stamina,&unit.Defence,&unit.Unittype,&unit.Owner)
		if err !=nil{
			panic(err)
		}
		unitsArr = append(unitsArr, unit)
		if unit.Unittype == "Warrior"{
			warrior := &game.Warrior{}
			warrior.SetName(unit.Name)
			warrior.SetHealth(unit.Health)
			warrior.SetMana(unit.Mana)
			warrior.SetStamina(unit.Stamina)
			warrior.SetDefence(unit.Defence)
			warrior.SetUnittype(unit.Unittype)
			warrior.SetAttacksPool()
			warrior.SetOwner(unit.Owner)
			Heroes = append(Heroes, warrior)
		}else if unit.Unittype == "Mage"{
			mage := &game.Mage{}
			mage.SetName(unit.Name)
			mage.SetHealth(unit.Health)
			mage.SetMana(unit.Mana)
			mage.SetStamina(unit.Stamina)
			mage.SetDefence(unit.Defence)
			mage.SetUnittype(unit.Unittype)
			mage.SetOwner(unit.Owner)
			mage.SetAttacksPool()
			Heroes = append(Heroes, mage)
		}else if unit.Unittype == "Rogue"{
			rogue := &game.Rogue{}
			rogue.SetName(unit.Name)
			rogue.SetHealth(unit.Health)
			rogue.SetMana(unit.Mana)
			rogue.SetStamina(unit.Stamina)
			rogue.SetDefence(unit.Defence)
			rogue.SetUnittype(unit.Unittype)
			rogue.SetAttacksPool()
			rogue.SetOwner(unit.Owner)
			Heroes = append(Heroes, rogue)
		}

	}
	if err!=nil{
		panic(err)
	}
	heroes := append(Heroes[:0], Heroes[1:]...)
	return heroes
}
