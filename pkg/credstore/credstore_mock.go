package credstore

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"

	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
)

type credHubStoreMock struct{}

func (c *credHubStoreMock) Put(key string, credentials interface{}) (interface{}, error) {
	return nil, nil
}

func getFileName(key string) string {
	sub := string(key[14:])
	fileName := strings.ReplaceAll(sub, "/", "-")
	return fileName
}

func (c *credHubStoreMock) PutValue(key string, credentials interface{}) (interface{}, error) {
	credstorePath := os.Getenv("DEV_MODE_ONLY")
	if _, err := os.Stat(credstorePath); os.IsNotExist(err) {
		err = os.MkdirAll(credstorePath, 0777)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create directory for mock credstore")
		}
	}
	file, err := os.Create(credstorePath + "/" + getFileName(key))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create file for mock credstore")
	}

	s := credentials.(string)
	b := []byte(s)

	_, err = file.Write(b)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to write to file for mock credstore")
	}

	file.Close()

	return nil, nil
}

func (c *credHubStoreMock) GetValue(key string) (string, error) {
	credstorePath := os.Getenv("DEV_MODE_ONLY")
	file, err := os.Open(credstorePath + "/" + getFileName(key))
	if err != nil {
		return "", errors.Wrap(err, "Failed to open file for mock credstore")
	}

	content, err := ioutil.ReadAll(file)
	file.Close()
	s := string(content)

	return s, nil
}

func (c *credHubStoreMock) Get(key string) (interface{}, error) {
	return nil, nil
}

func (c *credHubStoreMock) Delete(key string) error {
	return nil
}

func (c *credHubStoreMock) AddPermission(path string, actor string, ops []string) (*permissions.Permission, error) {
	return nil, nil
}

func (c *credHubStoreMock) DeletePermission(path string) error {
	return nil
}
