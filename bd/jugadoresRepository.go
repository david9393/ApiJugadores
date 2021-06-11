package bd

import (
	"log"
	"strings"
	"io/ioutil"
	"github.com/david9393/ApiJugadores/models"
)


func CrearJugadores(m []models.Items) error {

	for _, s := range m {
		db := Pool()
		const exec = `INSERT INTO public.jugadores(
			nombre, posicion, nacionalidad, club)
			VALUES ($1, $2, $3, $4);`

		_, err := db.Exec(exec, s.FirstName+" "+s.LastName, s.Position, s.Nation.Name, s.Club.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
func GetJugadoresEquipo(m models.Player) (error, []models.Player) {

	listPlayers := []models.Player{}
   
	db := Pool()
	const exec = `SELECT *
	FROM public.jugadores where  lower(club)=$1`

	rows, err := db.Query(exec, strings.ToLower(m.Club))
	if err != nil {
		log.Fatal(err)
		return err, listPlayers
	}
	defer rows.Close()
	for rows.Next() {
		player := &models.Player{}
		err = rows.Scan(&player.Name, &player.Position, &player.Nacionalidad,&player.Club)
		if err != nil {
			log.Fatal(err)
			return err, listPlayers
		}
		listPlayers = append(listPlayers, *player)
	}

	return nil, listPlayers

}
func GetJugadoresNombre(m models.Player) (error, []models.Player) {

	listPlayers := []models.Player{}
   
	db := Pool()
	const exec = `SELECT *
	FROM public.jugadores where  lower(nombre) LIKE '%' || $1 || '%'`

	rows, err := db.Query(exec, strings.ToLower(m.Club))
	if err != nil {
		log.Fatal(err)
		return err, listPlayers
	}
	defer rows.Close()
	for rows.Next() {
		player := &models.Player{}
		err = rows.Scan(&player.Name, &player.Position, &player.Nacionalidad,&player.Club)
		if err != nil {
			log.Fatal(err)
			return err, listPlayers
		}
		listPlayers = append(listPlayers, *player)
	}

	return nil, listPlayers

}
func MakeMigration() error {
	db := Pool()
    b, err := ioutil.ReadFile("./bd/models.sql")
    if err != nil {
        return err
    }

    rows, err := db.Query(string(b))
    if err != nil {
        return err
    }

    return rows.Close()
}

