package game


import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


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
	GetHP() float32
	GetName() string
	GetMana() float32
	GetStamina() float32
	GetDefence() float32
	GetUnitType() string
	setAdditionalStats(float32)
	PrintFullInfo(http.ResponseWriter)
}
//getters
func(mage *Mage) GetUnitType() string{
	unittype:=mage.unitType
	return unittype
}
func(war *Warrior) GetUnitType() string{
	unittype:=war.unitType
	return unittype
}
func(rog *Rogue) GetUnitType() string{
	unittype:=rog.unitType
	return unittype
}
func(mage *Mage) GetDefence() float32{
	defence:=mage.defence
	return defence
}
func(war *Warrior) GetDefence() float32{
	defence:=war.defence
	return defence
}
func(rog *Rogue) GetDefence() float32{
	defence:=rog.defence
	return defence
}
func(mage *Mage) GetStamina() float32{
	stamina:=mage.stamina
	return stamina
}
func(war *Warrior) GetStamina() float32{
	stamina:=war.stamina
	return stamina
}
func(rog *Rogue) GetStamina() float32{
	stamina:=rog.stamina
	return stamina
}
func(mage *Mage) GetMana() float32{
	manacount:=mage.mana
	return manacount
}
func(war *Warrior) GetMana() float32{
	manacount:=war.mana
	return manacount
}
func(rog *Rogue) GetMana() float32{
	manacount:=rog.mana
	return manacount
}
//setters
func(war *Warrior) SetHealth(hp float32) {
	war.health = hp
}
func(mage * Mage) SetHealth(hp float32) {
	mage.health = hp
}
func(rogue *Rogue) SetHealth(hp float32) {
	rogue.health = hp
}
func(war *Warrior) SetName(name string) {
	war.name = name
}
func(mage * Mage) SetName(name string) {
	mage.name = name
}
func(rogue *Rogue) SetName(name string) {
	rogue.name = name
}
func(war *Warrior) SetMana(mana float32) {
	war.mana = mana
}
func(mage* Mage) SetMana(mana float32) {
	mage.mana = mana
}
func(rogue * Rogue) SetMana(mana float32) {
	rogue.mana = mana
}
func(war *Warrior) SetStamina(stamina float32) {
	war.stamina = stamina
}
func(mage *Mage) SetStamina(stamina float32) {
	mage.stamina = stamina
}
func(rogue *Rogue) SetStamina(stamina float32) {
	rogue.stamina = stamina
}
func(war *Warrior) SetDefence(defence float32) {
	war.defence = defence
}
func(mage *Mage) SetDefence(defence float32) {
	mage.defence = defence
}
func(rogue *Rogue) SetDefence(defence float32) {
	rogue.defence = defence
}
func(war *Warrior) SetUnittype(ut string) {
	war.unitType = ut
}
func(mage *Mage) SetUnittype(ut string) {
	mage.unitType = ut
}
func(rogue *Rogue) SetUnittype(ut string) {
	rogue.unitType = ut
}

var Heroes []Hero
func(mage *Mage) setHP(num float32){
	mage.health = mage.health - num
}
func(war *Warrior) setHP(num float32){
	war.health = war.health - num
}
func(rog *Rogue) setHP(num float32){
	rog.health = rog.health - num
}

func createOwnMage(name string,health float32,mana float32,defence float32) *Mage {
	mage := Mage{}
	mage.unitType = "Mage"
	time.Sleep(10)

	mage.name = name
	mage.health = health
	mage.mana = mana
	mage.defence = defence
	mage.attackpool = append(mage.attackpool, AttacksPool{"blizzard",10,9})
	mage.attackpool = append(mage.attackpool, AttacksPool{"fireball",9,8})

	mage.attackpool = append(mage.attackpool, AttacksPool{"rock",1,0})
	return &mage
}

func createOwnWarrior(name string,health float32,stamina float32,defence float32) *Warrior{
	time.Sleep(10)
	warrior := Warrior{}
	warrior.unitType = "Warrior"

	warrior.name = name
	warrior.health = health
	warrior.stamina = stamina
	warrior.defence = defence
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"swordPunch",7,7})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"lunge",6,6})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"rock",2,0})
	return &warrior
}
func createOwnRogue(name string,health float32,stamina float32,defence float32) *Rogue{
	time.Sleep(10)
	rog :=Rogue{}
	rog.unitType = "Rogue"

	rog.name = name
	rog.health = health
	rog.stamina = stamina
	rog.defence = defence
	rog.attackpool = append(rog.attackpool, AttacksPool{"blow from below",9,9})
	rog.attackpool = append(rog.attackpool, AttacksPool{"backstab",25,30})
	rog.attackpool = append(rog.attackpool, AttacksPool{"rock",2,0})
	return &rog
}

func(mage *Mage) setAdditionalStats(num float32){
	mage.mana = 100
}
func(war *Warrior) setAdditionalStats(num float32){
	war.stamina = 100
}
func(rog *Rogue) setAdditionalStats(num float32){
	rog.stamina = 100
}

func(mage *Mage) GetHP() float32{
	return mage.health
}
func(war *Warrior) GetHP() float32{
	return war.health
}
func(rog *Rogue) GetHP() float32{
	return rog.health
}



func(mage *Mage) GetName() string{
	return mage.name
}
func(war *Warrior) GetName() string{
	return war.name
}
func(rog *Rogue) GetName() string{
	return rog.name
}

