package config

import (
	"os"
	"time"

	"bitbucket.org/optiisolutions/go-common-v2/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	ServiceName            = "timeline-svc"
	GracefulStopTimeoutSec = 5
	APIBasePath            = "/api"

	DBMigrationsBaseDir  = "internal/migrations/base_migrations"
	DBMigrationsShardDir = "internal/migrations/shard_migrations"

	PubSubRetryTime  = 500 * time.Millisecond
	PubSubMaxRetries = 3
)

type (
	Config struct {
		config.ServiceParams
		config.TlsValidatorParams
		config.GCP
		DbParams
		ServicesURLs

		ConcurrentSeedingNum string `json:"concurrent_seeding_num"`

		ZipkinURL string `json:"zipkin__endpoint"`
	}
	DbParams struct {
		DbMigrationPath  string `json:"db_migration_path"`
		Host             string `json:"db_host"`
		Port             string `json:"db_port"`
		User             string `json:"db_user"`
		Password         string `json:"db_password"`
		Database         string `json:"db_database"`
		NumConns         string `json:"db_conns"`
		DebugLogs        bool   `json:"db_debug_logs"`
		IsTLSEnabled     string `json:"postgres_tls_enabled"`
		SqlTLSRootCert   string `json:"PGSSLCERT"`
		SqlTLSClientCert string `json:"postgres_tls_client_cert"`
		SqlTLSClientKey  string `json:"postgres_tls_client_key"`
	}
	ServicesURLs struct {
		AdhocJobMgrURL   string `json:"adhoc_job_mgr_url"`
		DepartmentMgrURL string `json:"department_mgr_url"`
		EmployeeMgrURL   string `json:"employee_mgr_url"`
		LocationsMgrURL  string `json:"locations_mgr_url"`
		PropertyMgrURL   string `json:"property_mgr_url"`
	}
)

var cfg Config

const defaultDatabaseName = "timeline-svc"

func InitConfig(commit, builtAt string) *Config {
	//  load space .env variables first if available
	filename := "./files/.env"
	if _, err := os.Stat(filename); err == nil {
		_ = godotenv.Load(filename)
	}

	cfg := &Config{
		ServiceParams: config.ServiceParams{Environment: "local",
			Host:     "",
			BaseUri:  "",
			Port:     "8099",
			LogLevel: "debug",
		},
		DbParams: DbParams{
			Host:     "",
			Port:     "5432",
			User:     "",
			Password: "",
			Database: defaultDatabaseName,
			NumConns: "3",
		},
	}

	err := cfg.LoadEnvVariables(cfg, commit, builtAt)
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	err = cfg.ConfigureLogger()
	if err != nil {
		logrus.Fatalf("cannot ConfigureLogger: %s", err.Error())
	}
	SetCurrentCfg(*cfg)

	return cfg
}

func GetCurrentCfg() Config {
	return cfg
}

func SetCurrentCfg(c Config) {
	cfg = c
}
