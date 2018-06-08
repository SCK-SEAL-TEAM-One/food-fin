package foodFin

import (
	"testing"
)

func Test_searchRestaurants_Input_SCK_DOJOs_Location_Radius_200m_Should_be_3_Results(t *testing.T) {
	sckDojoLocation := location{
		lat:  13.8073167,
		long: 100.568995,
	}
	radius := float64(200)
	expected := restaurantsResponse{
		RestaurantList: []restaurant{
			restaurant{
				name:           "SOUL Food & drintcs",
				restaurantType: "Cafe",
				price:          "Low",
				distance:       150.00,
				location: location{
					lat:  13.8073120,
					long: 100.568980,
				},
			},
			restaurant{
				name:           "ลาบอุดร",
				restaurantType: "อาหารอีสาน",
				price:          "Low",
				distance:       80.00,
				location: location{
					lat:  13.8073284,
					long: 100.568153,
				},
			},
			restaurant{
				name:           "ทิศเหนือ",
				restaurantType: "อาหารเหนือ",
				price:          "Low",
				distance:       150.00,
				location: location{
					lat:  13.8073120,
					long: 100.568980,
				},
			},
		},
	}
	expectedNumberOfRestaurants := 3

	actualResult := searchRestaurants(sckDojoLocation, radius)
	actualNumberOfRestaurants := len(actualResult.RestaurantList)
	if expectedNumberOfRestaurants != actualNumberOfRestaurants {
		t.Error("Expected Number of restaurants is:", expectedNumberOfRestaurants, "but it is ", actualNumberOfRestaurants)
	}
	if expected.RestaurantList[0] != actualResult.RestaurantList[0] {
		t.Error("Expected restaurants No. 1 is:", expected, "but it is ", actualResult)
	}
	if expected.RestaurantList[1] != actualResult.RestaurantList[1] {
		t.Error("Expected restaurants No. 2 is:", expected, "but it is ", actualResult)
	}
	if expected.RestaurantList[2] != actualResult.RestaurantList[2] {
		t.Error("Expected restaurants No. 3 is:", expected, "but it is ", actualResult)
	}
}
