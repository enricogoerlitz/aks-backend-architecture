package constants

import "os"

var ENV_MODE = os.Getenv("MODE")
var ENV_DB_TYPE = os.Getenv("DB_TYPE")
var ENV_DB_WRITE_HOST = os.Getenv("DB_WRITE_HOST")
var ENV_DB_WRITE_PORT = os.Getenv("DB_WRITE_PORT")
var ENV_DB_WRITE_USER = os.Getenv("DB_WRITE_USER")
var ENV_DB_WRITE_PASSWORD = os.Getenv("DB_WRITE_PASSWORD")
var ENV_DB_READ_HOST = os.Getenv("DB_READ_HOST")
var ENV_DB_READ_PORT = os.Getenv("DB_READ_PORT")
var ENV_DB_READ_USER = os.Getenv("DB_READ_USER")
var ENV_DB_READ_PASSWORD = os.Getenv("DB_READ_PASSWORD")
var ENV_DB_NAME = os.Getenv("DB_NAME")
var ENV_CACHE_TYPE = os.Getenv("CACHE_TYPE")
var ENV_CACHE_HOSTS = os.Getenv("CACHE_HOSTS")
var ENV_CACHE_PORT = os.Getenv("CACHE_PORT")
var ENV_CACHE_USER = os.Getenv("CACHE_USER")
var ENV_CACHE_PASSWORD = os.Getenv("CACHE_PASSWORD")
