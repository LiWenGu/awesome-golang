package src

import (
	"errors"
	"fmt"
)

const PrivateKeyFile = "./private.pem"
const PublicKeyFile = "./public.pem"
const ECCPrivateKeyFile = "./eccPrivate.pem"
const ECCPublicKeyFile = "./eccPublic.pem"

func checkErr(message string, err error) error {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return errors.New(message)
	}
	return nil
}
