package gform

import (
	"sync"
	"time"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"

	"gofly/utils/tools/gcache"
	"gofly/utils/tools/glog"
)

// Config is the configuration management object.
type Config map[string]ConfigGroup

// ConfigGroup is a slice of configuration node for specified named group.
type ConfigGroup []ConfigNode

// ConfigNode is configuration for one node.
type ConfigNode struct {
	Hostname             string        `json:"hostname"`             // Host of server, ip or domain like: 127.0.0.1, localhost
	Hostport             string        `json:"hostport"`             // Port, it's commonly 3306.
	Username             string        `json:"username"`             // Authentication username.
	Password             string        `json:"password"`             // Authentication password.
	Dbname               string        `json:"dbname"`               // Default used database name.
	Type                 string        `json:"type"`                 // Database type: mysql, sqlite, mssql, pgsql, oracle.
	Sqlmode              string        `json:"sqlmode"`              // Set the database sql_mode, when the database type mariadb/mysql can be set
	Link                 string        `json:"link"`                 // (Optional) Custom link information for all configuration in one single string.
	Extra                string        `json:"extra"`                // (Optional) Extra configuration according the registered third-party database driver.
	Role                 string        `json:"role"`                 // (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
	Debug                bool          `json:"debug"`                // (Optional) Debug mode enables debug information logging and output.
	Prefix               string        `json:"prefix"`               // (Optional) Table prefix.
	DryRun               bool          `json:"dryRun"`               // (Optional) Dry run, which does SELECT but no INSERT/UPDATE/DELETE statements.
	Weight               int           `json:"weight"`               // (Optional) Weight for load balance calculating, it's useless if there's just one node.
	Charset              string        `json:"charset"`              // (Optional, "utf8" in default) Custom charset when operating on database.
	Protocol             string        `json:"protocol"`             // (Optional, "tcp" in default) See net.Dial for more information which networks are available.
	Timezone             string        `json:"timezone"`             // (Optional) Sets the time zone for displaying and interpreting time stamps.
	Namespace            string        `json:"namespace"`            // (Optional) Namespace for some databases. Eg, in pgsql, the `Name` acts as the `catalog`, the `NameSpace` acts as the `schema`.
	MaxIdleConnCount     int           `json:"maxIdle"`              // (Optional) Max idle connection configuration for underlying connection pool.
	MaxOpenConnCount     int           `json:"maxOpen"`              // (Optional) Max open connection configuration for underlying connection pool.
	MaxConnLifeTime      time.Duration `json:"maxLifeTime"`          // (Optional) Max amount of time a connection may be idle before being closed.
	QueryTimeout         time.Duration `json:"queryTimeout"`         // (Optional) Max query time for per dql.
	ExecTimeout          time.Duration `json:"execTimeout"`          // (Optional) Max exec time for dml.
	TranTimeout          time.Duration `json:"tranTimeout"`          // (Optional) Max exec time for a transaction.
	PrepareTimeout       time.Duration `json:"prepareTimeout"`       // (Optional) Max exec time for prepare operation.
	CreatedAt            string        `json:"createtime"`           // (Optional) The filed name of table for automatic-filled created datetime.
	UpdatedAt            string        `json:"updatetime"`           // (Optional) The filed name of table for automatic-filled updated datetime.
	DeletedAt            string        `json:"deletetime"`           // (Optional) The filed name of table for automatic-filled updated datetime.
	TimeMaintainDisabled bool          `json:"timeMaintainDisabled"` // (Optional) Disable the automatic time maintaining feature.
}

const (
	DefaultGroupName = "default" // Default group name.
)

// configs specifies internal used configuration object.
var configs struct {
	sync.RWMutex
	config Config // All configurations.
	group  string // Default configuration group.
}

func init() {
	configs.config = make(Config)
	configs.group = DefaultGroupName
}

// SetConfig sets the global configuration for package.
// It will overwrite the old configuration of package.
func SetConfig(config Config) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	for k, nodes := range config {
		for i, node := range nodes {
			nodes[i] = parseConfigNode(node)
		}
		config[k] = nodes
	}
	configs.config = config
}

// SetConfigGroup sets the configuration for given group.
func SetConfigGroup(group string, nodes ConfigGroup) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	for i, node := range nodes {
		nodes[i] = parseConfigNode(node)
	}
	configs.config[group] = nodes
}

// AddConfigNode adds one node configuration to configuration of given group.
func AddConfigNode(group string, node ConfigNode) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	configs.config[group] = append(configs.config[group], parseConfigNode(node))
}

