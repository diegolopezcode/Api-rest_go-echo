package services

import (
	"errors"
	"strings"

	"github.com/Diegolopez-ing/api_rest/models"
	"github.com/Diegolopez-ing/api_rest/repos"
)

type UsuarioService struct {
	repo *repos.UsuarioRepository
}

/* type UsuarioServices interface {
	FindAll() ([]*models.Usuario, error)
	FindById(id int) (*models.Usuario, error)
	FindByUsername(username string) (*models.Usuario, error)
	FindByUsernameAndPassword(username string, password string) (*models.Usuario, error)
	Save(usuario *models.Usuario) (*models.Usuario, error)
	Update(usuario *models.Usuario) (*models.Usuario, error)
	Delete(usuario *models.Usuario) (*models.Usuario, error)
} */

func NewUsuarioService(repo *repos.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo}
}

func (service *UsuarioService) Loggin(rs models.RequestSession) (*models.Usuario, error) {
	lenUsername := len(strings.TrimSpace(rs.Username))
	lenPassword := len(strings.TrimSpace(rs.Password))
	if lenUsername < 6 || lenPassword < 6 {
		return nil, errors.New("Username or Password is invalid")
	}
	return service.repo.Loging(rs)
}

/*
func (service *UsuarioService) FindAll() ([]*models.Usuario, error) {
	return service.repo.FindAll()
}

func (service *UsuarioService) FindById(id int) (**models.Usuario, error) {
	return service.repo.FindById(id)
}

func (service *UsuarioService) FindByUsername(username string) (*models.Usuario, error) {
	return service.repo.FindByUsername(username)
}

func (service *UsuarioService) FindByUsernameAndPassword(username string, password string) (*models.Usuario, error) {
	return service.repo.FindByUsernameAndPassword(username, password)
}

func (service *UsuarioService) Save(usuario *models.Usuario) (*models.Usuario, error) {
	return service.repo.Save(usuario)
}

func (service *UsuarioService) Update(usuario *models.Usuario) (*models.Usuario, error) {
	return service.repo.Update(usuario)
}

func (service *UsuarioService) Delete(usuario *models.Usuario) (*models.Usuario, error) {
	return service.repo.Delete(usuario)
}
*/
// Language: go
// Path: services/usuario_service.go
// Compare this snippet from models/usuario.go:
// package models
//
