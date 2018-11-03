package fantasytest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/34blast/gofantasyfootball/common"
	"flag"
	"fmt"
	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	//	ffutil "fantasyfootball/util"
)

const (
	DB_USER        string = "admin"
	DB_PASSWORD    string = "admin"
	DB_PERSON_NAME string = "person_db"
	DB_PLAYER_NAME string = "player_db"
	DB_TYPE        string = "couch"
	DB_URL         string = "http://admin:admin@localhost:5984/"
	DB_DELETE      bool   = false
	QUOTE          string = "\""
)

var (
	client   *kivik.Client
	personDB *kivik.DB
	playerDB *kivik.DB
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain, starting execution")
	// Create the test client
	var clientError error
	client, clientError = kivik.New(DB_TYPE, DB_URL)
	if clientError != nil {
		os.Exit(1)
	}

	flag.Parse()
	exitCode := m.Run()

	// Pretend to close our DB connection
	// this part is like tearDown in Java
	cleanUp()

	// Exit
	os.Exit(exitCode)
}

func TestPersonCreateDB(t *testing.T) {
	dbExists, existsErr := client.DBExists(context.TODO(), DB_PERSON_NAME)
	assert.NoError(t, existsErr)
	if dbExists {
		t.Log("Skip db creation tests,  db already exists, db_name = ", DB_PERSON_NAME)
	} else {
		t.Log("Testing DB create, db name = ", DB_PERSON_NAME)
		_, err := client.CreateDB(context.TODO(), DB_PERSON_NAME)
		assert.NoError(t, err, "Error creating DB, db name = ", DB_PERSON_NAME)
	}
}

func TestPlayerCreateDB(t *testing.T) {
	dbExists, existsErr := client.DBExists(context.TODO(), DB_PLAYER_NAME)
	assert.NoError(t, existsErr)
	if dbExists {
		t.Log("Skip db creation tests,  db already exists, db_name = ", DB_PLAYER_NAME)
	} else {
		t.Log("Testing DB create, db name = ", DB_PLAYER_NAME)
		_, err := client.CreateDB(context.TODO(), DB_PLAYER_NAME)
		assert.NoError(t, err, "Error creating DB, db name = ,", DB_PLAYER_NAME)
	}
}

func TestPersonConnection(t *testing.T) {
	returnDSN := client.DSN()
	assert.Equal(t, DB_URL, returnDSN, "Error DSN does not agree with initial value")

	// connect to example DB
	var dbConnectError error
	personDB, dbConnectError = client.DB(context.TODO(), DB_PERSON_NAME)
	assert.NoError(t, dbConnectError)
}

func TestPlayerConnection(t *testing.T) {
	returnDSN := client.DSN()
	assert.Equal(t, DB_URL, returnDSN, "Error DSN does not agree with initial value")

	// connect to example DB
	var dbConnectError error
	playerDB, dbConnectError = client.DB(context.TODO(), DB_PLAYER_NAME)
	assert.NoError(t, dbConnectError)
}

func TestCreatePersonDesignDoc(t *testing.T) {
	filePath := "person-views.json"
	viewFile, fileError := ioutil.ReadFile(filePath)
	assert.NoError(t, fileError)
	// use self make function to http POST the new design doc and views
	err := createDesignDoc(viewFile, DB_URL, DB_PERSON_NAME)
	assert.NoError(t, err)
}

func TestCreatePlayerDesignDoc(t *testing.T) {
	filePath := "player-views.json"
	viewFile, fileError := ioutil.ReadFile(filePath)
	assert.NoError(t, fileError)
	// use self make function to http POST the new design doc and views
	err := createDesignDoc(viewFile, DB_URL, DB_PLAYER_NAME)
	assert.NoError(t, err)
}

func TestInsertOwnerRich(t *testing.T) {
	rich := getRich()
	exists, idString := personExists(rich.GetLastName(), rich.GetFirstName())
	if exists {
		rich.SetID(idString)
		t.Log("Skipping creating Rich,  already exists")
	} else {
		richRev, createErr := personDB.Put(context.TODO(), rich.GetID(), rich)
		assert.NoError(t, createErr)
		rich.SetRev(richRev)
	}

}

func TestInsertOwnerTracy(t *testing.T) {
	tracy := getTracy()
	exists, idString := personExists(tracy.GetLastName(), tracy.GetFirstName())
	if exists {
		tracy.SetID(idString)
		t.Log("Skipping creating Tracy,  already exists")
	} else {
		tracyRev, createErr := personDB.Put(context.TODO(), tracy.GetID(), tracy)
		assert.NoError(t, createErr)
		tracy.SetRev(tracyRev)
	}

}

