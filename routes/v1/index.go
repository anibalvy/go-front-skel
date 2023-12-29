package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go-front-skel-001/config"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetUserList(c *fiber.Ctx) error {
    var films map[string]interface{}
    // films_list := `{"Films": { {"Title": "The Godfather", "Director": "Francis Ford Coppola"}, {"Title": "Blade Runner", "Director": "Ridley Scott"}, {"Title": "The Thing", "Director": "John Carpenter"} }}`
    films_list := `{"Films": [ {"Title": "The Godfather", "Director": "Francis Ford Coppola"}, {"Title": "Blade Runner", "Director": "Ridley Scott"}, {"Title": "The Thing", "Director": "John Carpenter"} ]}`
    json.Unmarshal([]byte(films_list), &films)
    fmt.Printf("films: %v\n", films)
    return c.Render("userlist", fiber.Map(films))
}

func GetUser(c *fiber.Ctx) error {

	// user:= User{
	//   FirstName: "Kanibal",
	//   LastName: "Valdes",
	// }

	conn, err := config.PGconn()
	if err != nil {
		log.Fatal("error loading config")
	}
	// fmt.Println("conn db: " + conn)

	defer conn.Close(context.Background())

	var greeting string
	// err = conn.QueryRow(context.Background(), "select to_jsonb(row(username, email)) from tb_users offset 1").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select json_agg(data) from variables_entorno ").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select json_object_agg(id,data) from variables_entorno ").Scan(&greeting)
	err = conn.QueryRow(context.Background(), "select json_object_agg(data->>'tag',data) from variables_entorno ").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select to_jsonb(row(username, email)) from tb_users limit 1 offset 1;").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select json_agg(to_json(row(username, email))) from tb_users limit 1 offset 1;").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select jsonb_agg(to_jsonb(row(username, email))) from tb_users limit 1 offset 1;").Scan(&greeting)
	// err = conn.QueryRow(context.Background(), "select 'xxxxx' from tb_users limit 1;").Scan(&greeting)
	if err != nil {
		fmt.Fprint(os.Stderr, "Query failed: %v\n", err)
		// os.Exit()
		// return "NOK", err

	}

	fmt.Println("GetUser:  " + greeting)
	// return greeting, err
    // x := map[string]string{}
    // json.Unmarshal(greeting,&x)


    empJson := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager"
	}`
    fmt.Printf("empJson %T\n", empJson)
    // Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	// json.Unmarshal([]byte(empJson), &result)
	json.Unmarshal([]byte(greeting), &result)
	// return c.Status(fiber.StatusOK).JSON(user)
	// return c.Status(fiber.StatusOK).JSON(result)
	return c.Status(fiber.StatusOK).JSON(fiber.Map(result))
	// return c.Status(fiber.StatusOK).
}

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("text/plain", "application/json")

	fmt.Println(c.GetReqHeaders()["Content-Type"])
	new_user := User{}
	if err := c.BodyParser(&new_user); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(new_user)

}
