package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cors struct {
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
}

type Certificates struct {
	public  string `json:"public"`
	private string `json:"private"`
}

type ServiceConfigs struct {
	Addres       string       `json:"address"`
	Cors         Cors         `json:"cors"`
	Certificates Certificates `json:"certificates"`
}

type DbConfigs struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserDb   string `json:"user_db"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type Config struct {
	Service ServiceConfigs `json:"service"`
	Db      DbConfigs      `json:"db"`
}

func (d *DbConfigs) GetConnectionString() string {
	cadena := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Bogotá", d.Host, d.Port, d.UserDb, d.DbName, d.Password)
	return cadena
}

func LoadServiceConfig(path string) (*Config, error) {
	var config Config
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config, nil
}
