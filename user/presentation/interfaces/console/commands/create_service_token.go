package commands

import (
	"fmt"
	"reflect"

	"github.com/ecommerce/user/domain/contracts"
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/domain/services"
	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/spf13/cobra"
)

var accessTokenRepository contracts.AccessTokenRepository

func NewCreateServiceTokenCmd(container *persistence.Container) *cobra.Command {
	var name string
	repository, err := container.Resolve(reflect.TypeOf((*contracts.AccessTokenRepository)(nil)))
	if err != nil {
		panic(err)
	}
	accessTokenRepository = repository.(contracts.AccessTokenRepository)

	cmd := &cobra.Command{
		Use:   "create-service-token",
		Short: "Creates a service token",
		Run: func(cmd *cobra.Command, args []string) {
			CreateServiceToken(cmd, args, name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the service that will use the token")

	return cmd
}

func CreateServiceToken(cmd *cobra.Command, args []string, name string) {
	token, err := services.GenerateServiceToken(name)
	if err != nil {
		fmt.Println("Error generating service token:", err)
		return
	}

	accessToken, err := entities.NewAccessTokenEntity(0, name, token)
	if err != nil {
		fmt.Println("Error creating access token:", err)
		return
	}
	accessTokenRepository.Create(accessToken)

	// TODO: assign a role to the token

	fmt.Println("Service token:", token)
	fmt.Println("Access token ID:", accessToken.GetID())
}
