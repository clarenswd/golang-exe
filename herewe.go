
package main;
import "fmt";


func main(){
    yourAge := 28;
    if yourAge >= 30{
        fmt.Println("you are over 30")
    }else if yourAge == 28 {
        fmt.Println("you are 28")
    }else{
        fmt.Println("you are less than 30")
    }


    switch yourAge{
        case 18 : fmt.Println("eighteen")
        case 28 : fmt.Println("eighteen plus ten")
        default : fmt.Println("non recognized param")
    }
}