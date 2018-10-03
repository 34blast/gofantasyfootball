package common

import (
	"strconv"
	"strings"
)

type unary struct {
	val string
}

const (
	WR      string = "WR"
	RB      string = "RB"
	QB      string = "QB"
	TE      string = "TE"
	DEFENSE string = "DEFENSE"
)

type Player struct {
	Position    string
	CurrentTeam string
	Ranking     int
	Person
}

func (p Player) GetPosition() string {
	return p.Position
}
func (p *Player) SetPosition(pPosition string) {
	p.Position = pPosition
}

func (p Player) GetCurentTeam() string {
	return p.CurrentTeam
}
func (p *Player) SetCurentTeam(pCurrentTeam string) {
	p.CurrentTeam = pCurrentTeam
}

func (p Player) GetRanking() int {
	return p.Ranking
}
func (p *Player) SetRanking(pRanking int) {
	p.Ranking = pRanking
}

func GetWR() string {
	return WR
}
func GetRB() string {
	return RB
}
func GetQB() string {
	return QB
}
func GetTE() string {
	return TE
}
func GetDEFENSE() string {
	return DEFENSE
}

func (p Player) String() string {

	var sb strings.Builder

	sb.WriteString("[Position:")
	sb.WriteString(p.Position)
	sb.WriteString("[ranking:")
	rankingString := strconv.Itoa(p.Ranking)
	sb.WriteString(rankingString)
	sb.WriteString("[current team:")
	sb.WriteString(p.CurrentTeam)
	sb.WriteString("] : ")

	if p.Validate() == nil {
		sb.WriteString(p.Person.String())
	} else {
		sb.WriteString(" [person null]")
	}

	return sb.String()

}

func (p *Player) SetAll(pPosition string, pCurrentTeam string, pRanking int) {
	p.Position = pPosition
	p.CurrentTeam = pCurrentTeam
	p.Ranking = pRanking
}

func CreatePlayer(pPosition string, pCurrentTeam string, pRanking int, pPerson Person) Player {
	var p Player
	p.SetAll(pPosition, pCurrentTeam, pRanking)
	p.Person = pPerson
	return p
}
