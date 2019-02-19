package models

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Creating a structure for Taxi
type Taxi struct {
	TaxiNo                  int `gorm:"primary_key"` // Unique identification of taxi
	numOfTravellers int `json:"numOfTravellers"` // At present how many number of passengers in taxi
	TaxiCapacity            int `json:"taxiCapacity"` // Maximum number of travellers can travel at once

}
