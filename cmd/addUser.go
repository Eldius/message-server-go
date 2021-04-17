package cmd

import (
	"github.com/eldius/message-server-go/logger"
	"github.com/eldius/message-server-go/repository"

	authRepo "github.com/eldius/jwt-auth-go/repository"
	"github.com/eldius/jwt-auth-go/user"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if c, err := user.NewCredentials(userAddUser, userAddPass); err == nil {
			logger.Logger().Println("admin?", userAddAdmin)
			c.Admin = userAddAdmin
			c.Name = userAddName
			repo := authRepo.NewRepositoryCustom(repository.GetDB())
			repo.SaveUser(&c)
			logger.Logger().Println("User succesfully saved.")
		}
	},
}

var (
	userAddUser  string
	userAddPass  string
	userAddName  string
	userAddAdmin bool
)

func init() {
	userCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&userAddUser, "user", "u", "", "-u <username>")
	addCmd.Flags().StringVarP(&userAddPass, "pass", "W", "", "-W <password>")
	addCmd.Flags().StringVarP(&userAddName, "name", "n", "", "-n <name>")
	addCmd.Flags().BoolVarP(&userAddAdmin, "admin", "a", false, "-a")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
