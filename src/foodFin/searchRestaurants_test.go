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
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
				},
			},
			restaurant{
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
				},
			},
			restaurant{
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
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

}
