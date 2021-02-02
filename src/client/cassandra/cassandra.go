package cassandra

import "github.com/gocql/gocql"

var (
	session *gocql.Session
)

func init() {
	// connect to cluster
	cluster := gocql.NewCluster("192.168.0.40") //, "192.168.0.41")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

// GetSession func
func GetSession() *gocql.Session {
	return session
}
