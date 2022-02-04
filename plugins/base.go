package plugins

type AuthPlugin interface{
  Login(username string, password string) (models.Profile, bool)
}
