package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/province/domain"
)

// ProvinceRepository is a struct that implements ProvinceRepository interface.
type ProvinceRepository struct {
	client *http.Client
}

// basePath is a constant that represents the base path of the API endpoint.
const basePath = "http://www.emsifa.com/api-wilayah-indonesia/api"

// New is a function that returns new ProvinceRepository instance.
func New() domain.ProvinceRepository {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &ProvinceRepository{
		client,
	}
}

// GetAll is a method that returns all Provinces.
func (pr *ProvinceRepository) GetAll() ([]domain.Province, error) {
	var provinces []domain.Province

	uri := basePath + "/provinces.json"

	rsp, err := pr.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusNotFound {
		return nil, helper.ErrDataNotFound
	}

	err = json.NewDecoder(rsp.Body).Decode(&provinces)
	if err != nil {
		return nil, err
	}

	return provinces, nil
}

// GetByID is a method that returns the Province with the specified ID.
func (pr *ProvinceRepository) GetByID(id string) (*domain.Province, error) {
	var province domain.Province

	uri := basePath + fmt.Sprintf("/province/%s.json", id)

	rsp, err := pr.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusNotFound {
		return nil, helper.ErrDataNotFound
	}

	err = json.NewDecoder(rsp.Body).Decode(&province)
	if err != nil {
		return nil, err
	}

	return &province, nil
}
