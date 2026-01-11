package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is interface pattern learning code")
	personOne := NewPerson("Kalu", 23, "Bhimdatta", true)
	isSuccess := ParticipateInElection(personOne)

	if isSuccess {
		fmt.Println("Vote cast successfully")
	} else {
		fmt.Println("Vote casting failed")
	}

	newCat := Cat{Name: "pussi"}

	isSuccess = ParticipateInElection(newCat)
}

type Person struct {
	Name                 string
	Age                  int
	Location             string
	InterestedInPolitics bool
}

type Cat struct {
	Name string
}

func (c Cat) CastVote() bool {
	return true
}

func (c Cat) RejectElection() bool {
	return false
}

func (c Cat) IsEligible() bool {
	return true
}

func (c Cat) IsInterested() bool {
	return true
}

type Voter interface {
	CastVote() bool
	RejectElection() bool
	IsEligible() bool
	IsInterested() bool
}

func NewCat(name string) *Cat {
	return &Cat{
		Name: "pussi",
	}
}

func NewPerson(name string, age int, location string, interestedInPolitics bool) *Person {
	return &Person{
		Name:                 name,
		Age:                  age,
		Location:             location,
		InterestedInPolitics: interestedInPolitics,
	}
}

// Interface implementation
func (p *Person) CastVote() bool {
	fmt.Printf("%s has voted successfully at %s\n", p.Name, p.Location)
	return true
}

func (p *Person) RejectElection() bool {
	return !p.IsInterested()
}

func (p *Person) IsEligible() bool {
	fmt.Printf("%s you are Elegible for Vote your age is:%v\n", p.Name, p.Age)
	return p.Age >= 18
}

func (p *Person) IsInterested() bool {
	return p.InterestedInPolitics
}

// Business logic
func ParticipateInElection(voter Voter) bool {
	if !voter.IsEligible() {
		fmt.Println("Voter is not eligible")
		return false
	}

	if !voter.IsInterested() {
		voter.RejectElection()
		return false
	}

	return voter.CastVote()
}
