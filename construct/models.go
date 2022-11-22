package construct

type Inventory struct {
	Applications []Application `json:"applications"`
}

type Application struct {
	AppName     string            `json:"appName"`
	Description string            `json:"description"`
	Records     map[string]Record `json:"records"`
}

type Record struct {
	Name string      `json:"-"`
	Data interface{} `json:"data"`
}
