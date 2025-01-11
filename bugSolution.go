func main() {

        x := 10

        defer func(x int) {

                fmt.Println("defer function: ", x)

        }(x)

        x = 5

        fmt.Println("main function: ", x)

}