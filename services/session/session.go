package session

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

type sessionData struct{
	Login string
	Points string
}

type Session struct{
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)
	s.data = make(map[string]*sessionData)
	return s
}
func GenerateId() string{
	b :=make([]byte,16)
	rand.Read(b)
	return fmt.Sprintf("%x",b)
}
func(s *Session) Init(login string,points int)string{
	sessionId := GenerateId()

	data:= &sessionData{Login: login,Points: strconv.Itoa(points)}
	s.data[sessionId] =data

	return sessionId
}

func(s*Session) Get(sessionId string) (string,int){
	data := s.data[sessionId]
	if(data == nil){
		return "",0
	} else{
		points,_ := strconv.Atoi(data.Points)
		return data.Login, points
	}
}