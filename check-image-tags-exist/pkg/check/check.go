package check

import (
	"errors"
)

type Checker interface {
	// CheckImageTagExist will return true if the tag already exist in your registry,
	// otherwise, return false.
	CheckImageTagExist(imageName string) (bool, error)
}

// NewChecker is Checker factory
func NewChecker(checkOptions *CheckerOptions) (Checker, error) {
	switch checkOptions.RegistryType {
	case "ecr":
		ecrCheck, err := NewECRCheck(checkOptions.Panic)
		if err != nil {
			return nil, err
		}
		return ecrCheck, nil
	// Example of a new regisry type.
	//
	// case "docker":
	// 	dockerCheck, err := NewDockerCheck(checkOptions.Panic)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return dockerCheck, nil
	default:
		return nil, errors.New("Invalid registry type.")
	}
}

type CheckerOptions struct {
	RegistryType string
	Panic        bool
}
