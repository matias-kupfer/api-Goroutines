package service

import (
	"encoding/json"
	"fmt"
	"github.com/matiascfgm/mashup-api/entity"
	"net/http"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type service struct{}

func NewCarDetailService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	// Go routines
	go carService.FetchData()
	go ownerService.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:             car.ID,
		Brand:          car.Brand,
		Model:          car.Model,
		Year:           car.Year,
		OwnerFirstName: owner.FirstName,
		OwnerLastName:  owner.LastName,
		OwnerEmail:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
