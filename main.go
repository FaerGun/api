package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
}

type Order struct {
	ID         uint        `json:"id"`
	Products   []OrderItem `json:"products"`
	TotalPrice int         `json:"total_price"`
	Date       string      `json:"date"`
}

type OrderItem struct {
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	ImageURL  string `json:"image_url"`
}

var orders = []Order{}
var nextOrderID uint = 1

var products = []Product{
	{ID: 1, Name: "Удилище Спиннинговое Major Craft Restive", Description: "Удилище обладает комфортной длинной для береговой джиговой ловли на средних и крупных реках. Также отлично подходит для ловли на колеблющиеся блёсны. При джиговой ловле в идеальных условиях начинает отстукивать с 10-ти грамм.", ImageURL: "https://avatars.mds.yandex.net/get-mpic/3614670/img_id8709495867293679554.jpeg/x332_trim", Price: 8600},
	{ID: 2, Name: "Сумка рыболовная TSURIBITO Superbag JPN 36*23*25cm, с держателями удилищ, серая", Description: "сумка Tsuribito для рыболовных аксессуаров в уникальном дизайне треснутая земля - яркий и функциональный атрибут профессионального рыбака.", ImageURL: "https://main-cdn.sbermegamarket.ru/big2/hlr-system/203/047/558/333/119/3/600011418474b1.jpeg", Price: 2900},
	{ID: 3, Name: "Катушка безынерционная Mifine с байтраннером SPEED", Description: "Mifine Speed - это семейство безынерционок, которые подойдут для установки на фидерные удилища.", ImageURL: "https://avatars.mds.yandex.net/i?id=5c7e0cea0ebafce079e3bdbf605c4993_l-7663003-images-thumbs&n=13", Price: 1600},
	{ID: 4, Name: "Леска плетеная DAIWA J-Braid X4 Yellow 135м", Description: "Высококлассный плетеный шнур отменного японского качества J-Braid X4 от фирмы Daiwa состоит из четырех волокон высококачественного полиэтилена, имеет очень плотное плетение и сечение, максимально приближенное к круглому.", ImageURL: "https://cdn1.ozone.ru/s3/multimedia-i/6543987270.jpg", Price: 900},
	{ID: 5, Name: "Воблер Bassday Sugar Deep Short Bill 75F", Description: "Инженеры Bassday разработали представленные модели для акваторий с неспешным течением, но приличной глубиной. Для успеха на меляках производители рекомендуют вести минноу так, чтобы он ударялся о дно.", ImageURL: "https://avatars.mds.yandex.net/i?id=6341d76a546b8010e127b8e0b1212b8a7f2706d7-9541119-images-thumbs&n=13", Price: 510},
	{ID: 6, Name: "Приманка XXL Fish джиговая Флажок Модель № 2", Description: "Приманка Мандула Флажок XXL Fish Модель №2 Представляем вам новинку - приманку Мандула Флажок XXL Fish. Благодаря своей уловистости, приманка заинтересовала многих любителей рыбалки и получила широкую распространенность.", ImageURL: "https://avatars.mds.yandex.net/get-mpic/5254781/2a0000018e596a0a07ad16b177c4271e3229/orig", Price: 110},
	{ID: 7, Name: "Надувная лодка Leader Тундра-380", Description: "Моторно-гребная лодка с НДНД. Модель представлена в двух размерах 325 мм и 380 мм.  Выполнена из ткани плотностью 750  гр. на кв.м.", ImageURL: "https://avatars.mds.yandex.net/i?id=60a61ae10063417954ce9288a4e97949_l-4264709-images-thumbs&n=13", Price: 49300},
}

var nextID uint = 11

func createOrder(c *gin.Context) {
	var incoming struct {
		Products   []OrderItem `json:"products"`
		TotalPrice int         `json:"total_price"`
	}

	if err := c.ShouldBindJSON(&incoming); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := Order{
		ID:         nextOrderID,
		Products:   incoming.Products,
		TotalPrice: incoming.TotalPrice,
		Date:       "2024-11-28", // Вы можете использовать текущую дату
	}
	nextOrderID++
	orders = append(orders, newOrder)

	c.JSON(http.StatusCreated, gin.H{"message": "Заказ сохранён", "order": newOrder})
}

func getAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func getProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	for _, product := range products {
		if product.ID == uint(id) {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = nextID
	nextID++
	products = append(products, product)
	c.JSON(http.StatusCreated, product)
}

func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	for i, product := range products {
		if product.ID == uint(id) {
			if err := c.ShouldBindJSON(&products[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			products[i].ID = uint(id)
			c.JSON(http.StatusOK, products[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	for i, product := range products {
		if product.ID == uint(id) {
			products = append(products[:i], products[i+1:]...)
			c.Status(http.StatusOK)
			fmt.Println("Product deleted:", product)
			return
		}
	}
	fmt.Println("Product not found with ID:", id)
	c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/products", getAllProducts)       // Получить все продукты
	router.GET("/products/:id", getProductByID)   // Получить продукт по ID
	router.POST("/products", createProduct)       // Создать новый продукт
	router.PUT("/products/:id", updateProduct)    // Обновить продукт по ID
	router.DELETE("/products/:id", deleteProduct) // Удалить продукт по ID
	router.POST("/orders", createOrder)           // Создать новый заказ
	router.GET("/orders", getAllOrders)

	router.Run(":8080")
}
