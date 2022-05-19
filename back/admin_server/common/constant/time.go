package constant

const (
	DateTimeLayout             = "2006-01-02 15:04:05"
	DateLayout                 = "2006-01-02"
	DateLayoutStr              = "20060102"
	TimeLayout                 = "15:04:05"
	DateTimeWithTimeZoneLayout = "2006-01-02 15:04:05-07:00"

	TimeSecond     = 1
	TimeFiveSecond = 5 * TimeSecond
	TimeMinute     = 60 * TimeSecond
	TimeHour       = 60 * TimeMinute
	TimeDay        = 24 * TimeHour
	TimeWeek       = 7 * TimeDay
	Time30Days     = 30 * TimeDay
)
