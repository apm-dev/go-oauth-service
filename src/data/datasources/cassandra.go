package datasources

import (
	"github.com/gocql/gocql"
)

var cassandra CassandraDB

type CassandraDB interface {
	GetSession() (*gocql.Session, error)
}

func GetCassandraClient() CassandraDB {
	if cassandra == nil {
		cluster := gocql.NewCluster("127.0.0.1")
		cluster.Keyspace = "oauth"
		cluster.Consistency = gocql.Quorum
		cassandra = &cassandraDB{Cluster: cluster}
	}
	return cassandra
}

type cassandraDB struct {
	Cluster *gocql.ClusterConfig
}

func init() {
	GetCassandraClient()
}

func (c *cassandraDB) GetSession() (*gocql.Session, error) {
	return c.Cluster.CreateSession()
}
