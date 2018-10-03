package common

import (
	"strings"
)

type Owner struct {
	TeamName string
	Person
}

func (pOwner Owner) GetTeamName() string {
	return pOwner.TeamName
}

func (pOwner *Owner) SetTeamName(pTeamName string) {
	pOwner.TeamName = pTeamName
}

func (pOwner Owner) ToString() string {

	var sb strings.Builder

	sb.WriteString("Values of Owner object")
	sb.WriteString("\n")
	sb.WriteString("-------------------------------------------------")
	sb.WriteString("\n")
	sb.WriteString("title          :")
	sb.WriteString(pOwner.TeamName)
	sb.WriteString("\n")
	sb.WriteString(pOwner.Person.ToString())
	sb.WriteString("\n")

	return sb.String()

}

func (pOwner Owner) String() string {

	var sb strings.Builder

	sb.WriteString("[title:")
	sb.WriteString(pOwner.TeamName)
	sb.WriteString("] : ")

	if pOwner.Validate() == nil {
		sb.WriteString(pOwner.Person.String())
	} else {
		sb.WriteString(" [person null]")
	}

	return sb.String()

}
