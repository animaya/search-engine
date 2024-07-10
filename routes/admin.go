package routes

import (
	"animaya/search-engine/db"
	"animaya/search-engine/utils"
	"animaya/search-engine/views"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(c *fiber.Ctx) error {
	return render(c, views.Login())
}

func DeleteAdminUserHandler(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(400).SendString("Email parameter is required")
	}

	var user db.User
	if err := db.DBConn.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	if err := db.DBConn.Delete(&user).Error; err != nil {
		return c.Status(500).SendString("Failed to delete user: " + err.Error())
	}

	return c.SendString("User deleted successfully")
}

type loginform struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func LoginPostHandler(c *fiber.Ctx) error {
	input := loginform{}

	if err := c.BodyParser(&input); err != nil {
		c.Status(500)
		return c.SendString("<h2>Error: something went wrong</h2>")
	}

	user := &db.User{}

	user, err := user.LoginAsAdmin(input.Email, input.Password)

	if err != nil {
		c.Status(401)
		c.Append("content-type", "text/html")
		return c.SendString("<h2>Error: Unauthorised</h2>")
	}

	signedToken, err := utils.CreateNewAuthToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		c.Status(401)
		return c.SendString("<h2>Error: Something went wrong logging in, please try again.</h2>")
	}

	cookie := fiber.Cookie{
		Name:     "admin",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Append("HX-Redirect", "/")
	return c.SendStatus(200)
}

func LogoutHandler(c *fiber.Ctx) error {
	c.ClearCookie("admin")
	c.Set("HX-Redirect", "/login")
	return c.SendStatus(200)
}

type AdminClaims struct {
	User                 string `json:"user"`
	Id                   string `json:"id"`
	jwt.RegisteredClaims `json:"claims"`
}

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("admin")
	if cookie == "" {
		fmt.Println("No admin cookie found")
		return c.Redirect("/login", 302)
	}

	token, err := jwt.ParseWithClaims(cookie, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKeyPath := os.Getenv("ECDSA_PUBLIC_KEY_PATH")
		publicKey, err := utils.LoadECDSAPublicKey(publicKeyPath)
		if err != nil {
			return nil, err
		}
		return publicKey, nil
	})
	if err != nil {
		fmt.Println("JWT parsing failed:", err)
		return c.Redirect("/login", 302)
	}

	claims, ok := token.Claims.(*AdminClaims)
	if ok && token.Valid {
		fmt.Println("JWT token is valid for user:", claims.User)
		return c.Next()
	}

	fmt.Println("JWT token is invalid or claims parsing failed")
	return c.Redirect("/login", 302)
}

type settingsform struct {
	Amount   int    `form:"amount"`
	SearchOn string `form:"seachOn"`
	AddNew   string `form:"addNew"`
}

func DashboardHandler(c *fiber.Ctx) error {
	settings := db.SearchSettings{}

	err := settings.Get()

	if err != nil {
		c.Status(500)
		return c.SendString("<h2>Error: Cannot get settings</h2>")
	}

	amount := strconv.FormatUint(uint64(settings.Amount), 10)

	return render(c, views.Home(amount, settings.SearchOn, settings.AddNew))
}

func DashboardPostHandler(c *fiber.Ctx) error {
	input := settingsform{}

	if err := c.BodyParser(&input); err != nil {
		c.Status(500)

		return c.SendString("<h2>Error: Cannot get settings</h2>")
	}

	addNew := false

	if input.AddNew == "on" {
		addNew = true
	}

	searchOn := false

	if input.SearchOn == "on" {
		searchOn = true
	}

	settings := &db.SearchSettings{}
	settings.Amount = uint(input.Amount)

	settings.SearchOn = searchOn
	settings.AddNew = addNew

	err := settings.Update()

	if err != nil {
		fmt.Println(err)
		return c.SendString("<h2>Error: Cannot update settings</h2>")
	}
	c.Append("HX-Refresh", "true")
	return c.SendStatus(200)
}