// parseConfigNode parses `Link` configuration syntax.
func parseConfigNode(node ConfigNode) ConfigNode {
	if node.Link != "" {
		node = *parseConfigNodeLink(&node)
	}
	if node.Link != "" && node.Type == "" {
		match, _ := gregex.MatchString(`([a-z]+):(.+)`, node.Link)
		if len(match) == 3 {
			node.Type = gstr.Trim(match[1])
			node.Link = gstr.Trim(match[2])
		}
	}
	return node
}

// AddDefaultConfigNode adds one node configuration to configuration of default group.
func AddDefaultConfigNode(node ConfigNode) {
	AddConfigNode(DefaultGroupName, node)
}

// AddDefaultConfigGroup adds multiple node configurations to configuration of default group.
func AddDefaultConfigGroup(nodes ConfigGroup) {
	SetConfigGroup(DefaultGroupName, nodes)
}

// GetConfig retrieves and returns the configuration of given group.
func GetConfig(group string) ConfigGroup {
	configs.RLock()
	defer configs.RUnlock()
	return configs.config[group]
}

// SetDefaultGroup sets the group name for default configuration.
func SetDefaultGroup(name string) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	configs.group = name
}

// GetDefaultGroup returns the { name of default configuration.
func GetDefaultGroup() string {
	defer instances.Clear()
	configs.RLock()
	defer configs.RUnlock()
	return configs.group
}

// IsConfigured checks and returns whether the database configured.
// It returns true if any configuration exists.
func IsConfigured() bool {
	configs.RLock()
	defer configs.RUnlock()
	return len(configs.config) > 0
}

// SetLogger sets the logger for orm.
func (c *Core) SetLogger(logger glog.ILogger) {
	c.logger = logger
}

// GetLogger returns the (logger) of the orm.
func (c *Core) GetLogger() glog.ILogger {
	return c.logger
}

// SetMaxIdleConnCount sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
func (c *Core) SetMaxIdleConnCount(n int) {
	c.dynamicConfig.MaxIdleConnCount = n
}

// SetMaxOpenConnCount sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func (c *Core) SetMaxOpenConnCount(n int) {
	c.dynamicConfig.MaxOpenConnCount = n
}

// SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are not closed due to a connection's age.
func (c *Core) SetMaxConnLifeTime(d time.Duration) {
	c.dynamicConfig.MaxConnLifeTime = d
}

// GetConfig returns the current used node configuration.
func (c *Core) GetConfig() *ConfigNode {
	internalData := c.GetInternalCtxDataFromCtx(c.db.GetCtx())
	if internalData != nil && internalData.ConfigNode != nil {
		// Note:
		// It so here checks and returns the config from current DB,
		// if different schemas between current DB and config.Dbname from context,
		// for example, in nested transaction scenario, the context is passed all through the logic procedure,
		// but the config.Dbname from context may be still the original one from the first transaction object.
		if c.config.Dbname == internalData.ConfigNode.Dbname {
			return internalData.ConfigNode
		}
	}
	return c.config
}

// SetDebug enables/disables the debug mode.
func (c *Core) SetDebug(debug bool) {
	c.debug.Set(debug)
}

// GetDebug returns the debug value.
func (c *Core) GetDebug() bool {
	return c.debug.Val()
}

// GetCache returns the internal cache object.
func (c *Core) GetCache() *gcache.Cache {
	return c.cache
}

// GetGroup returns the group string configured.
func (c *Core) GetGroup() string {
	return c.group
}

// SetDryRun enables/disables the DryRun feature.
func (c *Core) SetDryRun(enabled bool) {
	c.config.DryRun = enabled
}

// GetDryRun returns the DryRun value.
func (c *Core) GetDryRun() bool {
	return c.config.DryRun || allDryRun
}

// GetPrefix returns the table prefix string configured.
func (c *Core) GetPrefix() string {
	return c.config.Prefix
}

// GetSchema returns the schema configured.
func (c *Core) GetSchema() string {
	schema := c.schema
	if schema == "" {
		schema = c.db.GetConfig().Dbname
	}
	return schema
}

func parseConfigNodeLink(node *ConfigNode) *ConfigNode {
	var match []string
	if node.Link != "" {
		match, _ = gregex.MatchString(linkPattern, node.Link)
		if len(match) > 5 {
			node.Type = match[1]
			node.Username = match[2]
			node.Password = match[3]
			node.Protocol = match[4]
			array := gstr.Split(match[5], ":")
			if len(array) == 2 && node.Protocol != "file" {
				node.Hostname = array[0]
				node.Hostport = array[1]
				node.Dbname = match[6]
			} else {
				node.Dbname = match[5]
			}
			if len(match) > 6 && match[7] != "" {
				node.Extra = match[7]
			}
			node.Link = ""
		}
	}
	if node.Extra != "" {
		if m, _ := gstr.Parse(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node
}
