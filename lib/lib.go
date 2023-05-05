package lib
import (
	"encoding/binary"
	"fmt"
	"math"
)

func toHour(seconds float64) int {
	return int(seconds/60/60)
}

func toMinutes(seconds float64) int {
	return int(seconds/60)%60
}

func toSeconds(seconds float64) int {
	return int(seconds)%60
}

func fractionOfSecond(seconds float64) float64 {
	return (seconds - float64(int(seconds)))*1000
}

func PrintTime(time float64, fileNumber int) {
	fmt.Printf("Save %d: %.2f seconds: %02d:%02d:%02d.%3.0f\n", fileNumber, time, toHour(time), toMinutes(time), toSeconds(time), fractionOfSecond(time))
}

func FindFromOffsets(data []byte, offset int) float64 {
	time := float64(0)
	start := uint64(0)
	for i := 0; i <= offset; i++ {
		bit := (data[i/8] >> (7 - (i % 8))) & 1
		start = (start << 1) | uint64(bit)
		start = start & 0xFFFFFFFF_FFFFFFFF

		if i == offset {
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, start)
			time = math.Float64frombits(binary.BigEndian.Uint64(buf))
			break
		}
	}
	return time
}
