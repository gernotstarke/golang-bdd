package basket

import "github.com/cucumber/godog"

// based upon: https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go

// shelf to store articles that are available for shopping
type shopping struct {
	shelf  *Shelf
	basket *Basket
}

func (sh *shopping) addProductToBasket(productName string) error {
	sh.basket.AddItem(productName, sh.shelf.GetProductPrice(productName))
	return nil
}

func (sh *shopping) iShouldHaveProductsInTheBasket(nrOfProductsInBasket int) error {
	return godog.ErrPending
}

func (sh *shopping) theOverallBasketPriceShouldBe(overallPrice int) error {
	return godog.ErrPending
}

func (sh *shopping) addProductToShelf(product string, price float64) (err error) {
	sh.shelf.AddProduct(product, price)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	sh := &shopping{}

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()
	})

	ctx.Step(`^I add the "([^"]*)" to the basket$`, sh.addProductToBasket)
	ctx.Step(`^I should have (\d+) product\(s\) in the basket$`, sh.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be €(\d+)$`, sh.theOverallBasketPriceShouldBe)
	ctx.Step(`^there is a "([^"]*)", which costs €(\d+)$`, sh.addProductToShelf)
}
