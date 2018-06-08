package foodFin

import (
	"testing"
)

func Test_searchRestaurants_Input_SCK_DOJOs_Location_Radius_200m_Should_be_3_Results(t *testing.T) {
	sckDojoLocation := location{
		Lat:  13.8073167,
		Long: 100.568995,
	}
	radius := float64(200)
	expected := restaurantsResponse{
		RestaurantList: []restaurant{
			restaurant{
				Name:           "SOUL Food & Drinks",
				RestaurantType: "Cafe",
				Price:          "Low",
				Distance:       150.00,
				Location: location{
					Lat:  13.8073120,
					Long: 100.568980,
				},
			},
			restaurant{
				Name:           "ลาบอุดร",
				RestaurantType: "อาหารอีสาน",
				Price:          "Low",
				Distance:       80.00,
				Location: location{
					Lat:  13.8073284,
					Long: 100.568153,
				},
			},
			restaurant{
				Name:           "ทิศเหนือ",
				RestaurantType: "อาหารเหนือ",
				Price:          "Low",
				Distance:       50.00,
				Location: location{
					Lat:  13.8073120,
					Long: 100.568980,
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
		t.Error("Expected restaurants No. 1 is:", expected.RestaurantList[0], "but it is ", actualResult.RestaurantList[0])
	}
	if expected.RestaurantList[1] != actualResult.RestaurantList[1] {
		t.Error("Expected restaurants No. 2 is:", expected.RestaurantList[1], "but it is ", actualResult.RestaurantList[1])
	}
	if expected.RestaurantList[2] != actualResult.RestaurantList[2] {
		t.Error("Expected restaurants No. 3 is:", expected.RestaurantList[2], "but it is ", actualResult.RestaurantList[2])
	}
}
