package main

import "fmt"

var houses = []house{
	{
		1,
		adresses{
			"Сино",
			"Улица Пушкина",
			12},
		30_000,
		80,
		3,
	},
	{2,
		adresses{
			"Сомони",
			"Улица Сатторрова",
			9},
		45_000,
		87,
		4,
	}, {3,
		adresses{
			"Шохмансур",
			"Улица А.Дониша",
			4},
		40_000,
		75,
		3,
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