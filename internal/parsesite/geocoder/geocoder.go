package geocoder

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/omekov/geminitesttask/internal/parsesite/model"
)

var ApiKey string

const (
	geocodeApiUrl = "https://maps.googleapis.com/maps/api/geocode/json?"
)

func httpRequest(url string) (model.Results, error) {

	var results model.Results

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	if strings.ToUpper(results.Status) != "OK" {
		// If the status is not "OK" check what status was returned
		switch strings.ToUpper(results.Status) {
		case "ZERO_RESULTS":
			err = errors.New("No results found.")
			break
		case "OVER_QUERY_LIMIT":
			err = errors.New("You are over your quota.")
			break
		case "REQUEST_DENIED":
			err = errors.New("Your request was denied.")
			break
		case "INVALID_REQUEST":
			err = errors.New("Probably the query is missing.")
			break
		case "UNKNOWN_ERROR":
			err = errors.New("Server error. Please, try again.")
			break
		default:
			break
		}
	}
	return results, err
}

func Geocoding(address string) (model.Location, error) {

	var location model.Location

	// Create the URL based on the formated address
	u := geocodeApiUrl + "address=" + url.QueryEscape(address)
	if ApiKey != "" {
		u += "&key=" + ApiKey
	}
	results, err := httpRequest(u)
	if err != nil {
		log.Println(err)
		return location, err
	}

	location.Latitude = results.Results[0].Geometry.Location.Lat
	location.Longitude = results.Results[0].Geometry.Location.Lng

	return location, nil
}
