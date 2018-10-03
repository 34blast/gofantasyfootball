package common

import (
	"fmt"
	"strings"
)

type Address struct {
	Line1 string
	Line2 string
	Line3 string
	City string
	State string
	ZipCode string
}

func (pAddress Address) GetLine1() string {
	return pAddress.Line1
}
func (pAddress *Address) SetLine1(pLine1 string)  {
	pAddress.Line1 = pLine1
}

func (pAddress Address) GetLine2() string {
	return pAddress.Line2
}
func (pAddress *Address) SetLine2(pLine2 string)  {
	pAddress.Line2 = pLine2
}

func (pAddress Address) GetLine3() string {
	return pAddress.Line3
}
func (pAddress *Address) SetLine3(pLine3 string)  {
	pAddress.Line3 = pLine3
}

func (pAddress Address) GetCity() string {
	return pAddress.City
}
func (pAddress *Address) SetCity(pCity string)  {
	pAddress.City = pCity
}

func (pAddress Address) GetState() string {
	return pAddress.State
}
func (pAddress *Address) SetState(pState string)  {
	pAddress.State = pState
}

func (pAddress Address) GetZipCode() string {
	return pAddress.ZipCode
}
func (pAddress *Address) SetZipCode(pZipCode string)  {
	pAddress.ZipCode = pZipCode
}

func (pAddress *Address) SetAll(pLine1 string, pLine2 string, pLine3 string, pCity string, pState string, pZipCode string) {
	pAddress.SetLine1(pLine1)
	pAddress.SetLine2(pLine2)
	pAddress.SetLine3(pLine3)
	pAddress.SetCity(pCity)
	pAddress.SetState(pState)
	pAddress.SetZipCode(pZipCode)
}

func (a Address) String() string {

	return fmt.Sprintf("[Line1:%v, Line2:%v, Line3:%v, City:%v, State:%v, ZipCode:%v]",
		a.Line1, a.Line2, a.Line3, a.City, a.State, a.ZipCode)
}

func (pAddress Address) ToString() string {
	var sb strings.Builder
	
	sb.WriteString("Values of Address object")
	sb.WriteString("\n")
	sb.WriteString("-------------------------------------------------")
	sb.WriteString("\n")
	sb.WriteString("Line1          : ")
	sb.WriteString( pAddress.GetLine1())
	sb.WriteString("\n")
	sb.WriteString("Line2          : ")
	sb.WriteString( pAddress.GetLine2())
	sb.WriteString("\n")
	sb.WriteString("Line3          : ")
	sb.WriteString(pAddress.GetLine3())
	sb.WriteString("\n")
	sb.WriteString("City           : ")
	sb.WriteString(pAddress.GetCity())
	sb.WriteString("\n")
	sb.WriteString("State          : ")
	sb.WriteString(pAddress.GetState())
	sb.WriteString("\n")
	sb.WriteString("zip code       : ")
	sb.WriteString(pAddress.GetZipCode())
	sb.WriteString("\n")

	return sb.String();
}
