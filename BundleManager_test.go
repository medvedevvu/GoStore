package GoStore

import (
	"fmt"
	"testing"

	"github.com/youricorocks/shop_competition"
	sc "github.com/youricorocks/shop_competition"
)

func TestAddWrongNameProducts(t *testing.T) {
	// var shop shop_competition.Shop
	shop := NewStore()

	product := shop_competition.Product{
		Name:  "",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	if err := shop.AddProduct(product); err == nil /*&& err.(StoreError).errStatus == I*/ {
		t.Fatalf(" могу добавлять товар с пустым названием %v", product)
	}
	t.Fatalf("Finish %v", shop.ProductList)
}

func TestWhen3ProductToBundle(t *testing.T) {

	bundles := Bundles{}

	product1 := shop_competition.Product{
		Name:  "AAAA",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product2 := shop_competition.Product{
		Name:  "YYYYY",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product3 := shop_competition.Product{
		Name:  "ZZZZZ",
		Price: 100,
		Type:  shop_competition.ProductPremium,
	}

	err := bundles.AddBundle("Карзин1", product3, 1.999, product2, product1)

	if err != nil {
		t.Fatalf("%v", err)
	}

}

func TestWhen2ProductToBundle(t *testing.T) {

	bundles := Bundles{}

	product1 := shop_competition.Product{
		Name:  "AAAA",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product2 := shop_competition.Product{
		Name:  "YYYYY",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	err := bundles.AddBundle("Карзина1", product1, 1.999, product2)
	if err != nil {
		t.Fatalf("%v", err)
	}

}

func TestBundlesDisc(t *testing.T) {

	product1 := shop_competition.Product{
		Name:  "AAAA",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product2 := shop_competition.Product{
		Name:  "YYYYY",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product3 := shop_competition.Product{
		Name:  "ZZZZZ",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	products := []sc.Product{product1, product2, product3}

	bundles := Bundles{}
	bundles["Корзина"] = &sc.Bundle{products, sc.BundleNormal, 45}
	err := bundles.ChangeDiscount("Корзина", 50)

	if err != nil {
		t.Fatalf("%v", err)
	}

	dsk := bundles["Корзина"].Discount
	if dsk != 0.5 {
		t.Fatalf("%f == %f", dsk, 0.5)
	}
	fmt.Printf("%f == %f\n", dsk, 0.5)
}

func TestRemoveBundles(t *testing.T) {

	product1 := shop_competition.Product{
		Name:  "AAAA",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product2 := shop_competition.Product{
		Name:  "YYYYY",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product3 := shop_competition.Product{
		Name:  "ZZZZZ",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product4 := shop_competition.Product{
		Name:  "ZZZZZ",
		Price: 100,
		Type:  shop_competition.ProductSample,
	}

	products := []sc.Product{product1, product2, product3}
	products1 := []sc.Product{product1, product2, product4}

	bundles := Bundles{}
	bundles["Корзина"] = &sc.Bundle{products, sc.BundleNormal, 45}
	bundles["Корзина1"] = &sc.Bundle{products1, sc.BundleSample, 45}

	err := bundles.RemoveBundle("Корзина1")

	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(bundles)

}
