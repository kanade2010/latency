package latency

import (
	"fmt"
	"time"
)

var (
	green            = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white            = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow           = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red              = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue             = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta          = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan             = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset            = string([]byte{27, 91, 48, 109})
)

var (
	tagColor	 = blue
	methodColor  = yellow
	latencyColor = cyan
	resetColor   = reset
)

// Latency is the structure any formatter will be handed when time to log comes
type Latency struct {
	// Latency is how much time the method cost.
	Latency time.Duration
	// Method is method tag.
	Method  string
	MethodColor  string
	// tag
	Tag     string
	TagColor     string
	// Keys are the keys set on the request's context.
	//Keys map[string]interface{}
	start   time.Time
	end     time.Time
	
	//calculate the total Latency and average Latency
	average time.Duration
	total 	time.Duration
	cnts	int64
}

/*func New() *Latency {
	return &Latency{
		MethodColor: methodColor,
		TagColor:    tagColor,
		cnts:		 0,
	}
}*/

func New(tag, method string) *Latency {
	return &Latency{
		MethodColor: methodColor,
		TagColor:    tagColor,
		cnts:		 0,
		Tag:		 tag,
		Method:		 method,
	}
}

func NewWhithMethodColor(tag, method string, color string) *Latency {
	return &Latency{
		MethodColor: color,
		TagColor:    tagColor,
		cnts:		 0,
		Tag:		 tag,
		Method:		 method,
	}
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func GetColor(color string) string {
	switch {
	case color == "red":
		return red
	case color == "white":
		return white
	case color == "yellow":
		return yellow
	case color == "green":
		return green
	case color == "blue":
		return blue
	case color == "magenta":
		return magenta
	case color == "cyan":
		return cyan
	default:
		return reset
	}
}

func (f *Latency)Start() {
	f.start = time.Now()
}

func (f *Latency)End() string {

	f.cnts++
	f.end = time.Now()
	f.Latency = f.end.Sub(f.start)
	f.total += f.Latency

	return fmt.Sprintf("%s [%s] %s %v | %s %13v %s | %s %-8s %s \n",
		f.TagColor, f.Tag, resetColor,
		f.end.Format("2006/01/02 - 15:04:05"),
		latencyColor, f.Latency, resetColor,
		f.MethodColor, f.Method, resetColor,
	)
}

// get the total and average Latency
func (f *Latency)Total() string {
	//fmt.Println(f.total, ":",f.cnts)
	if (f.cnts == 0) {
		f.average = 0
	} else {
		f.average = time.Duration(int64(f.total)/(f.cnts))
	}
	//f.average = time.Duration(int64(f.total)/(f.cnts))
	//fmt.Println(f.total, ":",f.cnts,":",f.average)

	return fmt.Sprintf("%s [%s-total] %s %v | %s %19v %s | %s %13v %s | %s %13v %s | %s %-8s %s \n",
		f.TagColor, f.Tag, resetColor,
		f.end.Format("2006/01/02 - 15:04:05"),
		green, f.cnts, resetColor,
		cyan, f.total, resetColor,
		red, f.average, resetColor,
		f.MethodColor, f.Method, resetColor,
	)
}