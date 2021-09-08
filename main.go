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
	Rat