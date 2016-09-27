package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/jackc/pgx"
)

// don't have task data definition -- shouldn't matter since not used
type taskData string

type projectData struct {
	ID         string     `json:"id" dt:"-"`
	CratedAt   time.Time  `json:"created_at" db:"-"`
	Name       string     `json:"name"`
	Name2      string     `json:"name2"`
	Site       string     `json:"site"`
	Device     int        `json:"device"`
	BrowserBar int        `json:"browser_bar"`
	Type       int        `json:"type"`
	Email      string     `json:"email"`
	Status     int        `json:"status"`
	Tasks      []taskData `json:"tasks"`
}

func RandString(size int) string {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		panic("couldn't get rand data")
	}

	return hex.EncodeToString(buf)[:size]
}

func main() {
	config, err := pgx.ParseEnvLibpq()
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Initialing draft project
	draftProject := projectData{
		Name: "",
		Site: "",
		//TrackSesstion: track/session?uid={id}
		Device: 1,  // Default: Desktop
		Type:   1,  // TODO default or from request
		Email:  "", // TODO default or from request
		Status: 1,
		Tasks:  []taskData{},
	}

	query := `
    INSERT INTO projects (id, user_id, data)
    VALUES ($1, $2, $3::jsonb )
    `

	hash := RandString(10) // Length: 10
	uid := RandString(10)

	_, err = conn.Exec(query, hash, uid, &draftProject)
	if err != nil {
		log.Fatal(err)
	}
}
