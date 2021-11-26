package main

import "fmt"
import "time"
import "math"

func getMilliSecondsOfDay(t time.Time) int {
    return (3600000 * t.Hour()) + (60000 * t.Minute()) + (1000 * t.Second()) + (t.Nanosecond() / 1000000)
}


func main() {
    go getHazMazSaz()
    select {}
}

func getHazMazSaz() {
    for {
        t := time.Now()
        var currentTime int = getMilliSecondsOfDay(t)

        var haz, maz, saz int 
        var timeLeft float64

        haz = int(math.Floor(float64(currentTime) / 864000))
        timeLeft = float64(currentTime) - (float64(haz) * 864000)

        maz = int(math.Floor(float64(timeLeft) / 8640))
        timeLeft -= float64(maz) * 8640

        saz = int(math.Floor(float64(timeLeft) / 86.4))
        timeLeft -= float64(saz) * 86.4

        fmt.Printf("\033[2K\r %d : %d : %d", haz, maz, saz)

        time.Sleep(time.Second / 15)
    }
}
