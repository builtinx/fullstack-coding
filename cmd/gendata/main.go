package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

func main() {
	gofakeit.Seed(int64(time.Now().Nanosecond()))

	fmt.Println(`
CREATE TABLE companies (
	id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(256) NOT NULL DEFAULT '',
	mission VARCHAR(256) NOT NULL DEFAULT '',
	PRIMARY KEY (id),
	UNIQUE KEY name (name)
) ENGINE = InnoDB CHARSET=utf8mb4;
`)

	for i := 0; i < 10; i++ {
		fmt.Printf(
			"INSERT INTO companies (name, mission) VALUES ('%s', '%s');\n",
			gofakeit.BuzzWord(),
			gofakeit.Generate("We {hacker.verb} the {company.bs} so you can {hacker.verb} your {hipster.word}"),
		)
	}
}
