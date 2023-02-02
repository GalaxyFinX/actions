package check

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/dlclark/regexp2"
)

type ECRCheck struct {
	cfg   *aws.Config
	panic bool
}

func NewECRCheck(panic bool) (Checker, error) {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &ECRCheck{
		cfg:   &cfg,
		panic: panic,
	}, nil
}

// extractInfoFromImageName will split your image name into 3 differents part: repo, tag and registryID.
// Note that this is only required by AWS ECR sdk.
func extractInfoFromImageName(imageName string) (repo string, tag string, registryID string) {
	repoRegex := regexp2.MustCompile("(?<=\\/)(.*)(?=\\:)", 0)
	if m, _ := repoRegex.FindStringMatch(imageName); m != nil {
		repo = m.String()
	}

	tagRegex := regexp2.MustCompile("(?<=\\:)(.*)", 0)
	if m, _ := tagRegex.FindStringMatch(imageName); m != nil {
		tag = m.String()
	}

	registryIDRegex := regexp2.MustCompile("^[0-9]*(?=\\.)", 0)
	if m, _ := registryIDRegex.FindStringMatch(imageName); m != nil {
		registryID = m.String()
	}

	return repo, tag, registryID
}

func (ecrCheck ECRCheck) CheckImageTagExist(imageName string) (bool, error) {
	repo, tag, registryID := extractInfoFromImageName(imageName)

	client := ecr.NewFromConfig(*ecrCheck.cfg)
	paginator := ecr.NewListImagesPaginator(
		client,
		&ecr.ListImagesInput{
			RepositoryName: &repo,
			RegistryId:     &registryID,
		},
	)

	for paginator.HasMorePages() {
		out, err := paginator.NextPage(context.TODO())
		if err != nil {
			return false, err
		}

		for _, image := range out.ImageIds {
			if image.ImageTag != nil && *image.ImageTag == tag {
				if ecrCheck.panic {
					return false, errors.New("Tag already exists in the repo.")
				}

				return true, nil
			}
		}
	}

	return false, nil
}
