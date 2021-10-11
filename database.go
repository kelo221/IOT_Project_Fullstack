package main

import (
	"context"
	_ "context"
	"crypto/tls"
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var col driver.Collection

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
	db, err := client.Database(nil, "IOT_DATA")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		db, err = client.CreateDatabase(ctx, "IOT_DATA", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create a collection
	col, err = db.Collection(nil, "IOT_DATA_SENSOR")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		col, err = db.CreateCollection(ctx, "IOT_DATA_SENSOR", options)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func appendToDB(tpackage dataPackage) {

	ctx := context.Background()
	meta, err := col.CreateDocument(ctx, tpackage)
	if err != nil {
		// handle error
	}
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)

}

func dropDatabase() {

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
	db, err := client.Database(nil, "IOT_DATA")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		db, err = client.CreateDatabase(ctx, "IOT_DATA", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create a collection
	col, err = db.Collection(nil, "IOT_DATA_SENSOR")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		col, err = db.CreateCollection(ctx, "IOT_DATA_SENSOR", options)
		if err != nil {
			fmt.Println(err)
		}
	}

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

func aql(query string) []dataPackage {

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
	db, err := client.Database(nil, "IOT_DATA")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		db, err = client.CreateDatabase(ctx, "IOT_DATA", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	var dataPayload []dataPackage

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
		var doc dataPackage
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
