package modules

import (
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

var TimeModule = map[string]*object.Builtin{
	"timeNow": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "timeNow takes no arguments"}
			}
			return &object.Integer{Value: time.Now().Unix()}
		},
	},
	"timeNowNano": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "timeNowNano takes no arguments"}
			}
			return &object.Integer{Value: time.Now().UnixNano()}
		},
	},
	"timeSleep": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "timeSleep requires 1 argument: seconds (INT or FLOAT)"}
			}

			switch val := args[0].(type) {
			case *object.Integer:
				if val.Value < 0 {
					return &object.Error{Message: "timeSleep duration cannot be negative"}
				}
				time.Sleep(time.Duration(val.Value) * time.Second)
			case *object.Float:
				if val.Value < 0 {
					return &object.Error{Message: "timeSleep duration cannot be negative"}
				}
				nanos := int64(val.Value * 1_000_000_000)
				time.Sleep(time.Duration(nanos))
			case *object.Instance:
				// Handle instance-wrapped values
				if value, exists := val.Env.Get("value"); exists {
					switch innerVal := value.(type) {
					case *object.Integer:
						if innerVal.Value < 0 {
							return &object.Error{Message: "timeSleep duration cannot be negative"}
						}
						time.Sleep(time.Duration(innerVal.Value) * time.Second)
					case *object.Float:
						if innerVal.Value < 0 {
							return &object.Error{Message: "timeSleep duration cannot be negative"}
						}
						nanos := int64(innerVal.Value * 1_000_000_000)
						time.Sleep(time.Duration(nanos))
					default:
						return &object.Error{Message: "timeSleep instance value must be INTEGER or FLOAT, got " + string(innerVal.Type())}
					}
				} else {
					return &object.Error{Message: "timeSleep instance missing value"}
				}
			default:
				return &object.Error{Message: "timeSleep argument must be INTEGER or FLOAT, got " + string(args[0].Type())}
			}

			return &object.None{}
		},
	},
	"timeFormat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "timeFormat requires 1 or 2 arguments: timestamp, [format]"}
			}

			timestamp, ok := args[0].(*object.Integer)
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

			t := time.Unix(timestamp.Value, 0)
			formatted := t.Format(format)
			return &object.String{Value: formatted}
		},
	},
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
	"timeDate": {
		Fn: func(args ...object.Object) object.Object {
			var t time.Time
			if len(args) == 0 {
				t = time.Now()
			} else if len(args) == 1 {
				timestamp, ok := args[0].(*object.Integer)
				if !ok {
					return &object.Error{Message: "timeDate argument must be INTEGER timestamp"}
				}
				t = time.Unix(timestamp.Value, 0)
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
	"timeAddDuration": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeAddDuration requires 2 arguments: timestamp, seconds"}
			}

			timestamp, ok1 := args[0].(*object.Integer)
			duration, ok2 := args[1].(*object.Integer)
			if !ok1 || !ok2 {
				return &object.Error{Message: "timeAddDuration arguments must be INTEGERs"}
			}

			t := time.Unix(timestamp.Value, 0)
			newTime := t.Add(time.Duration(duration.Value) * time.Second)
			return &object.Integer{Value: newTime.Unix()}
		},
	},
	"timeDiff": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "timeDiff requires 2 arguments: timestamp1, timestamp2"}
			}

			ts1, ok1 := args[0].(*object.Integer)
			ts2, ok2 := args[1].(*object.Integer)
			if !ok1 || !ok2 {
				return &object.Error{Message: "timeDiff arguments must be INTEGERs"}
			}

			diff := ts1.Value - ts2.Value
			return &object.Integer{Value: diff}
		},
	},
}

