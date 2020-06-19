package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "runtime"
	"strings"
	"sync"

	devisecrypto "github.com/consyse/go-devise-encryptor"
	"github.com/rodolfobandeira/go-bcrypt-brute/models"
	"github.com/spf13/cobra"
)

var wg = sync.WaitGroup{}

func main() {
	var cmdConfig = &cobra.Command{
		Use:   "config",
		Short: "Shows your database (MySQL) config",
		Long:  "Reads the environment variables: DB_HOST, DB_USER, DB_PASS",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("\n",
				"DB_NAME: "+os.Getenv("DB_NAME")+"\n",
				"DB_USER: "+os.Getenv("DB_USER")+"\n",
				"DB_PASS: "+os.Getenv("DB_PASS")+"\n",
			)
		},
	}

	var cmdShowUsers = &cobra.Command{
		Use:   "show-users",
		Short: "Shows all users from your database (MySQL)",
		Long:  "Shows all users from your database that are non-archived",
		Run: func(cmd *cobra.Command, args []string) {
			users := models.GetUsers()

			for i, user := range users {
				fmt.Println(i, user)
			}
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "run [10k-most-common-passwords.txt]",
		Short: "Runs go-bcrypt-brute using a passlist file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			passlist := strings.Join(args, " ")
			file, err := os.Open(passlist)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			users := models.GetUsers()
			// runtime.GOMAXPROCS(20)

			fmt.Println("Pulling all active users and checking for weak passwords...")

			for scanner.Scan() {
				for _, user := range users {
					// fmt.Printf("[%d] Processing user_id: %d \n", i, user.ID)
					// done := make(chan bool)
					// go checkPassword(user, scanner.Text(), done)
					// checkPassword(user, scanner.Text())

					rawPassword := scanner.Text()
					encryptedPassword := user.EncryptedPassword
					passwordSalt := user.PasswordSalt

					wg.Add(1)
					go func(user models.User, rawPassword string, encryptedPassword string, passwordSalt string) {
						// fmt.Print("user_id[", user.ID, "]")

						// fmt.Println("Checking user_id: ", user.ID, "with rawPassword: ", rawPassword)

						val := devisecrypto.Compare(rawPassword, passwordSalt, encryptedPassword)

						if val {
							fmt.Print("Weak password user_id[", user.ID, "] -> ", rawPassword, "\n")
						}
						wg.Done()
					}(user, rawPassword, encryptedPassword, passwordSalt)

				}
				wg.Wait()
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			println("Done")

		},
	}

	var rootCmd = &cobra.Command{Use: "go-bcrypt-brute"}
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdConfig)
	rootCmd.AddCommand(cmdShowUsers)
	rootCmd.Execute()
}
