package conf

import (
	"goUtils/env"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DbHost             string
	DbUser             string
	DbPassword         string
	DbDatabase         string
	DbLog              bool
	DbInit             bool
	ListenAddr         string
	BaseUrl            string
	Version            string
	OpenIdURL          string
	OpenIdClientID     string
	OpenIdClientSecret string
	OpenIdLogoutPath   string
	S3Config           S3Config
}

type S3Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Region    string
	Bucket    string
}

func (c *Config) Load() {
	//c.DbHost = env.GetEnvOrElse("DB_HOST", "postgres.agritracking.svc.cluster.local:5432")
	c.DbHost = env.GetEnvOrElse("DB_HOST", "10.111.211.46:5432")
	c.DbUser = env.GetEnvOrElse("DB_USER", "agritracking")
	c.DbPassword = env.GetEnvOrElse("DB_PASSWORD", "Aj9Owie4keexoot8")
	c.DbDatabase = env.GetEnvOrElse("DB_DATABASE", "agritracking")
	dbLog, err := env.GetBoolEnvOrElse("DB_LOG", false)
	if err != nil {
		log.Fatal(err)
	}
	c.DbLog = dbLog
	dbInit, err := env.GetBoolEnvOrElse("DB_INIT", true)
	if err != nil {
		log.Fatal(err)
	}
	c.DbInit = dbInit
	c.ListenAddr = env.GetEnvOrElse("LISTEN_ADDR", "0.0.0.0:4000")
	c.BaseUrl = env.GetEnvOrElse("BASE_URL", "http://adms.bbs.gov.bd/")
	
	c.OpenIdURL = env.GetEnvOrElse("OPENID_URL", "http://123.49.39.107:8080/auth/realms/adms_public")
	//c.OpenIdURL = env.GetEnvOrElse("OPENID_URL", "https://auth.agricensus.org/auth/realms/adms_public")
	c.OpenIdClientID = env.GetEnvOrElse("OPENID_CLIENT_ID", "admspublic")
	c.OpenIdClientSecret = env.GetEnvOrElse("OPENID_CLIENT_SECRET", "367fe26b-ccc6-481f-8573-5a765a44d327")
	//c.OpenIdClientSecret = env.GetEnvOrElse("OPENID_CLIENT_SECRET", "9XLa3ANDYfEnjJ7ytGi87yZaa9AVuheU")
	c.OpenIdLogoutPath = env.GetEnvOrElse("OPENID_LOGOUT_PATH", "auth/realms/adms_public/protocol/openid-connect/logout")

	c.S3Config.Endpoint = env.GetEnvOrElse("S3_ENDPOINT", "minio.solutions.im:10443")
	//c.S3Config.Endpoint = env.GetEnvOrElse("S3_ENDPOINT", "minio.bbs.gov.bd:9443")
	//c.S3Config.Endpoint = env.GetEnvOrElse("S3_ENDPOINT", "10.102.204.35:9443")
	//c.S3Config.AccessKey = env.GetEnvOrElse("S3_ACCESS_KEY", "ANGOK5IO4AHK1PITH7DU")
	c.S3Config.AccessKey = env.GetEnvOrElse("S3_ACCESS_KEY", "U65Z81EH9S39NTNYZ71U")
	c.S3Config.SecretKey = env.GetEnvOrElse("S3_SECRET_KEY", "tOuTEgkp3gkAWsf3acExkxpBn+EYLbfJxJmHephF")
	//c.S3Config.SecretKey = env.GetEnvOrElse("S3_SECRET_KEY", "acieshohz7aecee9cotiephai7eiPeihaor4zooc")
	c.S3Config.Region = env.GetEnvOrElse("S3_REGION", "us-east-1")
	c.S3Config.Bucket = env.GetEnvOrElse("S3_BUCKET", "reports")
	ssl, err := env.GetBoolEnvOrElse("S3_SSL", true)
	if err != nil {
		log.Fatal(err)
	}
	c.S3Config.UseSSL = ssl
	setupLogger()

	log.Infof(`
starting server with the following configuration :
- Database Host : %s
- Database Name: %s
- Database User : %s
- Listen to : %s
- Base URL : %s`, c.DbHost, c.DbDatabase, c.DbUser, c.ListenAddr, c.BaseUrl)
}

func setupLogger() {
	log.SetFormatter(&nested.Formatter{
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "",
		HideKeys:        true,
		NoColors:        false,
		NoFieldsColors:  false,
		ShowFullLevel:   true,
		TrimMessages:    false,
	})
	log.SetLevel(log.DebugLevel)
}
