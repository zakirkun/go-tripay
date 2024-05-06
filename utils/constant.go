package utils

type TRIPAY_URL string
type TRIPAY_MODE string

const URL_DEVELOPMENT TRIPAY_URL = "https://tripay.co.id/api-sandbox/"
const URL_PRODUCTION TRIPAY_URL = "https://tripay.co.id/api/"

const MODE_DEVELOPMENT TRIPAY_MODE = "development"
const MODE_PRODUCTION TRIPAY_MODE = "production"
