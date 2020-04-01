package credstore

import (
	"io/ioutil"
	"os"

	"github.com/pivotal/cloud-service-broker/pkg/config"

	"github.com/pkg/errors"

	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
	"github.com/sirupsen/logrus"
)

//go:generate counterfeiter ./ CredStore
type CredStore interface {
	Put(key string, credentials interface{}) (interface{}, error)
	PutValue(key string, credentials interface{}) (interface{}, error)
	Get(key string) (interface{}, error)
	GetValue(key string) (string, error)
	Delete(key string) error
	AddPermission(path string, actor string, ops []string) (*permissions.Permission, error)
	DeletePermission(path string) error
}

type credhubStore struct {
	credHubClient *credhub.CredHub
	logger        *logrus.Logger
}

func NewCredhubStore(credStoreConfig *config.CredStoreConfig, logger *logrus.Logger) (CredStore, error) {
	if os.Getenv("DEV_MODE_ONLY") != "" {
		logger.Debugf("DEV_MODE_ONLY [%+v] - Creating Mock Credhub", os.Getenv("DEV_MODE_ONLY"))
		return &credHubStoreMock{}, nil
	}

	if !credStoreConfig.HasCredHubConfig() {
		return nil, errors.Errorf("CredHubConfig not found")
	}
	options := []credhub.Option{
		credhub.SkipTLSValidation(credStoreConfig.SkipSSLValidation),
		credhub.Auth(auth.UaaClientCredentials(credStoreConfig.UaaClientName, credStoreConfig.UaaClientSecret)),
		credhub.AuthURL(credStoreConfig.UaaURL),
	}

	if credStoreConfig.CaCertFile != "" {
		dat, err := ioutil.ReadFile(credStoreConfig.CaCertFile)
		if err != nil {
			return nil, err
		}

		if dat == nil {
			return nil, errors.Errorf("CredHub certificate is not valid: %s", credStoreConfig.CaCertFile)
		}
		options = append(options, credhub.CaCerts(string(dat)))
	}

	ch, err := credhub.New(credStoreConfig.CredHubURL, options...)

	if err != nil {
		return nil, err
	}

	return &credhubStore{
		credHubClient: ch,
		logger:        logger,
	}, err
}

func (c *credhubStore) Put(key string, credentials interface{}) (interface{}, error) {
	return c.credHubClient.SetCredential(key, "json", credentials)
}

func (c *credhubStore) PutValue(key string, credentials interface{}) (interface{}, error) {
	return c.credHubClient.SetCredential(key, "value", credentials)
}

func (c *credhubStore) Get(key string) (interface{}, error) {
	return c.credHubClient.GetLatestValue(key)
}

func (c *credhubStore) GetValue(key string) (string, error) {
	value, err := c.credHubClient.GetLatestValue(key)
	if err != nil {
		return "", err
	}
	return string(value.Value), nil
}

func (c *credhubStore) Delete(key string) error {
	return c.credHubClient.Delete(key)
}

func (c *credhubStore) AddPermission(path string, actor string, ops []string) (*permissions.Permission, error) {
	return c.credHubClient.AddPermission(path, actor, ops)
}

func (c *credhubStore) DeletePermission(path string) error {
	allPermissions, err := c.credHubClient.GetPermissions(path)
	if err != nil {
		return err
	}

	for _, permission := range allPermissions {
		p, err := c.credHubClient.GetPermissionByPathActor(path, permission.Actor)
		if err != nil {
			return err
		}
		_, err = c.credHubClient.DeletePermission(p.UUID)
		if err != nil {
			return err
		}

	}

	return err
}
