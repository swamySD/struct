package main

import "fmt"
type bot interface{
	getGreeting() string
}

type speech interface{
	getSpeaker () string
}
type englishBot struct{}
type spanishBot struct{}

type englishSpeaker struct{}
type teluguSpeaker struct{}
func main() {
 eb:=englishBot{}
 sb:=spanishBot{} 
 es:=englishSpeaker{}
 ts:=teluguSpeaker{}
 printGreeting(eb)
 printGreeting(sb)
 speaker(es)
 speaker(ts)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func speaker(s speech){
  fmt.Println(s.getSpeaker())
}

func (englishSpeaker) getSpeaker() string{
 return "Hi iam english Speaker"
}

func (teluguSpeaker) getSpeaker() string{
	return "Hi iam Telugu Speaker"
   }
func (englishBot) getGreeting() string {
	//very custom logic for generating an english  greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating an spanish  greeting
	return "Hola!"
}