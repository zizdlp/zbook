package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/oschwald/maxminddb-golang"
)

type GeoData struct {
	IPRangeCIDR  string
	CityNameEn   *string
	CityNameZhCn *string
	Latitude     *float64
	Longitude    *float64
}

type CityNames struct {
	En   *string `maxminddb:"en"`
	ZhCN *string `maxminddb:"zh-CN"`
}

type Location struct {
	Latitude  *float64 `maxminddb:"latitude"`
	Longitude *float64 `maxminddb:"longitude"`
}

type GeoRecord struct {
	City struct {
		Names CityNames `maxminddb:"names"`
	} `maxminddb:"city"`
	Location Location `maxminddb:"location"`
}

func processBatch(batchData []GeoData, db *pgxpool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to begin transaction")
		return
	}
	defer tx.Rollback(ctx)

	stmt := `
        INSERT INTO geoip (ip_range_cidr, city_name_en, city_name_zh_cn, latitude, longitude)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (ip_range_cidr) DO NOTHING;
    `

	for _, data := range batchData {
		_, err := tx.Exec(ctx, stmt, data.IPRangeCIDR, data.CityNameEn, data.CityNameZhCn, data.Latitude, data.Longitude)
		if err != nil {
			log.Error().Err(err).Msg("Failed to execute statement")
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction")
		return
	}
}

func main() {
	// Set up logging to output to standard error with console formatting
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if len(os.Args) != 2 {
		log.Error().Msg("Please provide the mmdb file path")
		return
	}
	dbPath := os.Args[1]

	// Record start time
	startTime := time.Now()

	connStr := "user=root password=secret dbname=zbook host=localhost sslmode=disable"
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse database connection string")
		return
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return
	}
	defer db.Close()

	reader, err := maxminddb.Open(dbPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open MMDB file:")
		return
	}
	defer reader.Close()

	batchSize := 1000
	var batchData []GeoData
	var wg sync.WaitGroup

	networkIter := reader.Networks()
	for networkIter.Next() {
		var geoRecord GeoRecord
		network, err := networkIter.Network(&geoRecord)
		if err != nil {
			log.Error().Err(err).Msgf("Error processing network: %v", err)
			continue
		}

		ipRange := network.String()

		cityNameEn := geoRecord.City.Names.En
		cityNameZhCn := geoRecord.City.Names.ZhCN
		latitude := geoRecord.Location.Latitude
		longitude := geoRecord.Location.Longitude

		batchData = append(batchData, GeoData{
			IPRangeCIDR:  ipRange,
			CityNameEn:   cityNameEn,
			CityNameZhCn: cityNameZhCn,
			Latitude:     latitude,
			Longitude:    longitude,
		})

		if len(batchData) >= batchSize {
			wg.Add(1)
			go processBatch(batchData, db, &wg)
			batchData = nil
		}
	}

	if len(batchData) > 0 {
		wg.Add(1)
		go processBatch(batchData, db, &wg)
	}

	wg.Wait()

	// Record end time and calculate duration
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	log.Info().Msgf("Data processing complete. Total time taken: %s", duration)
}
