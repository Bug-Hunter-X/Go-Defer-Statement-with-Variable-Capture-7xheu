func main() {



        x := 10

        defer func() {

                fmt.Println("defer function: ", x)

        }()

        x = 5

        fmt.Println("main function: ", x)

}