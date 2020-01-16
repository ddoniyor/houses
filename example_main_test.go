package main

import "fmt"

var houses = []house{
	{
			id:1,
			address:adresses{
			district:"Сино",
			street:"Улица Пушкина",
			distanceFromCenter:12},
			price:30_000,
			areaInKm:80,
			roomsAmount:3,
	},
	{
			id:2,
			address:adresses{
			district:"Сомони",
			street:"Улица Сатторрова",
			distanceFromCenter:9},
			price:45_000,
			areaInKm:87,
			roomsAmount:4,
	},
	{
			id:3,
			address:adresses{
			district:"Шохмансур",
			street:"Улица А.Дониша",
			distanceFromCenter:4},
			price:40_000,
			areaInKm:75,
			roomsAmount:3,
	},
}

func ExampleSortByPriceAsc() {
	sortedPriceAsc := sortByPriceAsc(houses)
	fmt.Println(sortedPriceAsc)
	// Output: [{1 {Сино Улица Пушкина 12} 30000 80 3} {3 {Шохмансур Улица А.Дониша 4} 40000 75 3} {2 {Сомони Улица Сатторрова 9} 45000 87 4}]

}

func ExampleSortByPriceDesc() {
	sortedPriceDesc := sortByPriceDesc(houses)
	fmt.Println(sortedPriceDesc)
	// Output: [{2 {Сомони Улица Сатторрова 9} 45000 87 4} {3 {Шохмансур Улица А.Дониша 4} 40000 75 3} {1 {Сино Улица Пушкина 12} 30000 80 3}]
}

func ExampleSortByDistanceFromCenterAsc() {
	sortedDistAsc := sortByDistanceFromCenterAsc(houses)
	fmt.Println(sortedDistAsc)
	// Output: [{3 {Шохмансур Улица А.Дониша 4} 40000 75 3} {2 {Сомони Улица Сатторрова 9} 45000 87 4} {1 {Сино Улица Пушкина 12} 30000 80 3}]
}

func ExampleSortByDistanceFromCenterDesc() {
	sortedDistDesc := sortByDistanceFromCenterDesc(houses)
	fmt.Println(sortedDistDesc)
	// Output: [{1 {Сино Улица Пушкина 12} 30000 80 3} {2 {Сомони Улица Сатторрова 9} 45000 87 4} {3 {Шохмансур Улица А.Дониша 4} 40000 75 3}]
}

func ExampleSearchByMaxPrice_ForManyOutPuts() {
	maxResults := searchByMaxPrice(houses, 300_000)
	fmt.Println(maxResults)
	// Output: [{1 {Сино Улица Пушкина 12} 30000 80 3} {2 {Сомони Улица Сатторрова 9} 45000 87 4} {3 {Шохмансур Улица А.Дониша 4} 40000 75 3}]
}

func ExampleSearchByMaxPrice_ForOneOutPut() {
	oneResult := searchByMaxPrice(houses, 30_000)
	fmt.Println(oneResult)
	// Output: [{1 {Сино Улица Пушкина 12} 30000 80 3}]
}

func ExampleSearchByMaxPrice_NoResults() {
	noResult := searchByMaxPrice(houses, 0)
	fmt.Println(noResult)
	// Output: []
}

func ExampleSearchByLimitPrice_HaveResults() {
	limit := searchByLimitPrice(houses, 20_000, 35_000)
	fmt.Println(limit)
	//Output: [{1 {Сино Улица Пушкина 12} 30000 80 3}]
}
func ExampleSearchByLimitPrice_NoResults() {
	limit := searchByLimitPrice(houses, 0, 0)
	fmt.Println(limit)
	//Output: []
}

func ExampleSearchByDistrict_OneResult() {
	district := searchByDistrict(houses, "Сомони")
	fmt.Println(district)
	//Output: [{2 {Сомони Улица Сатторрова 9} 45000 87 4}]
}
func ExampleSearchByDistrict_NoResult() {
	district := searchByDistrict(houses, "")
	fmt.Println(district)
	//Output: []
}
func ExampleSearchByDistricts_ManyResults() {
	district := searchByDistricts(houses, []string{"Сомони", "Шохмансур"})
	fmt.Println(district)
	//Output: [{2 {Сомони Улица Сатторрова 9} 45000 87 4} {3 {Шохмансур Улица А.Дониша 4} 40000 75 3}]
}

func ExampleSearchByDistricts_NoResults() {
	district := searchByDistricts(houses, []string{"", ""})
	fmt.Println(district)
	//Output: []
}
