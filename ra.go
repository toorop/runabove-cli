package main

import (
	"fmt"
	//"github.com/codegangsta/cli"
	"bufio"
	//"log"
	"github.com/Toorop/goabove"
	raApiAuth "github.com/Toorop/goabove/auth"
	raToken "github.com/Toorop/goabove/token"
	"github.com/Toorop/gopenstack"
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"os"
	"runtime"
)

var (
	raClient  *goabove.ApiClient // Runabove API client
	osClient  *gopenstack.Client
	osKeyring *gopenstack.Keyring
)

func init() {

	// Check keyring
	// Runabove API
	// Check for Consumer key
	ck := os.Getenv("RA_CONSUMER_KEY")

	// if No ConsumerKey, request one
	if len(ck) == 0 {
		var r []byte
		fmt.Println("No Runabove API consumer key (RA_CONSUMER_KEY) found in environnement vars !")
		for {
			fmt.Print("Do you have valid Consumer Key for runabove-cli ? (y/n) : ")
			r, _, _ = bufio.NewReader(os.Stdin).ReadLine()
			if r[0] == 110 || r[0] == 121 {
				break
			}
		}
		// Yes
		if r[0] == 121 {
			fmt.Println("\r\nThen run the following command :", NL)
			if runtime.GOOS == "windows" {
				fmt.Println("SET RA_CONSUMER_KEY=your_consumer_key", NL)
			} else {
				fmt.Println("export RA_CONSUMER_KEY=your_consumer_key", NL)
			}
			fmt.Println("and restart runabove-cli application.", NL)
			os.Exit(0)
		}

		// No get new credentials
		raCredentials, err := raApiAuth.GetApiCredential(RA_APPLICATION_KEY)
		dieOnError(err)

		fmt.Println("Your consumer key is :", raCredentials.ConsumerKey)

		fmt.Println("Now you need to validate it :")
		if runtime.GOOS != "windows" {
			fmt.Printf("If you have a browser available on this machine it will open to the validation page.\n\tIf not, copy and paste the link below in a browser to validate your key :\r\n\r\n%s\r\n", raCredentials.ValidationUrl)
			webbrowser.Open(raCredentials.ValidationUrl)
		} else {
			fmt.Printf("To do it just copy and paste the link below in a browser and follow instructions on Runabove website :\r\n\r\n%s\r\n", raCredentials.ValidationUrl)
		}

		fmt.Println("\r\nWhen it will be done run the following command : \r\n")
		if runtime.GOOS == "windows" {
			fmt.Printf("SET RA_CONSUMER_KEY=%s%s%s", raCredentials.ConsumerKey, NL, NL)
		} else {
			fmt.Printf("export RA_CONSUMER_KEY=%s%s%s", raCredentials.ConsumerKey, NL, NL)
		}
		fmt.Println("and restart runabove-cli.\r\n")
		os.Exit(0)
	}

	// We have consumer key (ck)
	// populate openstack keyring
	raClient = goabove.NewClient(RA_APPLICATION_KEY, RA_APPLICATION_SECRET, ck)
	token, err := raToken.New(raClient)
	dieOnError(err)

	osKeyring, err = token.GetGosKeyring()
	dieOnError(err)
}

func main() {
	app := cli.NewApp()
	app.Name = "ra"
	app.Usage = "runabove-cli (aka ra) brings Runabove services to the command line."
	app.Version = VERSION
	app.Author = "Stéphane Depierrepont aka Toorop"
	app.Email = "toorop@toorop.fr"
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [section] [command] [arguments]

SECTIONS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Description}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.CommandHelpTemplate = `NAME:
   {{.Name}} - {{.Description}}

USAGE:
   {{.Usage}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.SubcommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} command [command options] [arguments...]

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Description}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	// default action: help
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:        "storage",
			Usage:       "ra storage subcommand",
			Description: "Manage objects storage (Swift)",
			Subcommands: getObjectStorageCmds(),
		},
	}

	app.Run(os.Args)
	return
}
