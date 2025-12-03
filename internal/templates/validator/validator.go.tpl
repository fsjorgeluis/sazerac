package validators

import "errors"

func Validate{{ .Entity }}(input any) error {
    // TODO: validation logic here
    return errors.New("validator not implemented")
}