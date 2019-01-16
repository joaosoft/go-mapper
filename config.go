package mapper

// AppConfig ...
type AppConfig struct {
	Mapper *MapperConfig `json:"mapper"`
}

// MapperConfig ...
type MapperConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}
