package todolist

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	NameArticle string `json:"nameArticle"`
	IsDone      bool   `json:"isDone"`
}

var myMap = make(map[int]Task)

func CreateArticle(c *fiber.Ctx) error {
	task := Task{}

	if err := c.BodyParser(&task); err != nil {
		panic(err)
	}

	myMap[len(myMap)] = task

	for index, taskInMap := range myMap {
		fmt.Printf("%d: %s: %t\n", index, taskInMap.NameArticle, taskInMap.IsDone)
	}

	return nil
}

func UpdateArticle(c *fiber.Ctx) error {

	name := c.Params("name")
	task := Task{}

	if err := c.BodyParser(&task); err != nil {
		panic(err)
	}

	for index, taskInMap := range myMap {
		if taskInMap.NameArticle == name {
			myMap[index] = task

			for index, taskInMap := range myMap {
				fmt.Printf("%d: %s: %t\n", index, taskInMap.NameArticle, taskInMap.IsDone)
			}
			return nil
		}
	}
	return nil
}

func DeleteArticle(c *fiber.Ctx) error {
	name := c.Params("name")

	for index, taskInMap := range myMap {
		if taskInMap.NameArticle == name {
			delete(myMap, index)

			for index, taskInMap := range myMap {
				fmt.Printf("%d: %s: %t\n", index, taskInMap.NameArticle, taskInMap.IsDone)
			}
			return nil
		}
	}

	return nil
}

func ReadArticles(c *fiber.Ctx) error {

	for index, taskInMap := range myMap {
		fmt.Printf("%d: %s: %t\n", index, taskInMap.NameArticle, taskInMap.IsDone)
	}

	return nil
}
