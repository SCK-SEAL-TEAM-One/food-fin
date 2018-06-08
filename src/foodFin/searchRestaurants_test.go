package foodFin

import (
	"testing"
)

func Test_CalculateDistance_Input_Start_lat_36_dot_12_Start_log_86_dot_67_End_let_33_dot_94_End_log_118_dot_40_Should_be_2887_dot_2599506071097(t *testing.T) {

	startLat := 13.8124384
	startLong := 100.5644465
	endLat := 13.8121383
	endLong := 100.5633476
	expected := 123.26

	actualResult := CalculateDistance(startLat, startLong, endLat, endLong)
	if actualResult != expected {
		t.Error("Expected Distance is:", actualResult, "but it is", expected)
	}
}

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
				Distance:       724.47,
				Location: location{
					Lat:  13.8123479,
					Long: 100.5647322,
				},
			},
			restaurant{
				Name:           "ลาบอุดร",
				RestaurantType: "อาหารอีสาน",
				Price:          "Low",
				Distance:       792.1,
				Location: location{
					Lat:  13.8122417,
					Long: 100.5636951,
				},
			},
			restaurant{
				Name:           "ทิศเหนือ",
				RestaurantType: "อาหารเหนือ",
				Price:          "Low",
				Distance:       832.81,
				Location: location{
					Lat:  13.8120852,
					Long: 100.5630476,
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
