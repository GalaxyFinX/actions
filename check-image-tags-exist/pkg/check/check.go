package check

import (
	"errors"
)

type Checker interface {
	// CheckImageTagExist will return true if the tag already exist in your registry,
	// otherwise, return false.
	CheckImageTagExist(imageName string) (bool, error)
}

func NewChecker(registryType string) (Checker, error) {
	switch registryType {
	case "ecr":
		ecrCheck, err := NewECRCheck()
		if err != nil {
			return nil, err
		}
		return ecrCheck, nil
	// Example of a new regisry type.
	//
	// case "docker":
	// 	dockerCheck, err := NewDockerCheck()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return dockerCheck, nil
	default:
		return nil, errors.New("Invalid registry type.")
	}
}
