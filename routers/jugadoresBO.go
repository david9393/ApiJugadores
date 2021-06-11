package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/david9393/ApiJugadores/models"
	"github.com/labstack/echo"
)

func ObtenerJugadores(c echo.Context) error {
	response, err := http.Get("https://www.easports.com/fifa/ultimate-team/api/fut/item?page=1")
	responseObject := models.Respuesta{}
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	///fmt.Println(string(responseData))
	json.Unmarshal(responseData, &responseObject)
	for i := 1; i <= responseObject.TotalPages; i++ {
		url := "https://www.easports.com/fifa/ultimate-team/api/fut/item?page=" + strconv.Itoa(i)
		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		///fmt.Println(string(responseData))
		json.Unmarshal(responseData, &responseObject)

		err = bd.CrearJugadores(responseObject.Items)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

	}
	return c.JSON(http.StatusOK, "Se agrgaron correctamente los datos")

}
