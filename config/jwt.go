package config

import (
	// "api_v1/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

    // jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func Login(c *fiber.Ctx) error {
    user := c.FormValue("username")
    pass := c.FormValue("password")

    fmt.Printf("Login: validating %v with passwd=%v\n", user, pass)

    // check login in db
	conn, err := PGconn()
	if err != nil {
		log.Fatal("error loading config")
	}
	// fmt.Println("conn db: " + conn)
	defer conn.Close(context.Background())

    var query_result string
    err = conn.QueryRow(context.Background(),"select public.fn_user_validate($1,$2)", user, pass).Scan(&query_result)

	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
        return c.SendStatus(fiber.StatusUnauthorized)
	}

    fmt.Println("query_result: " + query_result)

    var result map[string]interface{}
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(query_result), &result)
    fmt.Printf("query_result :             %v\n",  result)
    if result["valid"] != true {
        fmt.Println("Login: password not valid")
        // return c.SendStatus(fiber.StatusUnauthorized)
        c.Context().SetStatusCode(fiber.StatusUnauthorized)
        return c.Redirect("/")
    }
    fmt.Println("Login: password is valid")

    // Create Claim
    expiration := time.Now().Add(time.Hour * time.Duration(Conf["jwt_expiration_time"].(int)) )
    claims := jwt.MapClaims{
        "name": result["username"],
        "rol":  result["rol"],
        "exp":  expiration.Unix(),
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Generate encoded token and send it as response
    // fmt.Printf("secret: %v\n", Conf["jwt_secret"])
    // fmt.Printf("secret: %v\n", Conf["jwt_secret"])
    t, err := token.SignedString(Conf["jwt_secret"]) // secret already in byte from config
    if err != nil {
        fmt.Printf("token err: %v", err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    fmt.Printf("token: %v\n", t)
    // return c.JSON(fiber.Map{"token": t, "exp": Conf["jwt_expiration_time"].(int)})
    cookie := fiber.Cookie{
        Name: "jwt",
        Value: t,
        Expires: expiration,
        HTTPOnly: true,
    }

    c.Cookie(&cookie)

    // return c.JSON(fiber.Map{ "logged": true })
    return c.Redirect("/users")

}

func Logout(c *fiber.Ctx) error {
    // reset cookie
    cookie := fiber.Cookie{
        Name: "jwt",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HTTPOnly: true,
    }

    c.Cookie(&cookie)

    return c.JSON(fiber.Map{ "logged": false})

}


func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
