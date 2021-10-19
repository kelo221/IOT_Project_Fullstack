package main

import (
	"context"
	_ "context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var db driver.Database
var tempCol driver.Collection
var userCol driver.Collection
var loginCol driver.Collection

func handleDatabase() {

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		fmt.Println(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "1234"),
	})
	if err != nil {
		fmt.Println(err)
	}

	// Create a database
	db, err = client.Database(nil, "IOT_DATA")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		db, err = client.CreateDatabase(ctx, "IOT_DATA", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create a collection for data
	tempCol, err = db.Collection(nil, "IOT_DATA_SENSOR")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		tempCol, err = db.CreateCollection(ctx, "IOT_DATA_SENSOR", options)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create a collection for users
	userCol, err = db.Collection(nil, "IOT_DATA_LOGIN")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		userCol, err = db.CreateCollection(ctx, "IOT_DATA_LOGIN", options)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create a collection for login dates on such
	loginCol, err = db.Collection(nil, "IOT_DATA_LOGS")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		loginCol, err = db.CreateCollection(ctx, "IOT_DATA_LOGS", options)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func appendToDB(tpackage dataPackageIn) {

	ctx := context.Background()
	meta, err := tempCol.CreateDocument(ctx, tpackage)
	if err != nil {
		// handle error
	}
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)

}

func dropDatabase() {

	fmt.Println("Database dropped")

	ctx := context.Background()
	query := "FOR u IN IOT_DATA_SENSOR REMOVE u IN IOT_DATA_SENSOR"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)

}

func createAccounts() {
	salt := []byte("salt")

	aqlNoReturn("UPSERT { username: 'v' } " +
		"INSERT { username: 'v', hash: '" + hex.EncodeToString(HashPassword([]byte("v"), salt)) + "', dateCreated: DATE_NOW() } " +
		"UPDATE {} IN IOT_DATA_LOGIN")

	aqlNoReturn("UPSERT { username: 'x' } " +
		"INSERT { username: 'x', hash: '" + hex.EncodeToString(HashPassword([]byte("x"), salt)) + "', dateCreated: DATE_NOW() } " +
		"UPDATE {} IN IOT_DATA_LOGIN")

	aqlNoReturn("UPSERT { username: 'teach' } " +
		"INSERT { username: 'teach', hash: '" + hex.EncodeToString(HashPassword([]byte("teach"), salt)) + "', dateCreated: DATE_NOW() } " +
		"UPDATE {} IN IOT_DATA_LOGIN")

}

type logData struct {
	User string `json:"user,omitempty"`
	Time int    `json:"time,omitempty"`
}

func getDBLogs() []logData {
	query := "FOR x IN IOT_DATA_LOGS RETURN x"

	var dataPayload []logData

	ctx := context.Background()
	//query = "FOR Speed IN IOT_DATA_SENSOR RETURN Speed"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc logData
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta.Rev)
		//fmt.Println(doc)
		dataPayload = append(dataPayload, doc)
	}

	return dataPayload

}

func aqlMQTT(query string) []dataPackageIn {

	var dataPayload []dataPackageIn

	ctx := context.Background()
	//query = "FOR Speed IN IOT_DATA_SENSOR RETURN Speed"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc dataPackageIn
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta.Rev)
		//fmt.Println(doc)
		dataPayload = append(dataPayload, doc)
	}

	return dataPayload
}

func aqlNoReturn(query string) {

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)

}

func aqlToString(query string) string {

	var result string

	ctx := context.Background()
	//query = "FOR Speed IN IOT_DATA_SENSOR RETURN Speed"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		_, err2 := cursor.ReadDocument(ctx, &result)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}

		//fmt.Println(result)
	}

	return result
}
