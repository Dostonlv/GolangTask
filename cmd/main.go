package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
	"sort"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	//c := controller.NewController(&cfg, jsonDb)

	//Product(c)
	//PrintShopCard()
	//PrintTotal()
	//PrintCount()
	//ActiveClient()
	//TopTen()
	TopTenUnder()
	//New()

}
func New() {
	New := jsonDb.NewShopCartRepo(config.Load().ShopCartFileName)
	New.AddShopCart(&models.Add{
		ProductId: "ce06cf2e-6577-46cb-96a3-1cea379bde4b",
		UserId:    "030b6999-03fd-4f9d-90f5-fd344f24178f",
		Count:     3,
	})
}

func Product(c *controller.Controller) {

	// c.CreateProduct(&models.CreateProduct{
	// 	Name:       "Smartfon vivo V25 8/256 GB",
	// 	Price:      4_860_000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })

	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	products, err := c.GetAllProduct(
		&models.GetListProductRequest{
			Offset:     0,
			Limit:      1,
			CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	for in, product := range products.Products {
		fmt.Println(in+1, product)
	}
}

func Category(c *controller.Controller) {
	// c.CreateCategory(&models.CreateCategory{
	// 	Name:     "Smartfonlar va telefonlar",
	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
	// })

	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(category)

}

func User(c *controller.Controller) {

	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err := c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}

//ProductID -> getByID -> nomini olvolaman
//UserID -> getByID -> nomini olvolaman
//print() -> ClientNAme : Username

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
	type SortCLient controller.Map
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
	type ProCount controller.Map
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
	type SortCLient controller.Map

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
