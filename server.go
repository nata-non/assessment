package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nata-non/assessment/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Please use server.go for main file")
	db.AutoMigrate()
	fmt.Println("start at port:", os.Getenv("PORT"))
	list := []model.Expenses{}
	app := fiber.New()
	app.Get("/expenses", func(c *fiber.Ctx) error {
		db.Find(&list)
		return c.Status(http.StatusOK).JSON(list)
	})
	app.Post("/expenses", func(ctx *fiber.Ctx) error {
		//a := new(model.User)
		p := struct {
			Title  string   `json:"title"`
			Amount int      `json:"amount"`
			Note   string   `json:"note"`
			Tags   []string `json:"tags"`
		}{}
		if err := ctx.BodyParser(&p); err != nil {
			return err
		}
		a := model.Expenses{
			Title:  p.Title,
			Amount: p.Amount,
			Note:   p.Note,
			Tags:   p.Tags,
		}
		db.Create(&a)
		return ctx.Status(http.StatusOK).JSON(a)
	})
	app.Get("/expenses/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		db.Find(&list, id)
		return ctx.Status(http.StatusOK).JSON(&list)
	})
	app.Put("/expenses/:id", func(ctx *fiber.Ctx) error {
		p := struct {
			Title  string   `json:"title"`
			Amount int      `json:"amount"`
			Note   string   `json:"note"`
			Tags   []string `json:"tags"`
		}{}
		if err := ctx.BodyParser(&p); err != nil {
			return err
		}
		id := ctx.Params("id")
		db.Find(&list, id)
		db.Model(&list).Updates(map[string]interface{}{"title": p.Title, "amount": p.Amount, "note": p.Note, "tags": &p.Tags})
		return ctx.Status(http.StatusOK).JSON(&list)
	})
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
