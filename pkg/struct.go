package pkg

type Artisttest struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:firstAlbum"`
	Relations    Relations
}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int 	     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relations struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
