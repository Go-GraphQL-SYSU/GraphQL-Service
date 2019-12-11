package GraphQL_Service

import (
	"encoding/json"
	"fmt"
	"regexp"

	modelsw "github.com/Go-GraphQL-SYSU/GraphQL-Service/sw"

	"github.com/boltdb/bolt"
)

const peopleBucket = "PeopleBucket"
const filmsBucket = "FilmsBucket"
const planetsBucket = "PlanetsBucket"
const speciesBucket = "SpeciesBucket"
const starshipsBucket = "StarshipsBucket"
const vehiclesBucket = "VehiclesBucket"

const basedURL = "http://localhost:8080/query/api/+[a-zA-Z_]+/"

const replaceURL = ""

func peopleConvertion(p1 *modelsw.People) *People {
	p := &People{}
	p.ID = p1.ID
	p.Name = p1.Name
	p.Height = &p1.Height
	p.Mass = &p1.Mass
	p.HairColor = &p1.HairColor
	p.SkinColor = &p1.SkinColor
	p.EyeColor = &p1.EyeColor
	p.BirthYear = &p1.BirthYear
	p.Gender = &p1.Gender
	return p
}

func planetConvertion(p1 *modelsw.Planets) *Planets {
	p := &Planets{}
	p.ID = p1.ID
	p.Name = p1.Name
	p.RotationPeriod = &p1.RotationPeriod
	p.OrbitalPeriod = &p1.OrbitalPeriod
	p.Diameter = &p1.Diameter
	p.Climate = &p1.Climate
	p.Gravity = &p1.Gravity
	p.Terrain = &p1.Terrain
	p.SurfaceWater = &p1.SurfaceWater
	p.Population = &p1.Population
	return p
}

func filmConvertion(f1 *modelsw.Films) *Films {
	f := &Films{}
	f.ID = f1.ID
	f.Title = f1.Title
	f.EpisodeID = &f1.EpisodeID
	f.OpeningCrawl = &f1.OpeningCrawl
	f.Director = &f1.Director
	f.Producer = &f1.Producer
	f.ReleaseDate = &f1.ReleaseDate
	return f
}

func speciesConvertion(s1 *modelsw.Species) *Species {
	s := &Species{}
	s.ID = s1.ID
	s.Name = &s1.Name
	s.Classification = &s1.Classification
	s.Designation = &s1.Designation
	s.AverageHeight = &s1.AverageHeight
	s.SkinColors = &s1.SkinColors
	s.HairColors = &s1.HairColors
	s.EyeColors = &s1.EyeColors
	s.AverageLifespan = &s1.AverageLifespan
	s.Language = &s1.Language
	return s
}

func starshipsConvertion(s1 *modelsw.Starships) *Starships {
	s := &Starships{}
	s.ID = s1.ID
	s.Name = &s1.Name
	s.Model = &s1.Model
	s.Manufacturer = &s1.Manufacturer
	s.CostInCredits = &s1.CostInCredits
	s.Length = &s1.Length
	s.MaxAtmospheringSpeed = &s1.MaxAtmospheringSpeed
	s.Crew = &s1.Crew
	s.Passengers = &s1.Passengers
	s.CargoCapacity = &s1.CargoCapacity
	s.Consumables = &s1.Consumables
	s.HyperdriveRating = &s1.HyperdriveRating
	s.Mglt = &s1.Mglt
	s.StarshipClass = &s1.StarshipClass
	return s
}

func vehiclesConvertion(v1 *modelsw.Vehicles) *Vehicles {
	v := &Vehicles{}
	v.ID = v1.ID
	v.Name = &v1.Name
	v.Model = &v1.Model
	v.Manufacturer = &v1.Manufacturer
	v.CostInCredits = &v1.CostInCredits
	v.Length = &v1.Length
	v.MaxAtmospheringSpeed = &v1.MaxAtmospheringSpeed
	v.Crew = &v1.Crew
	v.Passengers = &v1.Passengers
	v.CargoCapacity = &v1.CargoCapacity
	v.Consumables = &v1.Consumables
	v.VehicleClass = &v1.VehicleClass
	return v
}

