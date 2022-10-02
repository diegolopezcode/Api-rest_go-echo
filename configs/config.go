package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Port     uint   `json:"port"`
	UserDb   string `json:"user_db"`
	Password uint   `json:"password"`
	DbName   string `json:"db_name"`
}

type Config struct {
	Service ServiceConfigs `json:"service"`
	Db      DbConfigs      `json:"db"`
}

func (d *DbConfigs) GetConnectionString() string {
	cadena := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		d.Host,
		d.Port,
		d.UserDb,
		d.Password,
		d.DbName,
	)
	return cadena
}

func LoadServiceConfig(path string) (*ServiceConfigs, error) {
	configs := &ServiceConfigs{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo de configuraciones: %s", err.Error())
	}
	if err := json.Unmarshal(configFile, &configs); err != nil {
		return nil, fmt.Errorf("Error al convertir el archivo de configuraciones: %s", err.Error())
	}
	return configs, nil
}

func LoadDbConfig(path string) (*DbConfigs, error) {
	config := &DbConfigs{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo de configuraciones: %s", err.Error())
	}
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("Error al convertir el archivo de configuraciones: %s", err.Error())
	}
	return config, nil
}
