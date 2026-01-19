package modules

import (
	"fmt"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Helper function to extract integer value from Integer or Instance
func getIntegerValue(obj object.Object) (int64, bool) {
	switch val := obj.(type) {
	case *object.Integer:
		return val.Value, true
	case *object.Instance:
		if value, exists := val.Env.Get("value"); exists {
			if intVal, ok := value.(*object.Integer); ok {
				return intVal.Value, true
			}
		}
	}
	return 0, false
}

// Helper function to extract float value from Float or Instance
func getFloatValue(obj object.Object) (float64, bool) {
	switch val := obj.(type) {
	case *object.Float:
		return val.Value, true
	case *object.Instance:
		if value, exists := val.Env.Get("value"); exists {
			if floatVal, ok := value.(*object.Float); ok {
				return floatVal.Value, true
			}
		}
	}
	return 0, false
}

var TimeModule = map[string]*object.Builtin{
	// Get current time as Time object
	"now": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "now() takes no arguments"}
			}
			return &object.Time{Value: time.Now()}
		},
	},

	// Get current time as Unix timestamp (seconds)
	"timeNow": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "timeNow takes no arguments"}
			}
			return &object.Integer{Value: time.Now().Unix()}
		},
	},

	// Get current time as Unix timestamp (nanoseconds)
	"timeNowNano": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "timeNowNano takes no arguments"}
			}
			return &object.Integer{Value: time.Now().UnixNano()}
		},
	},

	// Sleep for specified duration
	"timeSleep": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "timeSleep requires 1 argument: seconds (INT or FLOAT)"}
			}

			// Try to get integer value first
			if intVal, ok := getIntegerValue(args[0]); ok {
				if intVal < 0 {
					return &object.Error{Message: "timeSleep duration cannot be negative"}
				}
				time.Sleep(time.Duration(intVal) * time.Second)
				return &object.None{}
			}

			// Try to get float value
			if floatVal, ok := getFloatValue(args[0]); ok {
				if floatVal < 0 {
					return &object.Error{Message: "timeSleep duration cannot be negative"}
				}
				nanos := int64(floatVal * 1_000_000_000)
				time.Sleep(time.Duration(nanos))
				return &object.None{}
			}

			// Check for Duration type
			if durVal, ok := args[0].(*object.Duration); ok {
				if durVal.Value < 0 {
					return &object.Error{Message: "timeSleep duration cannot be negative"}
				}
				time.Sleep(durVal.Value)
				return &object.None{}
			}

			return &object.Error{Message: "timeSleep argument must be INTEGER, FLOAT, or DURATION, got " + string(args[0].Type())}
		},
	},

	// Parse time from string
	"parseTime": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 && len(args) != 2 {
				return &object.Error{Message: "parseTime() takes 1 or 2 arguments (timeString, [format])"}
			}

			timeStr, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "First argument must be a STRING"}
			}

			format := time.RFC3339
			if len(args) == 2 {
				formatStr, ok := args[1].(*object.String)
				if !ok {
					return &object.Error{Message: "Second argument must be a STRING"}
				}
				format = formatStr.Value
			}

			parsedTime, err := time.Parse(format, timeStr.Value)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to parse time: %s", err)}
			}

			return &object.Time{Value: parsedTime}
		},
	},

	// Parse time and return Unix timestamp
	"timeParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeParse requires 2 arguments: format, timeString"}
			}

			format, ok1 := args[0].(*object.String)
			timeStr, ok2 := args[1].(*object.String)
			if !ok1 || !ok2 {
				return &object.Error{Message: "timeParse arguments must be STRINGs"}
			}

			t, err := time.Parse(format.Value, timeStr.Value)
			if err != nil {
				return &object.Error{Message: "timeParse failed: " + err.Error()}
			}

			return &object.Integer{Value: t.Unix()}
		},
	},

	// Format time
	"formatTime": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "formatTime() requires exactly 2 arguments"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			formatStr, ok := args[1].(*object.String)
			if !ok {
				return &object.Error{Message: "Second argument must be a STRING"}
			}

			formatted := timeObj.Value.Format(formatStr.Value)
			return &object.String{Value: formatted}
		},
	},

	// Format Unix timestamp
	"timeFormat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "timeFormat requires 1 or 2 arguments: timestamp, [format]"}
			}

			timestamp, ok := getIntegerValue(args[0])
			if !ok {
				return &object.Error{Message: "timeFormat first argument must be INTEGER timestamp"}
			}

			format := "2006-01-02 15:04:05"
			if len(args) == 2 {
				formatArg, ok := args[1].(*object.String)
				if !ok {
					return &object.Error{Message: "timeFormat second argument must be STRING format"}
				}
				format = formatArg.Value
			}

			t := time.Unix(timestamp, 0)
			formatted := t.Format(format)
			return &object.String{Value: formatted}
		},
	},

	// Time since a given time
	"timeSince": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "timeSince() requires exactly 1 argument"}
			}

			// Check if it's a Time object first
			if t, ok := args[0].(*object.Time); ok {
				elapsed := time.Since(t.Value)
				return &object.Duration{Value: elapsed}
			}

			// Try to get integer value (Unix timestamp)
			if timestamp, ok := getIntegerValue(args[0]); ok {
				elapsed := time.Since(time.Unix(timestamp, 0))
				return &object.Duration{Value: elapsed}
			}

			return &object.Error{Message: "Argument must be of type TIME or INTEGER (Unix timestamp)"}
		},
	},

	// Time until a given time
	"timeUntil": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "timeUntil() requires exactly 1 argument"}
			}

			// Check if it's a Time object first
			if t, ok := args[0].(*object.Time); ok {
				until := time.Until(t.Value)
				return &object.Duration{Value: until}
			}

			// Try to get integer value (Unix timestamp)
			if timestamp, ok := getIntegerValue(args[0]); ok {
				until := time.Until(time.Unix(timestamp, 0))
				return &object.Duration{Value: until}
			}

			return &object.Error{Message: "Argument must be of type TIME or INTEGER (Unix timestamp)"}
		},
	},

	// Add duration to time
	"addDuration": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "addDuration() requires exactly 2 arguments"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			durObj, ok := args[1].(*object.Duration)
			if !ok {
				return &object.Error{Message: "Second argument must be of type DURATION"}
			}

			newTime := timeObj.Value.Add(durObj.Value)
			return &object.Time{Value: newTime}
		},
	},

	// Add duration to Unix timestamp
	"timeAddDuration": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeAddDuration requires 2 arguments: timestamp, seconds"}
			}

			timestamp, ok1 := getIntegerValue(args[0])
			duration, ok2 := getIntegerValue(args[1])
			if !ok1 || !ok2 {
				return &object.Error{Message: "timeAddDuration arguments must be INTEGERs"}
			}

			t := time.Unix(timestamp, 0)
			newTime := t.Add(time.Duration(duration) * time.Second)
			return &object.Integer{Value: newTime.Unix()}
		},
	},

	// Calculate difference between timestamps
	"timeDiff": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeDiff requires 2 arguments"}
			}

			// Check if both are Time objects
			if t1, ok1 := args[0].(*object.Time); ok1 {
				if t2, ok2 := args[1].(*object.Time); ok2 {
					diff := t1.Value.Sub(t2.Value)
					return &object.Duration{Value: diff}
				}
				return &object.Error{Message: "Both arguments must be of the same type"}
			}

			// Try to get both as integers
			if val1, ok1 := getIntegerValue(args[0]); ok1 {
				if val2, ok2 := getIntegerValue(args[1]); ok2 {
					diff := val1 - val2
					return &object.Integer{Value: diff}
				}
				return &object.Error{Message: "Both arguments must be of the same type"}
			}

			return &object.Error{Message: "Arguments must be INTEGERs or TIMEs"}
		},
	},

	// Get date components
	"timeDate": {
		Fn: func(args ...object.Object) object.Object {
			var t time.Time
			if len(args) == 0 {
				t = time.Now()
			} else if len(args) == 1 {
				// Check if it's a Time object
				if timeObj, ok := args[0].(*object.Time); ok {
					t = timeObj.Value
				} else if intVal, ok := getIntegerValue(args[0]); ok {
					// Try to get integer value (Unix timestamp)
					t = time.Unix(intVal, 0)
				} else {
					return &object.Error{Message: "timeDate argument must be INTEGER timestamp or TIME"}
				}
			} else {
				return &object.Error{Message: "timeDate requires 0 or 1 arguments"}
			}

			year, month, day := t.Date()
			elements := []object.Object{
				&object.Integer{Value: int64(year)},
				&object.Integer{Value: int64(month)},
				&object.Integer{Value: int64(day)},
			}
			return &object.Array{Elements: elements}
		},
	},

	// Create duration from various units
	"seconds": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "seconds() requires exactly 1 argument"}
			}

			// Try to get integer value
			if intVal, ok := getIntegerValue(args[0]); ok {
				return &object.Duration{Value: time.Duration(intVal) * time.Second}
			}

			// Try to get float value
			if floatVal, ok := getFloatValue(args[0]); ok {
				nanos := int64(floatVal * float64(time.Second))
				return &object.Duration{Value: time.Duration(nanos)}
			}

			return &object.Error{Message: "Argument must be an INTEGER or FLOAT"}
		},
	},

	"minutes": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "minutes() requires exactly 1 argument"}
			}

			// Try to get integer value
			if intVal, ok := getIntegerValue(args[0]); ok {
				return &object.Duration{Value: time.Duration(intVal) * time.Minute}
			}

			// Try to get float value
			if floatVal, ok := getFloatValue(args[0]); ok {
				nanos := int64(floatVal * float64(time.Minute))
				return &object.Duration{Value: time.Duration(nanos)}
			}

			return &object.Error{Message: "Argument must be an INTEGER or FLOAT"}
		},
	},

	"hours": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "hours() requires exactly 1 argument"}
			}

			// Try to get integer value
			if intVal, ok := getIntegerValue(args[0]); ok {
				return &object.Duration{Value: time.Duration(intVal) * time.Hour}
			}

			// Try to get float value
			if floatVal, ok := getFloatValue(args[0]); ok {
				nanos := int64(floatVal * float64(time.Hour))
				return &object.Duration{Value: time.Duration(nanos)}
			}

			return &object.Error{Message: "Argument must be an INTEGER or FLOAT"}
		},
	},

	"milliseconds": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "milliseconds() requires exactly 1 argument"}
			}

			// Try to get integer value
			if intVal, ok := getIntegerValue(args[0]); ok {
				return &object.Duration{Value: time.Duration(intVal) * time.Millisecond}
			}

			// Try to get float value
			if floatVal, ok := getFloatValue(args[0]); ok {
				nanos := int64(floatVal * float64(time.Millisecond))
				return &object.Duration{Value: time.Duration(nanos)}
			}

			return &object.Error{Message: "Argument must be an INTEGER or FLOAT"}
		},
	},

	// Duration operations
	"durationToSeconds": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "durationToSeconds() requires exactly 1 argument"}
			}

			durObj, ok := args[0].(*object.Duration)
			if !ok {
				return &object.Error{Message: "Argument must be of type DURATION"}
			}

			return &object.Float{Value: durObj.Value.Seconds()}
		},
	},

	"durationToMinutes": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "durationToMinutes() requires exactly 1 argument"}
			}

			durObj, ok := args[0].(*object.Duration)
			if !ok {
				return &object.Error{Message: "Argument must be of type DURATION"}
			}

			return &object.Float{Value: durObj.Value.Minutes()}
		},
	},

	"durationToHours": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "durationToHours() requires exactly 1 argument"}
			}

			durObj, ok := args[0].(*object.Duration)
			if !ok {
				return &object.Error{Message: "Argument must be of type DURATION"}
			}

			return &object.Float{Value: durObj.Value.Hours()}
		},
	},

	"durationToMilliseconds": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "durationToMilliseconds() requires exactly 1 argument"}
			}

			durObj, ok := args[0].(*object.Duration)
			if !ok {
				return &object.Error{Message: "Argument must be of type DURATION"}
			}

			return &object.Integer{Value: durObj.Value.Milliseconds()}
		},
	},

	// Compare times
	"timeBefore": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeBefore() requires exactly 2 arguments"}
			}

			time1, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			time2, ok := args[1].(*object.Time)
			if !ok {
				return &object.Error{Message: "Second argument must be of type TIME"}
			}

			return &object.Boolean{Value: time1.Value.Before(time2.Value)}
		},
	},

	"timeAfter": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeAfter() requires exactly 2 arguments"}
			}

			time1, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			time2, ok := args[1].(*object.Time)
			if !ok {
				return &object.Error{Message: "Second argument must be of type TIME"}
			}

			return &object.Boolean{Value: time1.Value.After(time2.Value)}
		},
	},

	"timeEqual": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeEqual() requires exactly 2 arguments"}
			}

			time1, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			time2, ok := args[1].(*object.Time)
			if !ok {
				return &object.Error{Message: "Second argument must be of type TIME"}
			}

			return &object.Boolean{Value: time1.Value.Equal(time2.Value)}
		},
	},

	// Get time components
	"year": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "year() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Year())}
		},
	},

	"month": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "month() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Month())}
		},
	},

	"day": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "day() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Day())}
		},
	},

	"weekday": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "weekday() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Weekday())}
		},
	},

	"hour": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "hour() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Hour())}
		},
	},

	"minute": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "minute() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Minute())}
		},
	},

	"second": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "second() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: int64(timeObj.Value.Second())}
		},
	},

	// Unix timestamp operations
	"unix": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "unix() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: timeObj.Value.Unix()}
		},
	},

	"unixNano": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "unixNano() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Integer{Value: timeObj.Value.UnixNano()}
		},
	},

	"fromUnix": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fromUnix() requires exactly 1 argument"}
			}

			timestamp, ok := getIntegerValue(args[0])
			if !ok {
				return &object.Error{Message: "Argument must be an INTEGER"}
			}

			return &object.Time{Value: time.Unix(timestamp, 0)}
		},
	},

	"fromUnixNano": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fromUnixNano() requires exactly 1 argument"}
			}

			timestamp, ok := getIntegerValue(args[0])
			if !ok {
				return &object.Error{Message: "Argument must be an INTEGER"}
			}

			return &object.Time{Value: time.Unix(0, timestamp)}
		},
	},

	// Timezone operations
	"inLocation": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "inLocation() requires exactly 2 arguments"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "First argument must be of type TIME"}
			}

			locStr, ok := args[1].(*object.String)
			if !ok {
				return &object.Error{Message: "Second argument must be a STRING"}
			}

			loc, err := time.LoadLocation(locStr.Value)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to load location: %s", err)}
			}

			return &object.Time{Value: timeObj.Value.In(loc)}
		},
	},

	"utc": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "utc() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Time{Value: timeObj.Value.UTC()}
		},
	},

	"local": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "local() requires exactly 1 argument"}
			}

			timeObj, ok := args[0].(*object.Time)
			if !ok {
				return &object.Error{Message: "Argument must be of type TIME"}
			}

			return &object.Time{Value: timeObj.Value.Local()}
		},
	},
}
