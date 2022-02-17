package plugins

import "github.com/x-auth/common/models"

type AuthPlugin interface{
  Login(username string, password string) (models.Profile, error)
}
