package config

import (
	"fmt"
	"strconv"
)

// DatabaseKind ...
type DatabaseKind string

// Defines
const (
	MongoDBKind     DatabaseKind = "MONGO"
	RedisDBKind     DatabaseKind = "REDIS"
	PostgresDBKind  DatabaseKind = "POSTGRES"
	CassandraDBKind DatabaseKind = "CASSANDRA"
)

// Service ...
type Service struct {
	Name     string   `yaml:"name"`
	HTTP     HTTP     `yaml:"http"`
	GRPC     GRPC     `yaml:"grpc"`
	Mongo    Mongo    `yaml:"mongo"`
	Mail     Mail     `yaml:"mail"`
	Coinbase Coinbase `yaml:"coinbase"`
	CSE      CSE      `yaml:"cse"`
}

// HTTP ...
type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// GRPC ...
type GRPC struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Mail ...
type Mail struct {
	APIKey      string `yaml:"api_key"`
	SenderName  string `yaml:"sendername"`
	SenderEmail string `yaml:"senderemail"`
	Adapter     string `yaml:"adapter"`
}

// Coinbase ...
type Coinbase struct {
	APIKey     string `yaml:"api_key"`
	APISecret  string `yaml:"api_secret"`
	WebhookKey string `yaml:"webhook_key"`
}

// CSE ...
type CSE struct {
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

// IsUseMongo ...
func (t DatabaseKind) IsUseMongo() bool {
	return t == MongoDBKind
}

// IsUseRedis ...
func (t DatabaseKind) IsUseRedis() bool {
	return t == RedisDBKind
}

// IsUseCassandra ...
func (t DatabaseKind) IsUseCassandra() bool {
	return t == CassandraDBKind
}

// IsUsePostgres ...
func (t DatabaseKind) IsUsePostgres() bool {
	return t == PostgresDBKind
}

// Config ...
type Config struct {
	IsProduction bool

	ServiceName string
	GRPC        ConnAddress
	HTTP        ConnAddress

	UserStoreUseDBKind DatabaseKind

	SQL   SQL
	Mongo Mongo

	JaegerAddress string
	ConsulAddress string

	PrometheusPort string

	LogFolderPath string

	CoinbaseService       string
	CryptocurrencyService string
	OrderService          string
	UserService           string
	VerifyCodeService     string
	LendingService        string

	KafkaBrokers string // ex: localhost:27017,localhost:29993,...
}

// ConnAddress ...
type ConnAddress struct {
	Host string
	Port string
}

// String ...
func (c *ConnAddress) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// GetPortInt ...
func (c ConnAddress) GetPortInt() int {
	i2, _ := strconv.ParseInt(c.Port, 10, 64)
	return int(i2)
}

// SQL ...
type SQL struct {
	Namespace   string
	Automigrate bool

	Dialect string

	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	Timeout  int
}

// PostgresConnectionString ...
func (c SQL) PostgresConnectionString() string {
	sslmode := c.SSLMode
	if c.SSLMode == "" {
		sslmode = "disable"
	}
	if c.Timeout == 0 {
		c.Timeout = 30
	}

	s := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%v", c.Username, c.Password, c.Host, c.Port, c.Database, sslmode, c.Timeout)
	return s
}

// Mongo ...
type Mongo struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ConnectionString returns connection string for mongodb
func (c Mongo) ConnectionString() string {
	conn := "mongodb://"
	if c.Username != "" {
		conn += fmt.Sprintf("%s:%s@", c.Username, c.Password)
	}
	conn += fmt.Sprintf("%s:%s/%s", c.Address, c.Port, c.Database)
	return conn
}
