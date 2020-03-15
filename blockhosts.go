package main

import(
    "fmt"
    "os"
)

func main() {
    var filePath string
    var maker4 bool = false
    var maker6 bool = false
    var makerh bool = false
    var makerd bool = false

    if len(os.Args) > 1 && len(os.Args) < 7 {
        for i := 1; i < len(os.Args); i++ {
            s := os.Args[i]
            switch s {
            case "-f":
                maker4 = true
                break
            case "-s":
                maker6 = true
                break
            case "-h":
                makerh = true
                break
            case "-d":
                makerd = true
                break
            default:
                filePath = s
                break
            }
        }
        if filePath == "" {
            fmt.Println("Error file path, exited.")
            return
        }
        fmt.Println(filePath)
        fmt.Println(maker4)
        fmt.Println(maker6)
        fmt.Println(makerh)
        fmt.Println(makerd)
    } else {
        fmt.Println("Error argument(s), exited.")
        fmt.Println("The arguments are:")
        fmt.Println("  -f    Default: IPv4")
        fmt.Println("  -s    IPv6")
        fmt.Println("  -h    Default: Hosts(Windows, macOS, Linux)")
        fmt.Println("  -d    Dnsmasq")
    }
}
