package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bagashiz/gommerce/internal/app/city/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// CityRepository is a struct that implements CityRepository interface.
type CityRepository struct {
	client *http.Client
}

// basePath is a constant that represents the base path of the API endpoint.
const basePath = "http://www.emsifa.com/api-wilayah-indonesia/api"

// New is a function that returns new CityRepository instance.
func New() domain.CityRepository {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &CityRepository{
		client,
	}
}

// GetAll is a method that returns all Cities.
func (pr *CityRepository) GetAll(provinceID string) ([]domain.City, error) {
	var cities []domain.City

	uri := basePath + fmt.Sprintf("/regencies/%s.json", provinceID)

	rsp, err := pr.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusNotFound {
		return nil, helper.ErrDataNotFound
	}

	err = json.NewDecoder(rsp.Body).Decode(&cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

// GetByID is a method that returns the City with the specified ID.
func (pr *CityRepository) GetByID(id string) (*domain.City, error) {
	var city domain.City

	uri := basePath + fmt.Sprintf("/regency/%s.json", id)

	rsp, err := pr.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusNotFound {
		return nil, helper.ErrDataNotFound
	}

	err = json.NewDecoder(rsp.Body).Decode(&city)
	if err != nil {
		return nil, err
	}

	return &city, nil
}
