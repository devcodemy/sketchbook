package construct

type Inventory struct {
	Applications []Application `json:"applications"`
	// Applications map[string]Application `json:"applications"`
}

type Application struct {
	AppName string `json:"appName,omitempty"`
	// AppName     string            `json:"-"`
	Description string            `json:"description"`
	Records     map[string]Record `json:"records"`
	// Fields      map[string]Record `json:"datas"`
}

type Record struct {
	// Name string      `json:"name"`
	Name string      `json:"-"`
	Data interface{} `json:"data"`
	// Name string `json:"name"`
}
