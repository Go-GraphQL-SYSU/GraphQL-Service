package crawler

type Films struct {
	ID           string   `json:"id"`
	Url          string   `json:"url"`
	Title        string   `json:"title"`
	EpisodeID    int      `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	ReleaseDate  string   `json:"release_date"`
	Character    []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
}

type People struct {
	ID        string   `json:"id"`
	Url       string   `json:"url"`
	Name      string   `json:"name"`
	BirthYear string   `json:"birth_year"`
	EyeColor  string   `json:"eye_color"`
	Gender    string   `json:"gender"`
	HairColor string   `json:"hair_color"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	SkinColor string   `json:"skin_color"`
	HomeWorld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
}

type Planets struct {
	ID             string   `json:"id"`
	Url            string   `json:"url"`
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
}

type Species struct {
	ID              string   `json:"id"`
	Url             string   `json:"url"`
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	AverageLifespan string   `json:"average_lifespan"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
}

type Starships struct {
	ID                   string   `json:"id"`
	Url                  string   `json:"url"`
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	Mglt                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
}

type Vehicles struct {
	ID                   string   `json:"id"`
	Url                  string   `json:"url"`
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
}
