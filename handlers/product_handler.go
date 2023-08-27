package handlers

import (
	"net/http"
	"strconv"

	"github.com/firhan200/rest_fiber/db"
	"github.com/firhan200/rest_fiber/dto"
	"github.com/firhan200/rest_fiber/models"
	repositories "github.com/firhan200/rest_fiber/repositories/product"

	"github.com/gofiber/fiber/v2"
)

var (
	productRepository repositories.IProductRepository
)

func init() {
	dbObj, _ := db.GetConnection()
	productRepository = repositories.NewProductRepository(dbObj, &repositories.ProductRepository{})
}

func HandlerProducts(app *fiber.App) {
	productApi := app.Group("/products")
	productApi.Get("/", getAllProducts)
	productApi.Post("/", insertProduct)
	productApi.Put("/:id", updateProduct)
	productApi.Delete("/:id", deleteProduct)
}

func getAllProducts(c *fiber.Ctx) error {
	channel := make(chan []models.Product)

	go func() {
		products, _ := productRepository.GetAll()
		channel <- products
	}()

	//get from channel
	result := <-channel

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func insertProduct(c *fiber.Ctx) error {
	//get params
	p := &dto.ProductDto{}

	//validate
	if err := c.BodyParser(p); err != nil {
		return HandlerError(c, http.StatusBadRequest, err.Error())
	}

	if p.Name == "" {
		return HandlerError(c, http.StatusBadRequest, "Name cannot be empty")
	}

	if p.Price == 0 {
		return HandlerError(c, http.StatusBadRequest, "Price cannot be empty")
	}

	//insert to db
	isSuccess, err := productRepository.Add(&models.Product{
		Name:  p.Name,
		Price: p.Price,
	})

	if err != nil || !isSuccess {
		return HandlerError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func updateProduct(c *fiber.Ctx) error {
	//get params
	id := c.Params("id", "0")

	if id == "" {
		return HandlerError(c, http.StatusBadRequest, "Id cannot be empty")
	}

	idVal, _ := strconv.Atoi(id)
	idUint := uint(idVal)

	p := &dto.ProductDto{}

	//validate
	if err := c.BodyParser(p); err != nil {
		return HandlerError(c, http.StatusBadRequest, err.Error())
	}

	if p.Name == "" {
		return HandlerError(c, http.StatusBadRequest, "Name cannot be empty")
	}

	if p.Price == 0 {
		return HandlerError(c, http.StatusBadRequest, "Price cannot be empty")
	}

	//update to db
	isSuccess, err := productRepository.Update(idUint, p.Name, p.Price)

	if err != nil || !isSuccess {
		return HandlerError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func deleteProduct(c *fiber.Ctx) error {
	//get params
	id := c.Params("id")

	if id == "" {
		return HandlerError(c, http.StatusBadRequest, "Id cannot be empty")
	}

	idVal, _ := strconv.Atoi(id)
	idUint := uint(idVal)

	//delete
	isSuccess, err := productRepository.Delete(idUint)

	if err != nil || !isSuccess {
		return HandlerError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
