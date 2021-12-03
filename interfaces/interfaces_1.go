package main

import (
	"fmt"
	"math/rand"
	"time"
)

const count = 64

type AttacksPool struct{
	name string
	dmg float32
	cost float32
}
type Warrior struct {
	name string
	health float32
	mana float32
	stamina float32
	attackpool []AttacksPool
	defence float32
	unitType string
}
type Mage struct {
	name string
	health float32
	mana float32
	stamina float32
	attackpool []AttacksPool
	defence float32
	unitType string
}
type Rogue struct {
	name string
	health float32
	mana float32
	stamina float32
	attackpool []AttacksPool
	defence float32
	unitType string
}

type Hero interface{
	FindDmg(Hero)
	ChooseSkill() int
	printInfo()
	setHP(float32)
	getHP() float32
	getName() string
}

func(mage *Mage) setHP(num float32){
	mage.health = mage.health - num
}
func(war *Warrior) setHP(num float32){
	war.health = war.health - num
}
func(rog *Rogue) setHP(num float32){
	rog.health = rog.health - num
}

func(mage *Mage) getHP() float32{
	return mage.health
}
func(war *Warrior) getHP() float32{
	return war.health
}
func(rog *Rogue) getHP() float32{
	return rog.health
}


func(mage *Mage) getName() string{
	return mage.name
}
func(war *Warrior) getName() string{
	return war.name
}
func(rog *Rogue) getName() string{
	return rog.name
}

func(m *Mage) FindDmg(hero Hero){
	if m.mana >= m.attackpool[m.ChooseSkill()].cost {
		damage := m.attackpool[m.ChooseSkill()].dmg
		hero.setHP(damage)
		m.mana -= m.attackpool[m.ChooseSkill()].cost
		fmt.Printf(m.name + " done succ attack and dealed %.2f dmg by %s ," + hero.getName() + " have %.2f hp\n",m.attackpool[m.ChooseSkill()].dmg,m.attackpool[m.ChooseSkill()].name,hero.getHP())
	}else{
		fmt.Print("No mana\n")
	}
	m.mana += 1
}
func(wr *Warrior) FindDmg(hero Hero){
	if wr.stamina >= wr.attackpool[wr.ChooseSkill()].cost {
		damage := wr.attackpool[wr.ChooseSkill()].dmg
		hero.setHP(damage)
		wr.stamina -= wr.attackpool[wr.ChooseSkill()].cost
		fmt.Printf(wr.name + " done succ attack and dealed %.2f dmg by %s , " + hero.getName() +" have %.2f hp\n",wr.attackpool[wr.ChooseSkill()].dmg,wr.attackpool[wr.ChooseSkill()].name,hero.getHP())
	}else{
		fmt.Print("No stamina\n")
	}
	wr.stamina += 1
}
func(rog *Rogue) FindDmg(hero Hero){
	if rog.stamina >= rog.attackpool[rog.ChooseSkill()].cost {
		damage := rog.attackpool[rog.ChooseSkill()].dmg
		wasCrit := isCrit()
		if(wasCrit) {
			hero.setHP(2*damage)
		} else{
			hero.setHP(damage)
		}
		rog.stamina -= rog.attackpool[rog.ChooseSkill()].cost
		if(wasCrit) {
			fmt.Printf(rog.name+" done succ attack and dealed %.2f dmg by %s , "+hero.getName()+" have %.2f hp\n", 2*rog.attackpool[rog.ChooseSkill()].dmg, rog.attackpool[rog.ChooseSkill()].name, hero.getHP())
		} else{
			fmt.Printf(rog.name+" done succ attack and dealed %.2f dmg by %s , "+hero.getName()+" have %.2f hp\n", rog.attackpool[rog.ChooseSkill()].dmg, rog.attackpool[rog.ChooseSkill()].name, hero.getHP())
		}
	}else{
		fmt.Print("No stamina\n")
	}
	rog.stamina += 1
}

func(m Mage) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(m.attackpool))
}
func(wr Warrior) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(wr.attackpool))
}
func(rog Rogue) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(rog.attackpool))
}

