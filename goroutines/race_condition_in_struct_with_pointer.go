package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	person := Person{Name: "Mert", Age: 24, Skills: []string{}}
	fmt.Println(time.Now())

	wg := sync.WaitGroup{}
	wg.Add(2)
	lock := sync.Mutex{}
	// setFirstSkillSet(&person, &wg, &lock)
	// setSecondSkillSet(&person, &wg, &lock)
	// Above two methods takes 10 second

	go setFirstSkillSet(&person, &wg, &lock)
	go setSecondSkillSet(&person, &wg, &lock)
	// Above two methods takes 5 second
	wg.Wait() // blocking

	fmt.Println(person) // print {Mert 24 [Java Javascript GoLang Docker]} or {Mert 24 [GoLang Docker Java Javascript]}
	fmt.Println(time.Now())
}

func setFirstSkillSet(person *Person, wg *sync.WaitGroup, lock *sync.Mutex) {
	time.Sleep(5 * time.Second)
	firstSkills := []string{"Java", "Javascript"}
	for _, skill := range firstSkills {
		lock.Lock()
		person.Skills = append(person.Skills, skill)
		lock.Unlock()
	}
	wg.Done()
}

func setSecondSkillSet(person *Person, wg *sync.WaitGroup, lock *sync.Mutex) {
	time.Sleep(5 * time.Second)
	firstSkills := []string{"GoLang", "Docker"}
	for _, skill := range firstSkills {
		lock.Lock()
		person.Skills = append(person.Skills, skill)
		lock.Unlock()
	}
	wg.Done()
}

type Person struct {
	Name   string
	Age    int
	Skills []string
}
