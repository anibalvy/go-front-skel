package auth

import (
	"context"
	"encoding/json"
	"time"
	"go-front-skel-001/config"

    // jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func Login(c *fiber.Ctx) error {
    l := config.Logger()
    user := c.FormValue("username")
    pass := c.FormValue("password")
    remember := c.FormValue("remember")

    l.Info("Login: validating %v with passwd=%v, remember=%v", user, pass, remember)

    // check login in db
	conn, err := config.PGconn()
	if err != nil {
		l.Error("error loading config")
	}
	// fmt.Println("conn db: " + conn)
	defer conn.Close(context.Background())

    var query_result string
    err = conn.QueryRow(context.Background(),"select public.fn_user_validate($1,$2)", user, pass).Scan(&query_result)

	if err != nil {
		l.Error("Query failed: %v", err)
        return c.SendStatus(fiber.StatusUnauthorized)
	}

    l.Info("query_result: " + query_result)

    var result map[string]interface{}
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(query_result), &result)
    l.Info("query_result :             %v",  result)
    if result["valid"] != true {
        l.Info("Login: password not valid")
        // return c.SendStatus(fiber.StatusUnauthorized)
        c.Context().SetStatusCode(fiber.StatusUnauthorized)
        return c.Redirect("/")
    }
    l.Info("Login: password is valid")

    // Create Claim
    expiration := time.Now().Add(time.Hour * time.Duration(config.Conf["jwt_expiration_time"].(int)) )
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
    t, err := token.SignedString(config.Conf["jwt_secret"]) // secret already in byte from config
    if err != nil {
        l.Error("token err: %v", err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    l.Info("token: %v", t)
    // return c.JSON(fiber.Map{"token": t, "exp": Conf["jwt_expiration_time"].(int)})
    cookie := fiber.Cookie{
        Name: "jwt",
        Value: t,
        Expires: expiration,
        HTTPOnly: true,
    }

    c.Cookie(&cookie)

    // return c.JSON(fiber.Map{ "logged": true })
    return c.Redirect("/user")
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
