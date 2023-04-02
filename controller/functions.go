package controller

import (
	"app/config"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
	"sort"
	"time"
)

func PrintShopCard() {

	p, sh, u := ReturnValue()

	for _, shopcard := range sh {
		for _, user := range u {
			for _, prod := range p {

				if shopcard.UserId == user.Id {
					if shopcard.ProductId == prod.Id {

						fmt.Printf(`---------------ShopCard-----------------------
Client Name: %v
Name:%v
Price:%v
Count:%v
Total:%v,
Time: %v
`, user.Name, prod.Name, prod.Price, shopcard.Count, prod.Price*float64(shopcard.Count), shopcard.Time)
					}
				}
			}
		}
	}

}

func PrintTotal() {
	var sum float64
	p, sh, u := ReturnValue()
	//v := " "
	MapTotal := make(map[string]float64)

	for _, shopcard := range sh {
		sum = 0
		for _, user := range u {
			for _, prod := range p {

				if shopcard.UserId == user.Id {
					if shopcard.ProductId == prod.Id {
						sum += prod.Price * float64(shopcard.Count)
						MapTotal[user.Name] = sum
					}
				}
			}
		}
	}
	type SortCLient Map
	var tp []SortCLient
	for k, v := range MapTotal {
		tp = append(tp, SortCLient{k, int(v)})
	}
	for _, v := range tp {
		fmt.Printf(`-----------Total Price-------------
Name:%v,
TotalPrice:%v
`, v.Key, v.Value)
	}

}

func PrintCount() {
	var count int
	ProdCount := make(map[string]int)
	p, sh, _ := ReturnValue()
	for _, shop := range sh {
		count = 0
		for _, prod := range p {
			if shop.ProductId == prod.Id {
				count += shop.Count
				ProdCount[prod.Name] = count
			}
		}
	}
	type ProCount Map
	var sc []ProCount
	for k, v := range ProdCount {
		sc = append(sc, ProCount{k, v})
	}
	for _, v := range sc {
		fmt.Printf(`-----------ProductCount------------
ProductName: %v
Count:%v
`, v.Key, v.Value)

	}

}

func ReturnValue() ([]models.Product, []models.ShopCart, []models.User) {
	ShopCard := jsonDb.NewShopCartRepo(config.Load().ShopCartFileName)
	sh, err := ShopCard.Read()
	if err != nil {
		fmt.Println(err)
	}
	Product := jsonDb.NewProductRepo(config.Load().ProductFileName)
	p, err := Product.Read()
	if err != nil {
		fmt.Println(err)
	}
	User := jsonDb.NewUserRepo(config.Load().UserFileName)
	u, err := User.Read()
	if err != nil {
		fmt.Println(err)
	}
	return p, sh, u
}

func ActiveClient() {
	_, sh, u := ReturnValue()

	TopClient := make(map[string]int)
	for _, v := range u {
		count := 0
		for _, shop := range sh {

			if v.Id == shop.UserId {
				count++
				TopClient[v.Name] = count
			}
		}

	}
	type SortCLient Map

	var sc []SortCLient
	for k, v := range TopClient {
		sc = append(sc, SortCLient{k, v})
	}

	sort.Slice(sc, func(i, j int) bool {
		return sc[i].Value > sc[j].Value
	})

	for _, v := range sc {
		fmt.Printf(" %s: %d\n", v.Key, v.Value)
		break
	}
}

func TopTen() {
	var count int
	TopTen := make(map[string]int)
	p, sh, _ := ReturnValue()
	for _, shop := range sh {
		count = 0
		for _, prod := range p {
			if shop.ProductId == prod.Id {
				count += shop.Count
				TopTen[prod.Name] = count
			}
		}
	}
	keys := make([]string, 0, len(TopTen))

	for key := range TopTen {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return TopTen[keys[i]] > TopTen[keys[j]]
	})
	fmt.Println("------------------top ten selling products-------------")
	for _, k := range keys {
		fmt.Printf(`Name: %v Count: %v
`, k, TopTen[k])
	}
}
func TopTenUnder() {
	var count int
	TopTenUnder := make(map[string]int)
	p, sh, _ := ReturnValue()
	for _, shop := range sh {
		count = 0
		for _, prod := range p {
			if shop.ProductId == prod.Id {
				count += shop.Count
				TopTenUnder[prod.Name] = count
			}
		}
	}
	keys := make([]string, 0, len(TopTenUnder))

	for key := range TopTenUnder {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return TopTenUnder[keys[i]] < TopTenUnder[keys[j]]
	})

	fmt.Println("------------------top ten undersellers-------------")
	for _, k := range keys {
		fmt.Printf(`Name: %v Count: %v
`, k, TopTenUnder[k])
	}
}

func SortTime() {
	p, sh, u := ReturnValue()
	var orders = sh
	fromDate, err := time.Parse("2006-01-02 15:04:05", "2022-04-04 04:41:01")
	if err != nil {
		log.Fatal(err)
	}

	toDate, err := time.Parse("2006-01-02 15:04:05", "2023-04-01 13:27:39")
	if err != nil {
		log.Fatal(err)
	}

	filteredOrders := make([]models.ShopCart, 0)

	for _, order := range orders {
		orderTime, err := time.Parse("2006-01-02 15:04:05", order.Time)
		if err != nil {
			log.Fatal(err)
		}
		if orderTime.After(fromDate) && orderTime.Before(toDate) {
			filteredOrders = append(filteredOrders, order)
		}
	}

	sort.Slice(filteredOrders, func(i, j int) bool {
		return filteredOrders[i].Time < filteredOrders[j].Time
	})
	for _, shopcard := range filteredOrders {
		for _, user := range u {
			for _, prod := range p {

				if shopcard.UserId == user.Id {
					if shopcard.ProductId == prod.Id {

						fmt.Printf(`---------------ShopCard-----------------------
Client Name: %v
Name:%v
Price:%v
Count:%v
Total:%v,
Time: %v
`, user.Name, prod.Name, prod.Price, shopcard.Count, prod.Price*float64(shopcard.Count), shopcard.Time)
					}
				}
			}
		}
	}
}

func BestSellerAllTime() {
	type nimadi struct {
		Name  string
		Time  string
		Count int
	}
	var data []nimadi
	var count int
	//TopTenUnder := make(map[string]int)
	p, sh, _ := ReturnValue()
	var v string
	for _, shop := range sh {
		for _, prod := range p {
			if shop.ProductId == prod.Id {
				count = shop.Count
				v = prod.Name
				data1 := nimadi{v, shop.Time, count}
				data = append(data, data1)
				//fmt.Println("Name:", v, "Time: ", shop.Time, "count:", count)
			}
		}
	}

	counts := make(map[string]int)
	maxCount := 0
	result := make(map[string]nimadi)

	for _, d := range data {
		count, ok := counts[d.Name]
		if !ok {
			count = 0
		}
		counts[d.Name] = d.Count

		if d.Count > maxCount {
			maxCount = d.Count
		}

		if d.Count < count {
			continue
		}

		if d.Count > count {
			result[d.Name] = d
			continue
		}

		existing, ok := result[d.Name]
		if !ok {
			result[d.Name] = d
			continue
		}

		if d.Count > existing.Count {
			result[d.Name] = d
		}
	}

	var filtered []nimadi

	for _, r := range result {
		filtered = append(filtered, r)
	}

	for _, v := range filtered {
		fmt.Printf(`-----------------------
Name: %v,
Time: %v,
Count: %v
`, v.Name, v.Time, v.Count)
	}
}
