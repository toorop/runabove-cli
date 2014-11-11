package main

import (
	"errors"
)

var (
	ErrBadStoragePath = errors.New("Your storage path is not valid. REGION/CONTAINER/OBJECT")
	ErrCpNoRemotePath = errors.New("You must set at least one remote (on Runabove Storage) path")
)

func ErrNoSuchFileOrDirectory(path string) error {
	return errors.New(path + ": no such file or directory")
}

func ErrUnsuporterPathType(ptype string) error {
	return errors.New(ptype + ": unsuported path type")
}
