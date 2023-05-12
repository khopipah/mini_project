package route

import (
	"mini_project/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	// docter collection
	docter := e.Group("/docters")
	docter.GET("", controller.GetDoctersController)
	docter.GET("/:id", controller.GetDocterController)
	docter.POST("", controller.CreateDocterController)
	docter.PUT("/:id", controller.UpdateDocterController)
	docter.DELETE("/:id", controller.DeleteDocterController)

	// pasien collection
	pasien := e.Group("/pasiens")
	pasien.GET("", controller.GetPasiensController)
	pasien.GET("/:id", controller.GetPasienController)
	pasien.POST("", controller.CreatePasienController)
	pasien.PUT("/:id", controller.UpdatePasienController)
	pasien.DELETE("/:id", controller.DeletePasienController)

	// room collection
	room := e.Group("/rooms")
	room.GET("", controller.GetRoomsController)
	room.GET("/:id", controller.GetRoomController)
	room.POST("", controller.CreateRoomController)
	room.PUT("/:id", controller.UpdateRoomController)
	room.DELETE("/:id", controller.DeleteRoomController)

}
