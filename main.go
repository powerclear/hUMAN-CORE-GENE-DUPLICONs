package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	"encoding/json"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Config struct {
	CryptoStock             []string `yaml:"CRYPTO_STOCK"`
	CoinApiKey              string   `yaml:"COIN_API_KEY"`
	CoinApiExchangeCurrency string   `yaml:"COIN_API_EXCHANGE_CURRENCY"`
}

type SpecificRate struct {
	Time string
	Rate float64
}

type TimeSeriesData struct {
	TimePeriodStart string
	TimePeriodEnd   string
	RateClose       float64
}

func (c *Config) getConfig() *Config {

	os.Setenv("TERMINAL_CHECK_MARKET_CONFIG_PATH", "config.yaml")

	configFilePath := os.Getenv("TERMINAL_CHECK_MARKET_CONFIG_PATH")
	if len(configFilePath) == 0 {
		os.Exit(1)
	}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error in getting yaml file path", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("Error in reading yaml file")
		os.Exit(1)
	}

	return c
}

func getCurrentPrice(config Config, crypto string) (*SpecificRate, error) {
	client := &http.Client{}

	url := "https://rest.coinapi.io/v1/