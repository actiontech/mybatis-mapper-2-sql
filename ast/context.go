package ast

type Context struct {
	QueryType string // select, insert, update, delete
	Variable  map[string]string
	Sqls      map[string]*SqlNode
	Config    *Config
}

func NewContext(config *Config) *Context {
	return &Context{
		Variable: map[string]string{},
		Sqls:     map[string]*SqlNode{},
		Config:   config,
	}
}

func (c *Context) GetVariable(k string) (string, bool) {
	variable, ok := c.Variable[k]
	return variable, ok
}

func (c *Context) SetVariable(k, v string) {
	c.Variable[k] = v
}

func (c *Context) GetSql(k string) (*SqlNode, bool) {
	sql, ok := c.Sqls[k]
	return sql, ok
}

type Config struct {
	SkipErrorQuery bool
	WithQueryId    bool
}

type ConfigFn func() func(*Config)