func(m *Mage) FindDmg(hero Hero){
	if m.mana >= m.attackpool[m.ChooseSkill()].cost {
		damage := m.attackpool[m.ChooseSkill()].dmg
		hero.setHP(damage)
		m.mana -= m.attackpool[m.ChooseSkill()].cost
		fmt.Printf(m.name + " done succ attack and dealed %.2f dmg by %s ," + hero.GetName() + " have %.2f hp\n",m.attackpool[m.ChooseSkill()].dmg,m.attackpool[m.ChooseSkill()].name,hero.GetHP())
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
		fmt.Printf(wr.name + " done succ attack and dealed %.2f dmg by %s , " + hero.GetName() +" have %.2f hp\n",wr.attackpool[wr.ChooseSkill()].dmg,wr.attackpool[wr.ChooseSkill()].name,hero.GetHP())
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
			fmt.Printf(rog.name+" done succ attack and dealed %.2f dmg by %s , "+hero.GetName()+" have %.2f hp\n", 2*rog.attackpool[rog.ChooseSkill()].dmg, rog.attackpool[rog.ChooseSkill()].name, hero.GetHP())
		} else{
			fmt.Printf(rog.name+" done succ attack and dealed %.2f dmg by %s , "+hero.GetName()+" have %.2f hp\n", rog.attackpool[rog.ChooseSkill()].dmg, rog.attackpool[rog.ChooseSkill()].name, hero.GetHP())
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

func setWarrior(wr*Warrior) *Warrior {
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
func setMage(mg*Mage) *Mage {
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
func setRogue(rog*Rogue) *Rogue {
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

func createWarrior() *Warrior {
	warrior := Warrior{}
	warrior.unitType = "Warrior"
	return setWarrior(&warrior)
}
func createMage() *Mage {
	mage := Mage{}
	mage.unitType = "Mage"
	return setMage(&mage)
}
func createRogue() *Rogue {
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

func (w *Warrior) PrintFullInfo(wr http.ResponseWriter){
	fmt.Fprintf(wr," : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",w.name,w.unitType,w.health,w.stamina,w.mana,w.defence)
}
func (m *Mage) PrintFullInfo(w http.ResponseWriter){
	fmt.Fprintf(w," : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",m.name,m.unitType,m.health,m.stamina,m.mana,m.defence)
}
func (r *Rogue) PrintFullInfo(w http.ResponseWriter){
	fmt.Fprintf(w," : %s , %s have %.2f hp, %.2f stamina,%.2f mana and %.2f def \n",r.name,r.unitType,r.health,r.stamina,r.mana,r.defence)
}

func fight(h1 Hero,h2 Hero)  int{
	h1.FindDmg(h2)
	if h2.GetHP() <= 0 {

		fmt.Println(h2.GetName() + " dead")
		return 1
	} else {
		h2.FindDmg(h1)
		if h1.GetHP() <= 0 {
			fmt.Println(h1.GetName() + " dead")
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
	return hero.GetName()
}
func PrintInfo(hero Hero){
	hero.printInfo()
}
func checkForEven(count int) bool{
	if count%2 == 0{
		return true
	} else{
		return false
	}
}


func battleTillTheEnd(i int,heroes []Hero,message chan Hero){
	personNum1 := 2*i
	personNum2 := 2*i+1
	winner := 0
	for ;winner == 0; {
		winner = fight(heroes[personNum1], heroes[personNum2])
	}
	if (winner == 1) {
		message <- heroes[personNum1]
	} else{
		message <- heroes[personNum2]
	}
}
func recovery(heroes []Hero) []Hero {
	for i:= range(heroes) {
		heroes[i].setHP(-50)
		heroes[i].setAdditionalStats(-50)
	}
	return heroes
}
func CreateRandomOne() Hero{
	num := rand.Intn(3)
	var hero Hero
	switch num{
	case 0:
		temp:= createWarrior()
		hero = temp
	case 1:
		temp := createMage()
		hero = temp
	case 2:
		temp := createRogue()
		hero = temp
	}
	return hero
}
func StartBattle(count int) Hero {
	if(!checkForEven(count)){
		print("Error! Not even count of players")
		return nil
	}
	message := make(chan Hero)
	Heroes = make([]Hero,count)
	tempheroes := make([]Hero,count)
	rand.Seed(time.Now().UnixNano())
	for i := range Heroes {
		num := rand.Intn(3)
		switch num{
		case 0:
			temp:= createWarrior()
			Heroes[i] = temp
		case 1:
			temp:= createMage()
			Heroes[i] = temp
		case 2:
			temp:= createRogue()
			Heroes[i] = temp
		}
	}
	println("Players : ")
	for i:= range(Heroes){
		PrintInfo(Heroes[i])
	}
	for ;len(Heroes)!= 0 &&!theLastOneRemained(Heroes);{
		time.Sleep(1*time.Millisecond)
		tempheroes =nil
		heroesLength := len(Heroes)
		for i:=0;i < heroesLength/2;i++ {
			 go battleTillTheEnd(i, Heroes, message)
		}
		for i:=0;i<heroesLength/2;i++ {
			tempheroes = append(tempheroes, <-message)
		}
		if(heroesLength!=2) {
			recovery(tempheroes)
		}
		Heroes = tempheroes
	}
	print(Heroes[0].GetName())
	println(" survived")
	PrintInfo(Heroes[0])
	return Heroes[0]
}

func main(){
	var count int = 1024
	StartBattle(count)
}


