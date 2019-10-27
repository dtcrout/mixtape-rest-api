package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Track struct {
	Artist string
	Song   string
}

func (t Track) String() string {
	return fmt.Sprintf("Artist: %s, Song: %s", t.Artist, t.Song)
}

func ReadCSV(filename string) [][]string {
	// Open CSV file
	f, _ := os.Open(filename)

	// Close file
	defer f.Close()

	// Create reader
	reader := csv.NewReader(f)

	// Read all records
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}

// Hello world handler function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	lines := ReadCSV("resources/tracks_2.csv")

	// Create slice to hold tracks
	var tracks []Track

	for i, line := range lines {
		// Skip header
		if i == 0 {
			continue
		}
		track := Track{Artist: line[0], Song: line[1]}
		tracks = append(tracks, track)
	}

	for _, track := range tracks {
		fmt.Println(track)
	}
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
