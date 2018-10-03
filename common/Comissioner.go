package common

import (
	"strings"
)

type Commissioner struct {
	Title string
	Person
	Address
}

func (c Commissioner) GetTitle() string {
	return c.Title;
}

func (c *Commissioner) SetTitle(pTitle string)  {
	c.Title = pTitle;
}

func (c Commissioner) String() string {
	var sb strings.Builder
	sb.WriteString("[Title:")
	sb.WriteString(c.Title)
	sb.WriteString("] : ")
	sb.WriteString(c.Person.String())
	sb.WriteString(c.Address.String())
	
	return sb.String()
}

func (pCommissioner Commissioner) ToString() string {

	var sb strings.Builder

	sb.WriteString("Values of Commissioner object")
	sb.WriteString("\n")
	sb.WriteString("-------------------------------------------------")
	sb.WriteString("\n")
	sb.WriteString("Title          :")
	sb.WriteString(pCommissioner.Title)
	sb.WriteString("\n")
	sb.WriteString(pCommissioner.Person.ToString())
	sb.WriteString(pCommissioner.Address.ToString())
	sb.WriteString("\n")

	return sb.String()

}