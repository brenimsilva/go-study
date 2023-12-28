package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/products", getProducts)
    router.POST("/products", addProduct)
    router.PUT("/products", editProduct)
    router.Run("localhost:9090")
}

type product struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Value   float64 `json:"value"`
}

var products = productList {
    {Name: "IPhone 11", Id: 1, Value: 2780.94},
    {Name: "Macbook m1 16 polegadas", Id: 2, Value: 15670.28},
}

func getProducts(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, products)
}

func addProduct(context *gin.Context) {
    var newProduct product
    if err := context.BindJSON(&newProduct); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    products = append(products, newProduct)
    context.IndentedJSON(http.StatusCreated, newProduct)
}

func editProduct(context *gin.Context) {
    var newProduct product
    if err := context.BindJSON(&newProduct); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var product *product = products.getProduct(func(p product) bool {
        return p.Id == newProduct.Id
    })

    *product = newProduct
    context.IndentedJSON(http.StatusOK, newProduct)
}

type productList []product

func (products *productList) filter(f func(product) bool) productList{
    var filteredList = make([]product, 0, cap(*products))

    for _, value := range *products {
        if(f(value)) {
            filteredList = append(filteredList, value)
        }
    }

    return filteredList
}

func (products *productList) getProduct(predicate func(product) bool) *product {
    var product *product
    for _, value := range *products {
        if(predicate(value)) {
            product = &value;
        }
    }
    return product
}
