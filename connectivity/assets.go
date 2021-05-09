package connectivity

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type assetym struct {
	assetid                int
	year, month, yearmonth string
}

func readAssets(path string, args ...string) (assets []assetym, err error) {
	var (
		startym string
		endym   string
	)
	if len(args) > 0 {
		startym = args[0]
	}
	if len(args) > 1 {
		endym = args[1]
	}
	if endym < startym {
		return nil, errors.New("Start year month is greater than end year month!")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		assetid, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		yearmonth := fmt.Sprintf("%s%s", line[2], line[1])

		if startym != "" && yearmonth < startym {
			continue
		}
		if endym != "" && yearmonth > endym {
			continue
		}

		asset := assetym{
			assetid:   assetid,
			year:      line[2],
			month:     line[1],
			yearmonth: yearmonth,
		}
		assets = append(assets, asset)
	}

	return assets, err
}
