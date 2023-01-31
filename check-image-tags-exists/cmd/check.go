package cmd

import (
	"errors"
	"fmt"

	"github.com/GalaxyFinX/actions/check-image-tag-exists/pkg/check"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Usage: tagcheck check [OPTIONS] NAME[:TAG] NAME[:TAG] NAME[:TAG]...",
	Long: `Check if a tag already exist in your OCI Registry. 
Return "1" if a tag already exist, otherwise return "0"`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("You must provide atleast 1 image name.")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		registryType, _ := cmd.Flags().GetString("type")
		check, err := check.NewChecker(registryType)
		if err != nil {
			return err
		}

		// The flag indicate whether or not images in the args list exist.
		flag := true
		for _, imageName := range args {
			isExist, err := check.CheckImageTagExist(imageName)
			if err != nil {
				return err
			}
			if isExist == false {
				flag = false
			}
		}

		if flag {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP("type", "t", "ecr", "Type of your OCI registry (currently support: ecr)")
	checkCmd.MarkFlagRequired("type")
}
