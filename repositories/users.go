package repositories

import (
	"fmt"

	"devstream.in/blogs/models"
)

func RetrieveUserByEmail(email string) (user *models.User, err error) {
	results, err := Db.Query(
		"SELECT id, name, email, username, password FROM users WHERE email = $1;",
		email,
	)

	if err != nil {
		return user, fmt.Errorf("could not retrieve user details for email %s", email)
	}

	defer results.Close()

	if results.Next() {
		err = results.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.Password,
		)

		if err != nil {
			return user, fmt.Errorf("could not parse details for email %s", email)

		}
	} else {
		user = nil
		err = fmt.Errorf("there exists no user for given email")
	}

	return user, err
}

func RetrieveUserById(id string) (user *models.User, err error) {
	results, err := Db.Query(
		"SELECT id, name, email, username, password FROM users WHERE id = $1;",
		id,
	)

	if err != nil {
		return user, fmt.Errorf("could not retrieve user details for id %s", id)
	}

	defer results.Close()

	if results.Next() {
		err = results.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.Password,
		)

		if err != nil {
			return user, fmt.Errorf("could not parse details for id %s", id)

		}
	} else {
		user = nil
		err = fmt.Errorf("there exists no user for given id")
	}

	return user, err
}
func RetrieveUserByUsername(username string) (user *models.User, err error) {
	results, err := Db.Query(
		"SELECT id, name, email, username, password FROM users WHERE username = $1;",
		username,
	)

	if err != nil {
		return user, fmt.Errorf("could not retrieve user details for username %s", username)
	}

	defer results.Close()

	if results.Next() {
		err = results.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.Password,
		)

		if err != nil {
			return user, fmt.Errorf("could not parse details for username %s", username)
		}
	} else {
		user = nil
		err = fmt.Errorf("there exists no user for given username")
	}

	return user, err
}
