package connectivity

import (
	"fmt"
	"os"
	"path/filepath"
)

type param struct {
	typeid, assetid        int
	year, month, yearmonth string
}

func createFolder(base, year, month string) error {
	path := filepath.Join(base, year, month)
	err := os.MkdirAll(path, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func getFilename(path string, p param) string {
	file := fmt.Sprintf("%s%s.parquet", p.typeid, p.assetid)
	return filepath.Join(path, file)
}

func (p *param) getQuery() string {
	return fmt.Sprintf(`
		select
			client_id, asset_id, year_month, type_id, time, nats_time, payload, worker_time
		from
			connectivity.data
		where
			year_month = '%s/%s' and
			asset_id = %d and
			type_id = %d
	`, p.year, p.month, p.assetid, p.typeid)
}

func (c *Connectivity) getWorkSet(assets []assetym) (params []param) {
	for _, typeid := range c.typeids {
		for _, asset := range assets {
			p := param{
				typeid:    typeid,
				assetid:   asset.assetid,
				year:      asset.year,
				month:     asset.month,
				yearmonth: asset.yearmonth,
			}
			params = append(params, p)
		}
	}
	return params
}

func toChunks(params []param, chunkcount int) (chunks [][]param) {
	arraysize := len(params)
	chunksize := int(arraysize / chunkcount)

	for i := 0; i < chunkcount; i++ {
		var chunk []param
		if i != chunkcount-1 {
			chunk = params[i*chunksize : (i+1)*chunksize]
		} else {
			chunk = params[i*chunksize : arraysize]
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}
