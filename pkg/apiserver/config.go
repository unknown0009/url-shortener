package apiserver

type Config struct {
	BindAddr    string `json:"bind_addr"`
	DatabaseUrl string `json:"database_url"`
	Storage     string `json:"storage"`
}