func GetPeopleByID(ID string, db *bolt.DB) (*People, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}
	p1 := &People{}

	p := &modelsw.People{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(peopleBucket))
		v := b.Get([]byte(ID))
		if v == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(v), replaceURL)
		err = json.Unmarshal([]byte(rep), p)
		if err != nil {
			fmt.Println(err)
		}
		p1 = peopleConvertion(p)

		// homeID
		homeID := p.HomeWorld
		homeID = homeID[0 : len(homeID)-1]
		homeBuck := tx.Bucket([]byte(planetsBucket))
		planetData := homeBuck.Get([]byte(homeID))
		planet := &modelsw.Planets{}
		err = json.Unmarshal([]byte(planetData), planet)
		p1.Homeworld = planetConvertion(planet)

		// films
		for _, it := range p.Films {
			it = it[0 : len(it)-1]
			filmBuck := tx.Bucket([]byte(filmsBucket))
			filmData := filmBuck.Get([]byte(it))
			film := &modelsw.Films{}
			err = json.Unmarshal([]byte(filmData), film)
			if err != nil {
				fmt.Println(err)
			}
			p1.Films = append(p1.Films, filmConvertion(film))
		}

		// species
		for _, it := range p.Species {
			it = it[0 : len(it)-1]
			speciesBuck := tx.Bucket([]byte(speciesBucket))
			speciesData := speciesBuck.Get([]byte(it))
			spec := &modelsw.Species{}
			err = json.Unmarshal([]byte(speciesData), spec)
			if err != nil {
				fmt.Println(err)
			}
			p1.Species = append(p1.Species, speciesConvertion(spec))
		}

		// vehicles
		for _, it := range p.Vehicles {
			it = it[0 : len(it)-1]
			vehiclesBuck := tx.Bucket([]byte(vehiclesBucket))
			vehiclesData := vehiclesBuck.Get([]byte(it))
			veh := &modelsw.Vehicles{}
			err = json.Unmarshal([]byte(vehiclesData), veh)
			if err != nil {
				fmt.Println(err)
			}
			p1.Vehicles = append(p1.Vehicles, vehiclesConvertion(veh))
		}

		// starships
		for _, it := range p.Starships {
			it = it[0 : len(it)-1]
			starshipsBuck := tx.Bucket([]byte(starshipsBucket))
			starshipsData := starshipsBuck.Get([]byte(it))
			starship := &modelsw.Starships{}
			err = json.Unmarshal([]byte(starshipsData), starship)
			if err != nil {
				fmt.Println(err)
			}
			p1.Starships = append(p1.Starships, starshipsConvertion(starship))
		}

		return nil
	})

	return p1, err
}

func GetFilmByID(ID string, db *bolt.DB) (*Films, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}

	f1 := &Films{}
	f := &modelsw.Films{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(filmsBucket))
		v := b.Get([]byte(ID))

		if v == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(v), replaceURL)
		err = json.Unmarshal([]byte(rep), f)
		if err != nil {
			fmt.Println(err)
		}
		f1 = filmConvertion(f)

		// characters
		for _, it := range f.Character {
			it = it[0 : len(it)-1]
			pBuck := tx.Bucket([]byte(peopleBucket))
			peopleData := pBuck.Get([]byte(it))
			people := &modelsw.People{}
			err = json.Unmarshal([]byte(peopleData), people)
			if err != nil {
				fmt.Println(err)
			}
			f1.Characters = append(f1.Characters, peopleConvertion(people))
		}

		// planets
		for _, it := range f.Planets {
			it = it[0 : len(it)-1]
			planetsBuck := tx.Bucket([]byte(planetsBucket))
			planetData := planetsBuck.Get([]byte(it))
			planet := &modelsw.Planets{}
			err = json.Unmarshal([]byte(planetData), planet)
			if err != nil {
				fmt.Println(err)
			}
			f1.Planets = append(f1.Planets, planetConvertion(planet))
		}

		// starships
		for _, it := range f.Starships {
			it = it[0 : len(it)-1]
			starshipBuck := tx.Bucket([]byte(starshipsBucket))
			starshipData := starshipBuck.Get([]byte(it))
			starship := &modelsw.Starships{}
			err = json.Unmarshal([]byte(starshipData), starship)
			if err != nil {
				fmt.Println(err)
			}
			f1.Starships = append(f1.Starships, starshipsConvertion(starship))
		}

		// vehicles
		for _, it := range f.Vehicles {
			it = it[0 : len(it)-1]
			vehicleBuck := tx.Bucket([]byte(vehiclesBucket))
			vehicleData := vehicleBuck.Get([]byte(it))
			vehicle := &modelsw.Vehicles{}
			err = json.Unmarshal([]byte(vehicleData), vehicle)
			if err != nil {
				fmt.Println(err)
			}
			f1.Vehicles = append(f1.Vehicles, vehiclesConvertion(vehicle))
		}

		// species
		for _, it := range f.Species {
			it = it[0 : len(it)-1]
			specieBuck := tx.Bucket([]byte(speciesBucket))
			specieData := specieBuck.Get([]byte(it))
			specie := &modelsw.Species{}
			err = json.Unmarshal([]byte(specieData), specie)
			if err != nil {
				fmt.Println(err)
			}
			f1.Species = append(f1.Species, speciesConvertion(specie))
		}
		return nil
	})

	return f1, err
}

func GetPlanetByID(ID string, db *bolt.DB) (*Planets, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}

	p1 := &Planets{}
	p := &modelsw.Planets{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(planetsBucket))
		v := b.Get([]byte(ID))

		if v == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(v), replaceURL)
		err = json.Unmarshal([]byte(rep), p)
		if err != nil {
			fmt.Println(err)
		}
		p1 = planetConvertion(p)

		// residents
		for _, it := range p.Residents {
			it = it[0 : len(it)-1]
			pBuck := tx.Bucket([]byte(peopleBucket))
			peopleData := pBuck.Get([]byte(it))
			people := &modelsw.People{}
			err = json.Unmarshal([]byte(peopleData), people)
			if err != nil {
				fmt.Println(err)
			}
			p1.Residents = append(p1.Residents, peopleConvertion(people))
		}

		// films
		for _, it := range p.Films {
			it = it[0 : len(it)-1]
			filmBuck := tx.Bucket([]byte(filmsBucket))
			filmData := filmBuck.Get([]byte(it))
			film := &modelsw.Films{}
			err = json.Unmarshal([]byte(filmData), film)
			if err != nil {
				fmt.Println(err)
			}
			p1.Films = append(p1.Films, filmConvertion(film))
		}

		return nil
	})

	return p1, err
}

