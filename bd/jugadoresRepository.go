package bd

import "github.com/david9393/ApiJugadores/models"

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