func setWarrior(wr* Warrior) *Warrior{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)

	names := []string{"Dave","John","Kenny","Ron","Henry","Nugget","Sara","Ronald","Andy","Susan"}
	warrior := *wr

	warrior.name = names[num]
	warrior.health = 100
	warrior.stamina = 100
	warrior.defence = 1 + 0.1*float32(rand.Intn(4))
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"swordPunch",7,7})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"lunge",6,6})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"rock",2,0})
	return &warrior
}
func setMage(mg* Mage) *Mage{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	names := []string{"Dave","John","Kenny","Ron","Henry","Nugget","Sara","Ronald","Andy","Susan"}
	mage := *mg

	mage.name = names[num]
	mage.health = 100
	mage.mana = 100
	mage.defence = 0.5 - 0.1*float32(rand.Intn(3))
	mage.attackpool = append(mage.attackpool, AttacksPool{"blizzard",10,9})
	mage.attackpool = append(mage.attackpool, AttacksPool{"fireball",9,8})

	mage.attackpool = append(mage.attackpool, AttacksPool{"rock",1,0})
	return &mage
}
func setRogue(rog*Rogue) *Rogue{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)

	names := []string{"Dave","John","Kenny","Ron","Henry","Nugget","Sara","Ronald","Andy","Susan"}
	rogue := *rog

	rogue.name = names[num]
	rogue.health = 100
	rogue.stamina = 100
	rogue.defence = 1 + 0.1*float32(rand.Intn(2))
	rogue.attackpool = append(rogue.attackpool, AttacksPool{"blow from below",9,9})
	rogue.attackpool = append(rogue.attackpool, AttacksPool{"backstab",25,30})
	rogue.attackpool = append(rogue.attackpool, AttacksPool{"rock",2,0})
	return &rogue
}

func createWarrior() *Warrior{
	warrior := Warrior{}
	warrior.unitType = "Warrior"
	return setWarrior(&warrior)
}
func createMage() *Mage{
	mage := Mage{}
	mage.unitType = "Mage"
	return setMage(&mage)
}
func createRogue() *Rogue{
	rogue := Rogue{}
	rogue.unitType = "Rogue"
	return setRogue(&rogue)
}

func (w *Warrior) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",w.name,w.unitType,w.health,w.stamina,w.mana,w.defence)
}
func (m *Mage) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",m.name,m.unitType,m.health,m.stamina,m.mana,m.defence)
}
func (r *Rogue) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",r.name,r.unitType,r.health,r.stamina,r.mana,r.defence)
}

func fight(h1 Hero,h2 Hero)  int{
	h1.FindDmg(h2)
	if h2.getHP() <= 0 {

		fmt.Println(h2.getName() + " dead")
		return 1
	} else {
		h2.FindDmg(h1)
		if h1.getHP() <= 0 {
			fmt.Println(h1.getName() + " dead")
			return 2
		}
	}
	return 0
}

func theLastOneRemained(heroes []Hero) bool{
	if(len(heroes) == 1){
		return true
	} else{
		return false
	}
}
func isCrit()bool{
	if(rand.Intn(5) == 0){
		return true
	} else{
		return false
	}
}
func printName(hero Hero) string{
	return hero.getName()
}
func printInfo(hero Hero){
	hero.printInfo()
}
func main(){
	heroes := make([]Hero,count)
	rand.Seed(time.Now().UnixNano())
	for i := range heroes{
		num := rand.Intn(3)
		switch num{
		case 0:
			temp:= createWarrior()
			heroes[i] = temp
		case 1:
			temp:= createMage()
			heroes[i] = temp
		case 2:
			temp:= createRogue()
			heroes[i] = temp
		}
	}
	println("Players : ")
	for i:= range(heroes){
		printInfo(heroes[i])
	}
	for ;!theLastOneRemained(heroes);{
		personNum1 := rand.Intn(len(heroes))
		time.Sleep(1)
		personNum2 := rand.Intn(len(heroes))
		if (personNum1 == personNum2){
			for ;personNum1 == personNum2;{
				loc := rand.Intn(len(heroes))
				personNum2  = loc
			}
		}
		winner := fight(heroes[personNum1], heroes[personNum2])
		if (winner == 1) {
			heroes = append(heroes[:personNum2], heroes[personNum2+1:]...)
		} else if(winner == 2){
			heroes = append(heroes[:personNum1], heroes[personNum1+1:]...)
		}
				
	}
	print(heroes[0].getName())
	print(" survived")
	printInfo(heroes[0])

}
