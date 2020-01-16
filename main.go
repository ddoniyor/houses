package main

import "sort"

type adresses struct {
	district string
	street string
	distanceFromCenter int64
}


type house struct {
	id int64
	address adresses
	price int64
	areaInKm int64
	roomsAmount int64
}


func sortByPriceAsc(houses []house)[]house{

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}


func sortByPriceDesc(houses []house)[]house{

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}


func sortByDistanceFromCenterAsc(houses []house)[]house{

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].address.distanceFromCenter < result[j].address.distanceFromCenter
	})
	return result
}


func sortByDistanceFromCenterDesc(houses []house)[]house{

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].address.distanceFromCenter > result[j].address.distanceFromCenter
	})
	return result
}

func searchByMaxPrice (houses []house, limit int64)[]house{
	out := make([]house,0)
	copy(out,houses)
	for _, house:=range houses{
		if house.price <= limit{
			out = append (out,house)
		}
	}
	return out
}

func searchByLimitPrice(houses []house, minPrice,maxPrice int64)[]house{
	result := make([]house,0)
	copy(result,houses)
	for _, house := range houses{
		if house.price > minPrice && house.price < maxPrice{
			result = append(result,house)
		}
	}
	return result
}

func searchByDistrict (houses []house,getDistrict string)[]house{
	result := make([]house,0)
	copy(result,houses)
	for _, house := range houses{
		if house.address.district == getDistrict{
			result = append(result,house)
		}
	}
	return result
}

func searchByDistricts(houses []house,getDistricts []string)[]house{
	result := make([]house,0)
	copy(result,houses)
	for _,house := range houses{
		for _,getDistrict := range getDistricts{
			if house.address.district == getDistrict{
				result = append(result,house)
			}
		}
	}
	return result
}

func main() {
}
