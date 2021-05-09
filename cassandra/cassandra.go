package cassandra

import (
	"github.com/gocql/gocql"
)

type Cassandra struct {
	Hosts    []string
	Keyspace string

	Session *gocql.Session
}

func NewCassandra(hosts []string, keyspace string) *Cassandra {
	return &Cassandra{
		Hosts: hosts, Keyspace: keyspace,
	}
}

func (c *Cassandra) Connect() (_ *gocql.Session, err error) {
	cluster := gocql.NewCluster(c.Hosts...)
	cluster.Keyspace = c.Keyspace
	c.Session, err = cluster.CreateSession()
	c.Session.SetConsistency(gocql.Consistency(gocql.LocalOne))
	if err != nil {
		return nil, err
	}
	return c.Session, nil
}
