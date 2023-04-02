package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
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
	//1-task
	//controller.SortTime()
	//2-task
	//controller.PrintShopCard()
	//3-task
	//controller.PrintTotal()
	//4-task
	//controller.PrintCount()
	//5-task
	//controller.TopTen()
	//6-task
	//controller.TopTenUnder()
	//7-task
	//controller.BestSellerAllTime()
	//9-task
	//controller.ActiveClient()

	//NewShopcard()

}
func NewShopcard() {
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
