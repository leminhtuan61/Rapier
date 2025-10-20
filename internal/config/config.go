package config

var Protocol = "tcp"
var Port = ":3000"
var MaxConnection = 20000
var MaxKeyNumber int = 1000000
var EvictionRatio = 0.1

var EvictionPolicy string = "allkeys-lru"

var EpoolMaxSize = 16
var EpoolLruSampleSize = 5

var ListenerNumber int = 2
