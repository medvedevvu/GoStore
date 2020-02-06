package GoStore

import (
	"testing"

	sc "github.com/youricorocks/shop_competition"
)

func TestStore_AddProduct1(t *testing.T) {
	/*
	   НЕ добавляет пробникик в товары "sample product can not have price"
	*/
	list := ProductList{}

	product := sc.Product{
		Name:  "Banana",
		Price: 1111.0,
		Type:  sc.ProductNormal,
	}

	product1 := sc.Product{
		Name:  "Banana1",
		Price: 0.0,
		Type:  sc.ProductSample,
	}

	if err := list.AddProduct(product); err != nil && err.(StoreError).errStatus == E {
		t.Fatal(err)
	}

	if err := list.AddProduct(product1); err != nil && err.(StoreError).errStatus == E {
		t.Fatal(err)
	}
	t.Log(list)

}

func TestStore_RemoveProduct1(t *testing.T) {

	list := ProductList{}

	product := sc.Product{
		Name:  "Banana",
		Price: 1111.0,
		Type:  sc.ProductNormal,
	}

	product1 := sc.Product{
		Name:  "Apple",
		Price: 2223.0,
		Type:  sc.ProductNormal,
	}

	product2 := sc.Product{
		Name:  "Orange",
		Price: 0.0,
		Type:  sc.ProductSample,
	}

	list.AddProduct(product)
	list.AddProduct(product1)
	list.AddProduct(product2)
	/*
	   не могу удалить так как Orange не добавился
	*/
	if err := list.RemoveProduct("Orange"); err != nil && err.(StoreError).errStatus == E {
		t.Fatal(err)
	}

	t.Log(list)

}

func TestModifyProduct1(t *testing.T) {
	list := ProductList{}

	product := sc.Product{
		Name:  "Banana",
		Price: 1111.0,
		Type:  sc.ProductNormal,
	}

	product1 := sc.Product{
		Name:  "Apple",
		Price: 2223.0,
		Type:  sc.ProductNormal,
	}

	product2 := sc.Product{
		Name:  "Orange",
		Price: 1000.0,
		Type:  sc.ProductPremium,
	}

	list.AddProduct(product)
	list.AddProduct(product1)
	list.AddProduct(product2)

	product3 := sc.Product{
		Name:  "Orange",
		Price: 0.0,              // 788
		Type:  sc.ProductSample, //ProductNormal,
	}
	/*
	   все операции с ProductSample - sample product can not have price
	*/
	err := list.ModifyProduct(product3)

	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%v %v %v", list["Orange"].Name,
		list["Orange"].Price,
		list["Orange"].Type)
}

func TestAccountMoney(t *testing.T) {
	products := []sc.Product{}

	shop := NewStore()

	product := sc.Product{
		Name:  "Banana",
		Price: 11.11,
		Type:  sc.ProductNormal,
	}

	product1 := sc.Product{
		Name:  "Apple",
		Price: 22.23,
		Type:  sc.ProductNormal,
	}

	product2 := sc.Product{
		Name:  "Orange",
		Price: 10.10,
		Type:  sc.ProductPremium,
	}

	products = append(products, product, product1, product2)

	order := sc.Order{
		Products: products,
		Bundles:  nil, /* реализация другая */
	}

	shop.Register("Dimas")
	shop.AddBalance("Dimas", 100)

	val, err := shop.CalculateOrder("Dimas", order)
	if err != nil && err.(StoreError).errStatus == E {
		t.Logf(" %f --- %v ", val, err)
		acc := shop.GetAccounts(sc.SortByBalance)
		if acc[0].Balance == 100 {
			/* не списались деньги с Account */
			t.Fatalf(" order money=%f  Account bill=%f \n", val, acc[0].Balance)
		}
	}
	t.Fatal(val, err)

}
