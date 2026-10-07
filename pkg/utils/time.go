package utils

import (
	"time"
)

// Daftar layout waktu yang umum digunakan
const (
	LayoutDateTime     = "2006-01-02 15:04:05"
	LayoutDateTimeZone = "2006-01-02 15:04:05 MST"
	LayoutDateOnly     = "2006-01-02"
	LayoutTimeOnly     = "15:04:05"
	LayoutReadable     = "02 January 2006 15:04"
)

// FormatToJakarta mengubah objek time.Time ke zona waktu Asia/Jakarta dengan layout kustom
func FormatToJakarta(t time.Time, layout string) string {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// Fallback menggunakan UTC+7 manual jika tzdata internal sistem tidak tersedia
		loc = time.FixedZone("WIB", 7*60*60)
	}

	return t.In(loc).Format(layout)
}

// FormatNowToJakarta mengambil waktu saat ini dan langsung memformatnya ke zona Asia/Jakarta
func FormatNowToJakarta(layout string) string {
	return FormatToJakarta(time.Now(), layout)
}

// ParseJakartaToTime melakukan parsing string waktu lokal Jakarta kembali menjadi objek time.Time
func ParseJakartaToTime(timeStr, layout string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("WIB", 7*60*60)
	}

	return time.ParseInLocation(layout, timeStr, loc)
}