func TestInsertCommisionerBrandon(t *testing.T) {
	brandon := getBrandon()
	exists, idString := personExists(brandon.GetLastName(), brandon.GetFirstName())
	if exists {
		brandon.SetID(idString)
		t.Log("Skipping creating Brandon,  already exists")
	} else {
		brandonRev, createErr := personDB.Put(context.TODO(), brandon.GetID(), brandon)
		assert.NoError(t, createErr)
		brandon.SetRev(brandonRev)
	}
}

func TestInsertPlayers(t *testing.T) {
	players := getPlayers()

	for _, player := range players {
		exists, idString := playerExists(player.GetLastName(), player.GetFirstName())
		if exists {
			player.SetID(idString)
			t.Log("Skipping creating player, already exists, player is ", player.GetFirstName(), " ", player.GetLastName())
		} else {
			rev, err := playerDB.Put(context.TODO(), player.GetID(), player)
			assert.NoError(t, err)
			player.SetRev(rev)
		}

	}
}

func TestUpdateTracy(t *testing.T) {
	tracy := getPersonByName("Scott", "Tracy")
	tracy.SetOtherPhone("972 555-1212")
	tracy.SetEmail("tracy@tracy.com")
	updateRev, updateError := personDB.Put(context.TODO(), tracy.GetID(), tracy)
	assert.NoError(t, updateError)
	tracy.SetRev(updateRev)
}

func TestQueryPersonrByFullName(t *testing.T) {
	position := "ScottBrandon"
	value, _ := json.Marshal(position)
	options := map[string]interface{}{
		"key": string(value),
	}
	rows, queryError := personDB.Query(context.TODO(), "_design/personDoc", "_view/fullname-view", options)
	assert.NoError(t, queryError)
	count := 0
	for rows.Next() {
		count++
		var person common.Person
		scanError := rows.ScanValue(&person)
		assert.NoError(t, scanError)
		assert.Equal(t, "Scott", person.GetLastName())
		assert.Equal(t, "Brandon", person.GetFirstName())
	}
	assert.Equal(t, 1, count, "Count should be 1")
}

func TestQueryPlayerByTEPosition(t *testing.T) {
	position := "TE"
	value, _ := json.Marshal(position)
	options := map[string]interface{}{
		"key": string(value),
	}
	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "_view/position-view", options)
	assert.NoError(t, queryError)
	for rows.Next() {
		var player common.Player
		scanError := rows.ScanValue(&player)
		assert.NoError(t, scanError)
		assert.Equal(t, common.TE, player.GetPosition())
	}
}

func TestSortPlayersByRanking(t *testing.T) {
	options := map[string]interface{}{
		"startKey":   "9999.9999",
		"endKey":     "-111.111",
		"descending": "true",
	}
	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "_view/ranking-view", options)
	assert.NoError(t, queryError)
	count := 0
	prevRank := 9999
	for rows.Next() {
		count++
		var player common.Player
		scanError := rows.ScanValue(&player)
		assert.NoError(t, scanError)
		currentRank := player.GetRanking()
		if prevRank < currentRank {
			errorMessage := fmt.Sprintf("previous ranks %d,  curent rank %d\n", prevRank, player.GetRanking())
			t.Error(errorMessage)
		}
		prevRank = player.GetRanking()
	}
}

// This one get them all by ranking, remove the unneeded positions then check
func TestSortWRPlayersByRankingManual(t *testing.T) {
	options := map[string]interface{}{
		"startKey":   "9999.9999",
		"endKey":     "-111.111",
		"descending": "true",
	}
	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "ranking-view", options)
	assert.NoError(t, queryError)
	count := 0
	prevRank := 9999
	var playerArray []common.Player

	for rows.Next() {
		count++
		var player common.Player
		scanError := rows.ScanValue(&player)
		assert.NoError(t, scanError)
		currentRank := player.GetRanking()
		if player.GetPosition() == "WR" {
			playerArray = append(playerArray, player)
		}
		if prevRank < currentRank {
			errorMessage := fmt.Sprintf("previous ranks %d,  curent rank %d\n", prevRank, player.GetRanking())
			t.Error(errorMessage)
		}
		prevRank = player.GetRanking()
	}
	for _, player := range playerArray {
		assert.Equal(t, "WR", player.GetPosition(), "ERROR: position is not WR, position = ", player.GetPosition())
	}
}