func GetSpeciesByID(ID string, db *bolt.DB) (*Species, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}

	s1 := &Species{}
	s := &modelsw.Species{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(speciesBucket))
		v := b.Get([]byte(ID))

		if v == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(v), replaceURL)
		err = json.Unmarshal([]byte(rep), s)
		if err != nil {
			fmt.Println(err)
		}
		s1 = speciesConvertion(s)

		// Homeworld
		homeID := s.Homeworld
		homeID = homeID[0 : len(homeID)-1]
		homeIDBuck := tx.Bucket([]byte(planetsBucket))
		planetData := homeIDBuck.Get([]byte(homeID))
		planet := &modelsw.Planets{}
		err = json.Unmarshal([]byte(planetData), planet)
		if err != nil {
			fmt.Println(err)
		}
		s1.Homeworld = planetConvertion(planet)

		// people
		for _, it := range s.People {
			it = it[0 : len(it)-1]
			pBuck := tx.Bucket([]byte(peopleBucket))
			peopleData := pBuck.Get([]byte(it))
			people := &modelsw.People{}
			err = json.Unmarshal([]byte(peopleData), people)
			if err != nil {
				fmt.Println(err)
			}
			s1.People = append(s1.People, peopleConvertion(people))
		}

		// films
		for _, it := range s.Films {
			it = it[0 : len(it)-1]
			filmBuck := tx.Bucket([]byte(filmsBucket))
			filmData := filmBuck.Get([]byte(it))
			film := &modelsw.Films{}
			err = json.Unmarshal([]byte(filmData), film)
			if err != nil {
				fmt.Println(err)
			}
			s1.Films = append(s1.Films, filmConvertion(film))
		}

		return nil
	})

	return s1, err
}

func GetStarshipByID(ID string, db *bolt.DB) (*Starships, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}

	s1 := &Starships{}
	s := &modelsw.Starships{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(starshipsBucket))
		v := b.Get([]byte(ID))

		if v == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(v), replaceURL)
		err = json.Unmarshal([]byte(rep), s)
		if err != nil {
			fmt.Println(err)
		}
		s1 = starshipsConvertion(s)

		// pilots
		for _, it := range s.Pilots {
			it = it[0 : len(it)-1]
			pBuck := tx.Bucket([]byte(peopleBucket))
			peopleData := pBuck.Get([]byte(it))
			people := &modelsw.People{}
			err = json.Unmarshal([]byte(peopleData), people)
			if err != nil {
				fmt.Println(err)
			}
			s1.Pilots = append(s1.Pilots, peopleConvertion(people))
		}

		// films
		for _, it := range s.Films {
			it = it[0 : len(it)-1]
			filmBuck := tx.Bucket([]byte(filmsBucket))
			filmData := filmBuck.Get([]byte(it))
			film := &modelsw.Films{}
			err = json.Unmarshal([]byte(filmData), film)
			if err != nil {
				fmt.Println(err)
			}
			s1.Films = append(s1.Films, filmConvertion(film))
		}

		return nil
	})

	return s1, err
}

func GetVehicleByID(ID string, db *bolt.DB) (*Vehicles, error) {
	var err error
	if db == nil {
		db, err = bolt.Open("data/data.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	}

	v1 := &Vehicles{}
	v := &modelsw.Vehicles{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(vehiclesBucket))
		vals := b.Get([]byte(ID))

		if vals == nil {
			return err
		}

		re, _ := regexp.Compile(basedURL)
		rep := re.ReplaceAllString(string(vals), replaceURL)
		err = json.Unmarshal([]byte(rep), v)
		if err != nil {
			fmt.Println(err)
		}
		v1 = vehiclesConvertion(v)

		// pilots
		for _, it := range v.Pilots {
			it = it[0 : len(it)-1]
			pBuck := tx.Bucket([]byte(peopleBucket))
			peopleData := pBuck.Get([]byte(it))
			people := &modelsw.People{}
			err = json.Unmarshal([]byte(peopleData), people)
			if err != nil {
				fmt.Println(err)
			}
			v1.Pilots = append(v1.Pilots, peopleConvertion(people))
		}

		// films
		for _, it := range v.Films {
			it = it[0 : len(it)-1]
			filmBuck := tx.Bucket([]byte(filmsBucket))
			filmData := filmBuck.Get([]byte(it))
			film := &modelsw.Films{}
			err = json.Unmarshal([]byte(filmData), film)
			if err != nil {
				fmt.Println(err)
			}
			v1.Films = append(v1.Films, filmConvertion(film))
		}
		return nil
	})

	return v1, err
}
