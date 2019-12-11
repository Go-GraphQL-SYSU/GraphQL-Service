package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"net/http"
	"sync"
)

const (
	dbPath          = "./data/data.db"
	dbMode          = 0600
	PeopleBucket    = "PeopleBucket"
	PlanetsBucket   = "PlanetsBucket"
	StarshipsBucket = "StarshipsBucket"
	FilmsBucket     = "FilmsBucket"
	SpeciesBucket   = "SpeciesBucket"
	VehiclesBucket  = "VehiclesBucket"
)

type Crawler struct {
	client *http.Client
}

// Get a crawler using the default client
func GetCrawlerSync() *Crawler {
	return &Crawler{
		client: http.DefaultClient,
	}
}

// Get a crawler using a custom client
func GetCrawlerAsync(c *http.Client) *Crawler {
	return &Crawler{
		client: c,
	}
}

// Use coroutines to crawl different information to improve performance
func ScrapAndSave() {
	// Open the db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbPath, dbMode, nil)
	if err != nil {
		log.Fatalf("Error when open bolt database:%s\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error when close the database:%s\n", err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(6)

	go func() {
		people, err := GetCrawlerAsync(&http.Client{}).getPeopleInfo()
		if err != nil {
			log.Printf("Error when crawl people info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range people {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (people):%s\n", err)
				continue
			}

			if err := writeToDB(db, PeopleBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (people):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()

	go func() {
		planets, err := GetCrawlerAsync(&http.Client{}).getPlanetsInfo()
		if err != nil {
			log.Printf("Error when crawl planets info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range planets {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (planets):%s\n", err)
				continue
			}

			if err := writeToDB(db, PlanetsBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (planets):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()

	go func() {
		starships, err := GetCrawlerAsync(&http.Client{}).getStarshipsInfo()
		if err != nil {
			log.Printf("Error when crawl starships info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range starships {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (starships):%s\n", err)
				continue
			}

			if err := writeToDB(db, StarshipsBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (starships):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()

	go func() {
		films, err := GetCrawlerAsync(&http.Client{}).getFilmsInfo()
		if err != nil {
			log.Printf("Error when crawl films info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range films {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (films):%s\n", err)
				continue
			}

			if err := writeToDB(db, FilmsBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (films):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()

	go func() {
		species, err := GetCrawlerAsync(&http.Client{}).getSpeciesInfo()
		if err != nil {
			log.Printf("Error when crawl species info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range species {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (species):%s\n", err)
				continue
			}

			if err := writeToDB(db, SpeciesBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (species):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()

	go func() {
		vehicles, err := GetCrawlerAsync(&http.Client{}).getVehiclesInfo()
		if err != nil {
			log.Printf("Error when crawl vehicles info:%s\n", err)
			wg.Done()
			return
		}
		for _, val := range vehicles {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Printf("Error when marshal json (vehicles):%s\n", err)
				continue
			}

			if err := writeToDB(db, VehiclesBucket, val.ID, jsonData); err != nil {
				log.Printf("Error when write jsondata to db (vehicles):%s\n", err)
				continue
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

// Write key-[] byte key-value pairs to a bucket named bucket
func writeToDB(db *bolt.DB, bucket, key string, data []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), data)
	})
}

// Request method used by the crawler, while parsing the returned json data into v
func (c *Crawler) requestAndSave(u string, v interface{}) error {
	fmt.Println(u)
	res, err := c.client.Get(u)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("Error when close response:%s\n", err)
		}
	}()

	return json.NewDecoder(res.Body).Decode(v)
}

// Crawl all people info
func (c *Crawler) getPeopleInfo() ([]People, error) {
	peopleInfo := struct {
		Next   *string  `json:"next"`
		Person []People `json:"results"`
	}{}
	requestURL, err := composeUrl("people")
	if err != nil {
		return nil, err
	}

	var people []People
	for {
		err = c.requestAndSave(requestURL, &peopleInfo)
		if err != nil {
			return nil, err
		}

		people = append(people, peopleInfo.Person...)

		if peopleInfo.Next == nil {
			break
		}
		requestURL = *peopleInfo.Next
	}

	for key := range people {
		people[key].ID = getIDFormUrl(people[key].Url)
		replaceSingleOldString(&people[key].HomeWorld)
		replaceSingleOldString(&people[key].Url)
		replaceOldString(people[key].Starships)
		replaceOldString(people[key].Vehicles)
		replaceOldString(people[key].Films)
		replaceOldString(people[key].Species)
	}

	fmt.Println(people)
	return people, nil
}

// Crawl all planets info
func (c *Crawler) getPlanetsInfo() ([]Planets, error) {
	planetsInfo := struct {
		Next   *string   `json:"next"`
		Planet []Planets `json:"results"`
	}{}
	requestURL, err := composeUrl("planets")
	if err != nil {
		return nil, err
	}

	var planets []Planets
	for {
		err = c.requestAndSave(requestURL, &planetsInfo)
		if err != nil {
			return nil, err
		}

		planets = append(planets, planetsInfo.Planet...)
		if planetsInfo.Next == nil {
			break
		}
		requestURL = *planetsInfo.Next
	}

	for key := range planets {
		planets[key].ID = getIDFormUrl(planets[key].Url)
		replaceSingleOldString(&planets[key].Url)
		replaceOldString(planets[key].Films)
		replaceOldString(planets[key].Residents)
	}

	fmt.Println(planets)
	return planets, nil
}

// Crawl all starships info
func (c *Crawler) getStarshipsInfo() ([]Starships, error) {
	starshipsInfo := struct {
		Next     *string     `json:"next"`
		Starship []Starships `json:"results"`
	}{}
	requestURL, err := composeUrl("starships")
	if err != nil {
		return nil, err
	}

	var starships []Starships
	for {
		err = c.requestAndSave(requestURL, &starshipsInfo)
		if err != nil {
			return nil, err
		}

		starships = append(starships, starshipsInfo.Starship...)
		if starshipsInfo.Next == nil {
			break
		}
		requestURL = *starshipsInfo.Next
	}

	for key := range starships {
		starships[key].ID = getIDFormUrl(starships[key].Url)
		replaceSingleOldString(&starships[key].Url)
		replaceOldString(starships[key].Films)
		replaceOldString(starships[key].Pilots)
	}

	fmt.Println(starships)
	return starships, nil
}

// Crawl all films info
func (c *Crawler) getFilmsInfo() ([]Films, error) {
	filmsInfo := struct {
		Next *string `json:"next"`
		Film []Films `json:"results"`
	}{}
	requestURL, err := composeUrl("films")
	if err != nil {
		return nil, err
	}

	var films []Films
	for {
		err = c.requestAndSave(requestURL, &filmsInfo)
		if err != nil {
			return nil, err
		}
		films = append(films, filmsInfo.Film...)

		if filmsInfo.Next == nil {
			break
		}
		requestURL = *filmsInfo.Next
	}

	for key := range films {
		films[key].ID = getIDFormUrl(films[key].Url)
		replaceSingleOldString(&films[key].Url)
		replaceOldString(films[key].Species)
		replaceOldString(films[key].Vehicles)
		replaceOldString(films[key].Starships)
		replaceOldString(films[key].Planets)
		replaceOldString(films[key].Character)
	}

	fmt.Println(films)
	return films, nil
}

// Crawl all species info
func (c *Crawler) getSpeciesInfo() ([]Species, error) {
	speciesInfo := struct {
		Next   *string   `json:"next"`
		Specie []Species `json:"results"`
	}{}
	requestURL, err := composeUrl("species")
	if err != nil {
		return nil, err
	}

	var species []Species
	for {
		err = c.requestAndSave(requestURL, &speciesInfo)
		if err != nil {
			return nil, err
		}
		species = append(species, speciesInfo.Specie...)

		if speciesInfo.Next == nil {
			break
		}
		requestURL = *speciesInfo.Next
	}

	for key := range species {
		species[key].ID = getIDFormUrl(species[key].Url)
		replaceSingleOldString(&species[key].Url)
		replaceSingleOldString(&species[key].Homeworld)
		replaceOldString(species[key].Films)
		replaceOldString(species[key].People)
	}

	fmt.Println(species)
	return species, nil
}

// Crawl all vehicles info
func (c *Crawler) getVehiclesInfo() ([]Vehicles, error) {
	vehiclesInfo := struct {
		Next    *string    `json:"next"`
		Vehicle []Vehicles `json:"results"`
	}{}
	requestURL, err := composeUrl("vehicles")
	if err != nil {
		return nil, err
	}

	var vehicles []Vehicles
	for {
		err = c.requestAndSave(requestURL, &vehiclesInfo)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehiclesInfo.Vehicle...)

		if vehiclesInfo.Next == nil {
			break
		}
		requestURL = *vehiclesInfo.Next
	}

	for key := range vehicles {
		vehicles[key].ID = getIDFormUrl(vehicles[key].Url)
		replaceSingleOldString(&vehicles[key].Url)
		replaceOldString(vehicles[key].Films)
		replaceOldString(vehicles[key].Pilots)
	}

	fmt.Println(vehicles)
	return vehicles, nil
}
