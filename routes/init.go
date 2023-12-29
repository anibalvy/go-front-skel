package routes

import (
	"go-front-skel-001/config"
	"go-front-skel-001/routes/v1"
        "go-front-skel-001/routes/auth"
	"go-front-skel-001/components/index"
	"go-front-skel-001/components/users"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

        "github.com/a-h/templ"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

    base:= app.Group("")
    // render index
    base.Get("/", adaptor.HTTPHandler(templ.Handler(index.Index("John Doe"))))
    base.Get("/login", adaptor.HTTPHandler(templ.Handler(index.Login())))
    base.Post("/login", auth.Login)
    base.Get("/logout", auth.Logout)

    // path: /users
    user := app.Group("/user")
    // JWT Middleware
    user.Use(jwtware.New(jwtware.Config{
	// TokenLookup: "header:Authorization",  //is the default, not need to declare, use auth scheme Bearer, authScheme param is only for header
	// TokenLookup: "cookie:jwt,header:Authorization",  // to allow seamless authorization from browsers
	// TokenLookup: "cookie:jwt",  // to allow seamless authorization from browsers
	// TokenLookup: "query:jwt",  //  other method
	// TokenLookup: "param:jwt",  // other method
	TokenLookup: "header:Authorization,cookie:jwt", // to allow multiple authorization methods
	AuthScheme:  "Bearer",
	SigningKey: jwtware.SigningKey{
		JWTAlg: jwtware.HS256,
		Key:    config.Conf["jwt_secret"]}, // []byte( config.Conf["jwt_secret"].(string)
    }))
    // users.Get( "", handleUser)
    // users.Get("", v1.GetUser)
    user.Get("", adaptor.HTTPHandler(templ.Handler(users.User("id_1"))))
    user.Get("list", v1.GetUserList)
    // users.Post("", handleCreateUser)
    user.Post("", v1.CreateUser)
}