/**
This Tests out getting the higest ranking wide receivers.  To support this a view was created with 2 values in the key
Ranking and Position.  The function looks like this
		function(doc, meta)
		 {
		   if (doc.Ranking && doc.Position)
		  {
		     emit([doc.Position, doc.Ranking],doc);
		  }
		 }

once the view is created you can test on the command line like below
	--- curl command -------
		curl -g -X GET 'http://localhost:5984/player_db/_design/playerDoc/_view/positionandrank-view?startkey=["WR",9999]&endkey=["WR",-9999]&descending=true'

The key is the parameter passed to the view positionandrankd-view below
		tartkey=["WR",9999]&endkey=["WR",-9999]&descending=true

Kivik gives alot of trouble when passing options.  It took a while, but I realized it does not support
startKey and endKey as the curl command does,  so instead use start_key and end_key accordingly
Also it has some trouble with understanding what is to be json encoded and what is not. You can use various
methods of encoding to get the desired rsult

	// Manual way to make the required syntax to Kivik
	options := map[string]interface{}{"start_key": `["WR",9999]`, "end_key": `["WR",-1]`,"descending": "true"}
**/
func TestSortWRPlayersByRankingByView(t *testing.T) {

	startKeyString := "[\"WR\",999]"
	endKeyString := "[\"WR\",-1]"
	options := map[string]interface{}{
		"start_key":  startKeyString,
		"end_key":    endKeyString,
		"descending": "true",
	}

	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "positionandrank-view", options)
	assert.NoError(t, queryError)
	count := 0
	prevRank := 9999
	for rows.Next() {
		count++
		var player common.Player
		scanError := rows.ScanValue(&player)
		assert.NoError(t, scanError)
		currentRank := player.GetRanking()
		assert.Equal(t, "WR", player.GetPosition(), "ERROR: position is not WR, position = ", player.GetPosition())
		if prevRank < currentRank {
			errorMessage := fmt.Sprintf("previous ranks %d,  curent rank %d\n", prevRank, player.GetRanking())
			t.Error(errorMessage)
		}
		prevRank = player.GetRanking()
	}
	if count < 1 {
		t.Error("No Players Found, should be more than 1")
	}
}

func TestDeletePersonDB(t *testing.T) {
	if DB_DELETE {
		t.Log("Testing DB delete")
		err := client.DestroyDB(context.TODO(), DB_PERSON_NAME)
		assert.NoError(t, err, "Error deleting DB, db name = ", DB_PERSON_NAME)
	} else {
		t.Log("Testing DB delete, skipping for db name = ", DB_PERSON_NAME)
	}
}

func TestDeletePlayerDB(t *testing.T) {
	if DB_DELETE {
		t.Log("Testing DB delete")
		err := client.DestroyDB(context.TODO(), DB_PLAYER_NAME)
		assert.NoError(t, err, "Error deleting DB, db name = ", DB_PLAYER_NAME)
	} else {
		t.Log("Testing DB delete, skipping for db name = ", DB_PLAYER_NAME)
	}
}
func getRich() common.Owner {
	var rich common.Owner
	richId := xid.New()
	richIdString := richId.String()
	rich.SetAll("Scott", "Richard", "rmscott@rmscott.com", "214 555-1212", "", 52, richIdString, "")
	rich.SetTeamName("Rich's rejects")
	rich.SetID(richIdString)

	return rich
}

func getTracy() common.Owner {
	var tracy common.Owner
	tracyId := xid.New()
	tracyIdString := tracyId.String()
	tracy.SetAll("Scott", "Tracy", "tdscott@rmscott.com", "214 555-1212", "", 53, tracyIdString, "")
	tracy.SetTeamName("Blitzburg Squeelers")
	tracy.SetID(tracyIdString)

	return tracy
}

func getBrandon() common.Commissioner {
	var brandon common.Commissioner
	brandonId := xid.New()
	brandonIdString := brandonId.String()
	brandon.Address.SetAll("1221 Big Falls Drive", "", "", "Flower Mound", "Texas", "75028")
	brandon.Person.SetAll("Scott", "Brandon", "bmscott@bmscott.com", "214 555-1212", "", 22, brandonIdString, "")
	brandon.SetTitle("The commissioner")
	brandon.SetID(brandonIdString)

	return brandon
}

func getPlayers() []common.Player {
	var playerArray []common.Player

	var ajGreen common.Player
	ajGreen.SetAll(common.WR, "Bengals", 92)
	ajGreen.SetLastName("Green")
	ajGreen.SetFirstName("AJ")
	ajID := xid.New()
	ajIDString := ajID.String()
	ajGreen.SetID(ajIDString)
	playerArray = append(playerArray, ajGreen)

	var julioJones common.Player
	julioJones.SetAll(common.WR, "Falcons", 92)
	julioJones.SetLastName("Jones")
	julioJones.SetFirstName("Julio")
	jjID := xid.New()
	jjIDString := jjID.String()
	julioJones.SetID(jjIDString)
	playerArray = append(playerArray, julioJones)

	var mSanu common.Player
	mSanu.SetAll(common.WR, "Falcons", 63)
	mSanu.SetLastName("Sanu")
	mSanu.SetFirstName("Mohammed")
	msID := xid.New()
	msIDString := msID.String()
	mSanu.SetID(msIDString)
	playerArray = append(playerArray, mSanu)

	var jRoss common.Player
	jRoss.SetAll(common.WR, "Bengals", 33)
	jRoss.SetLastName("Ross")
	jRoss.SetFirstName("John")
	jrID := xid.New()
	jrIDString := jrID.String()
	jRoss.SetID(jrIDString)
	playerArray = append(playerArray, jRoss)

	var tKroft common.Player
	tKroft.SetAll(common.TE, "Bengals", 43)
	tKroft.SetLastName("Kroft")
	tKroft.SetFirstName("Tyler")
	tkID := xid.New()
	tkIDString := tkID.String()
	tKroft.SetID(tkIDString)
	playerArray = append(playerArray, tKroft)

	return playerArray
}

