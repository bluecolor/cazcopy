package connectivity

import "fmt"

type Connectivity struct {
	assets  string
	startym string
	endym   string
	typeids []int
}

type param struct {
	typeid, assetid        int
	year, month, yearmonth string
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

func NewConnectivity(assets, startym, endym string, typeids []int) *Connectivity {
	return &Connectivity{
		assets: assets, startym: startym, endym: endym, typeids: typeids,
	}
}

func (c *Connectivity) Run(tmpPath string, parallel int) error {
	assets, err := readAssets(c.assets, c.startym, c.endym)
	if err != nil {
		return err
	}
	return nil
}
