package common

import (
	"errors"
	"fmt"
	"strings"
)

const MIN_NAME int = 2
const LAST_NAME_ERROR string = "ERROR: last name must be greater than 2 or greater in length"
const FIRST_NAME_ERROR string = "ERROR: first name must be greater than 2 or greater in length"

type PersonInterface interface {
	GetID() string
	SetID(pID string)
	GetRev() string
	SetRev(pRev string)
	GetLastName() string
	SetLastName(pLastName string)
	GetFirstName() string
	SetFirstName(pFirstName string)
	GetEmail() string
	SetEmail(pEmail string)
	GetCellPhone() string
	SetCellPhone(pCellPhone string)
	GetOtherPhone() string
	SetOtherPhone(pOtherPhone string)

	SetAll(pLastName string, pFirstName string, pEmail string, pCellPhone string, pOtherPhone string, pAge int, pID string, pRev string) error
	Validate() error

	ToString() string
	PrintVerbose()

	String() string
}

type Person struct {
	ID         string `json:"_id"`
	Rev        string `json:"_rev,omitempty"`
	LastName   string
	FirstName  string
	Email      string
	CellPhone  string
	OtherPhone string
	Age        int
}

func (p Person) GetID() string {
	return p.ID
}
func (p *Person) SetID(pID string) {
	p.ID = pID
}
func (p Person) GetRev() string {
	return p.Rev
}
func (p *Person) SetRev(pRev string) {
	p.Rev = pRev
}
func (p Person) GetLastName() string {
	return p.LastName
}
func (p *Person) SetLastName(pLastName string) {
	p.LastName = pLastName
}
func (p Person) GetFirstName() string {
	return p.FirstName
}
func (p *Person) SetFirstName(pFirstName string) {
	p.FirstName = pFirstName
}

func (p Person) GetEmail() string {
	return p.Email
}
func (p *Person) SetEmail(pEmail string) {
	p.Email = pEmail
}

func (p Person) GetCellPhone() string {
	return p.CellPhone
}
func (p *Person) SetCellPhone(pCellPhone string) {
	p.CellPhone = pCellPhone
}

func (p Person) GetOtherPhone() string {
	return p.OtherPhone
}
func (p *Person) SetOtherPhone(pOtherPhone string) {
	p.OtherPhone = pOtherPhone
}
func (p Person) GetAge() int {
	return p.Age
}
func (p *Person) SetAge(pAge int) {
	p.Age = pAge
}

func (p Person) ToString() string {
	var sb strings.Builder

	sb.WriteString("Values of Person object")
	sb.WriteString("\n")
	sb.WriteString("-------------------------------------------------")
	sb.WriteString("\n")
	sb.WriteString("ID             : ")
	sb.WriteString(p.GetID())
	sb.WriteString("Rev            : ")
	sb.WriteString(p.GetRev())
	sb.WriteString("first name     : ")
	sb.WriteString(p.GetFirstName())
	sb.WriteString("\n")
	sb.WriteString("last name      : ")
	sb.WriteString(p.GetLastName())
	sb.WriteString("\n")
	sb.WriteString("email          : ")
	sb.WriteString(p.GetEmail())
	sb.WriteString("\n")
	sb.WriteString("cell phone     : ")
	sb.WriteString(p.GetCellPhone())
	sb.WriteString("\n")
	sb.WriteString("other phone    : ")
	sb.WriteString(p.GetOtherPhone())
	sb.WriteString("Age            : ")
	sb.WriteString(string(p.GetAge()))
	sb.WriteString("\n")

	return sb.String()
}

func (p Person) String() string {

	return fmt.Sprintf("[id:%v, rev:%v, first name:%v, last name:%v, email:%v, cellPhone:%v, other phone:%v, age:%v]",
		p.ID, p.Rev, p.FirstName, p.LastName, p.Email, p.CellPhone, p.OtherPhone, p.Age)
}

func (p Person) PrintVerbose() {

	fmt.Println("Values of Person object")
	fmt.Println("-------------------------------------------------")
	fmt.Println("ID             : ", p.GetID())
	fmt.Println("Rev            : ", p.GetRev())
	fmt.Println("first name     : ", p.GetFirstName())
	fmt.Println("last name      : ", p.GetLastName())
	fmt.Println("email          : ", p.GetEmail())
	fmt.Println("cell phone     : ", p.GetCellPhone())
	fmt.Println("other phone    : ", p.GetOtherPhone())
	fmt.Println("age            : ", p.GetAge())
}

func (p *Person) SetAll(pLastName string, pFirstName string, pEmail string, pCellPhone string, pOtherPhone string, pAge int, pID string, pRev string) error {
	p.LastName = pLastName
	p.FirstName = pFirstName
	p.Email = pEmail
	p.CellPhone = pCellPhone
	p.OtherPhone = pOtherPhone
	p.Age = pAge
	p.ID = pID
	p.Rev = pRev

	return p.Validate()
}

func (p Person) Validate() error {
	if len(p.LastName) < MIN_NAME {
		return errors.New(LAST_NAME_ERROR)
	} else if len(p.FirstName) < MIN_NAME {
		return errors.New(FIRST_NAME_ERROR)
	} else {
		return nil
	}
}
