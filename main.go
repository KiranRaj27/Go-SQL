package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	album "github.com/kiranraj27/gosql/handler"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "sample",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := album.AlbumsByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := album.AlbumByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := album.AddAlbum(db, album.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}
