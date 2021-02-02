package cassandra

import "github.com/gocql/gocql"

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to cluster
	cluster = gocql.NewCluster("192.168.0.40") //, "192.168.0.41")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

// GetSession func
func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
