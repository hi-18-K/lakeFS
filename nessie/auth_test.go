package nessie

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/treeverse/lakefs/auth"
	"github.com/treeverse/lakefs/auth/crypt"
	"github.com/treeverse/lakefs/auth/model"
	"github.com/treeverse/lakefs/config"
	"github.com/treeverse/lakefs/db"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/viper"
	genclient "github.com/treeverse/lakefs/api/gen/client"
)

const (
	accessKeyID     = "Sneakers"
	secretAccessKey = "Setec Astronomy"
)

func TestSuperuserWithPassedCreds(t *testing.T) {
	ctx, log, _ := setupTest(t)

	cfg := config.NewConfig()
	dbParams := cfg.GetDatabaseParams()
	log.WithField("db_config", fmt.Sprintf("%q", dbParams)).Info("connect to DB")
	pool := db.BuildDatabaseConnection(dbParams)

	viper.SetEnvPrefix("LAKEFS") // Fetch lakeFS envariables in config
	authService := auth.NewDBAuthService(
		pool,
		crypt.NewSecretStore(cfg.GetAuthEncryptionSecret()),
		cfg.GetAuthCacheConfig())

	_, err := auth.SetupAdminUser(authService, &model.SuperuserConfiguration{
		User: model.User{
			Username: "cosmo",
		},
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	})
	if err != nil {
		t.Fatal("failed to setup admin user: ", err)
	}

	// Set up the client to use this authentication.
	endpointURL := viper.GetString("endpoint_url")
	u, err := url.Parse(endpointURL)
	if err != nil {
		t.Fatalf("Failed to parse endpoint URL %s: %s", endpointURL, err)
	}
	apiBasePath := genclient.DefaultBasePath
	if u.Path != "" {
		apiBasePath = u.Path
	}
	r := httptransport.New(u.Host, apiBasePath, []string{u.Scheme})
	r.DefaultAuthentication = httptransport.BasicAuth(accessKeyID, secretAccessKey)
	client.Transport = r

	// Use it for some minimal test
	listRepositories(t, ctx)
}
