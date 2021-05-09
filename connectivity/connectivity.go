package connectivity

import (
	"context"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/bluecolor/cazcopy/cassandra"
	azb "github.com/xitongsys/parquet-go-source/azblob"
	"github.com/xitongsys/parquet-go/writer"
)

type Connectivity struct {
	assets  string
	startym string
	endym   string
	typeids []int
}

type ConnectivityRecord struct {
	ClientId   int32  `parquet:"name=client_id, type=INT32"`
	AssetId    int32  `parquet:"name=asset_id, type=INT32"`
	YearMonth  string `parquet:"name=year_month, type=BYTE_ARRAY"`
	TypeId     int32  `parquet:"name=type_id, type=INT32"`
	Time       int32  `parquet:"name=time, type=INT32, convertedtype=DATE"`
	NatsTime   int32  `parquet:"name=nats_time, type=INT32, convertedtype=DATE"`
	Payload    string `parquet:"name=payload, type=BYTE_ARRAY"`
	WorkerTime int32  `parquet:"name=worker_time, type=INT32, convertedtype=DATE"`
}

func NewConnectivityRecord(r map[string]interface{}) ConnectivityRecord {
	return ConnectivityRecord{
		ClientId:   r["client_id"].(int32),
		AssetId:    r["asset_id"].(int32),
		YearMonth:  r["year_month"].(string),
		TypeId:     r["type_id"].(int32),
		Time:       r["time"].(int32),
		NatsTime:   r["nats_time"].(int32),
		Payload:    r["payload"].(string),
		WorkerTime: r["worker_time"].(int32),
	}
}

func NewConnectivity(assets, startym, endym string, typeids []int) *Connectivity {
	return &Connectivity{
		assets: assets, startym: startym, endym: endym, typeids: typeids,
	}
}

func startWorker(cassandra *cassandra.Cassandra, params []param, credential azblob.Credential) error {
	_, err := cassandra.Connect()
	ctx := context.Background()
	if err != nil {
		return err
	}

	for _, p := range params {
		if err != nil {
			return err
		}
		url := ""
		options := azb.WriterOptions{}
		azw, err := azb.NewAzBlobFileWriter(ctx, url, credential, options)
		pw, err := writer.NewParquetWriter(fw, new(student), 4)
		if err != nil {
			return err
		}

		query := p.getQuery()
		iter := cassandra.Session.Query(query).Iter()
		// columns := iter.Columns()
		scanner := iter.Scanner()
		var buffer []ConnectivityRecord
		for scanner.Next() {
			var record map[string]interface{}

			if err := scanner.Scan(&record); err != nil {
				return err
			}
			cr := NewConnectivityRecord(record)
			if err = azw.Write(cr); err != nil {
				return err
			}
		}

	}

	return nil
}

func (c *Connectivity) Run(cassandra *cassandra.Cassandra, tmpPath string, parallel int) error {
	assets, err := readAssets(c.assets, c.startym, c.endym)
	if err != nil {
		return err
	}
	params := c.getWorkSet(assets)
	chunks := toChunks(params, parallel)

	for _, chunk := range chunks {
		startWorker(cassandra, chunk)
	}
	return nil
}