func cleanUp() {
	playerDB = nil
	personDB = nil
	client = nil
}

// This document sends an HTTP POST with the contained JSON data to create a design document with views
// the DB URL is expected to contain the user id and the password as well as the database
// e.g. http://user:pass@hostname:5984/example_db
func createDesignDoc(pJSONData []byte, pDBURL string, pDBName string) error {

	isSlashLast := strings.HasSuffix(pDBURL, "/")

	var sb strings.Builder
	sb.WriteString(pDBURL)
	if !isSlashLast {
		sb.WriteString("/")
	}
	sb.WriteString(pDBName)
	fullURL := sb.String()

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fullURL,
		bytes.NewReader(pJSONData),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	fmt.Println("response status from Do = ", resp.Status)
	return err
}

func personExists(pLastName string, pFirstName string) (bool, string) {
	flag := false
	idString := "id"

	fullName := pLastName + pFirstName
	value, _ := json.Marshal(fullName)
	options := map[string]interface{}{
		"key": string(value),
	}

	rows, queryError := personDB.Query(context.TODO(), "_design/personDoc", "_view/fullname-view", options)
	if queryError != nil {
		panic(queryError)
	}
	count := 0
	for rows.Next() {
		count++
		idString = rows.ID()

		flag = true
		if count > 1 {
			fmt.Printf("WARNING: there are more than one row for last name = %s first name = %s row total %d\n", pLastName, pFirstName, count)
			break
		}
	}
	return flag, idString

} // end of personExists()

func playerExists(pLastName string, pFirstName string) (bool, string) {
	flag := false
	idString := "id"

	fullName := pLastName + pFirstName
	value, _ := json.Marshal(fullName)
	options := map[string]interface{}{
		"key": string(value),
	}

	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "_view/fullname-view", options)
	if queryError != nil {
		panic(queryError)
	}
	count := 0
	for rows.Next() {
		count++
		idString = rows.ID()

		flag = true
		if count > 1 {
			fmt.Printf("WARNING: there are more than one row for last name = %s first name = %s row total %d\n", pLastName, pFirstName, count)
			break
		}
	}
	return flag, idString

} // end of playerExists()

func getPersonById(pID string) common.Person {
	var person common.Person
	row := personDB.Get(context.TODO(), pID)
	fetchError := row.Err
	if fetchError != nil {
		panic(fetchError)
	} else {
		err := row.ScanDoc(&person)
		if err != nil {
			panic(err)
		}
	}
	person.SetRev(row.Rev)
	return person
}

func getPlayerById(pID string) common.Player {
	var player common.Player
	row := playerDB.Get(context.TODO(), pID)
	fetchError := row.Err
	if fetchError != nil {
		panic(fetchError)
	} else {
		err := row.ScanDoc(&player)
		if err != nil {
			panic(err)
		}
	}
	player.SetRev(row.Rev)

	return player
}

func getPersonByName(pLastName string, pFirstName string) common.Person {

	var person common.Person

	fullName := pLastName + pFirstName
	value, _ := json.Marshal(fullName)
	options := map[string]interface{}{
		"key": string(value),
	}

	rows, queryError := personDB.Query(context.TODO(), "_design/personDoc", "_view/fullname-view", options)
	if queryError != nil {
		panic(queryError)
	}

	for rows.Next() {
		rows.ScanValue(&person)
		break
	}

	return person

} // end of getPersonByName()

func getPlayerByName(pLastName string, pFirstName string) common.Player {

	var player common.Player

	fullName := pLastName + pFirstName
	value, _ := json.Marshal(fullName)
	options := map[string]interface{}{
		"key": string(value),
	}

	rows, queryError := playerDB.Query(context.TODO(), "_design/playerDoc", "_view/fullname-view", options)
	if queryError != nil {
		panic(queryError)
	}

	for rows.Next() {
		rows.ScanValue(&player)
		break
	}

	return player

} // end of getPlayerByName()
