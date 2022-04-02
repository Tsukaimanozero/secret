package secret

type Configs struct {
	Name     string
	Password string
}

func (c *Configs) TestConfig() error {
	return nil
}
func TestConfigSelect() error {
	return nil
}
