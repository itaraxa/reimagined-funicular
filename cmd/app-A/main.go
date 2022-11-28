package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func makeApp() *cli.App {
	return &cli.App{
		Name:        "ACULa",
		Description: "Antivirus Check USB before connect to ULS",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Usage:   "start ACULa",
				Aliases: []string{"s"},
				Action:  startACULa,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "logLevel", Value: "info"},
					&cli.StringFlag{Name: "config", Value: "ACULa.json"},
				},
			},
		},
	}
}

func main() {
	app := makeApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Critical error. Cannot start application: %v\n", err)
	}
}

func startACULa(c *cli.Context) error {
	/*
		1. Проверка новых подключений флешек по КД
		->  В отдельной горутине
			2. Если нашел флешку, то примонтируй ее в режиме для чтения
			3. Запусти ее антивирусную проверку
			4. Если проверка прошла успешно, то создай сертификат
			5. Перемонтируй флешку для записи
			6. Запиши на нее сертификат
			7. Отмонтируй флешку
	*/

	return nil
}
