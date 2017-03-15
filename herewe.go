
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


    //arrays


    var favNums2[5] float64;
    favNums2[0] = 163;
    favNums2[1] = 1.3;
    favNums2[2] = 1.63;
    favNums2[3] = 13;
    favNums2[4] = 3;

    fmt.Println(favNums2[3])


    favNums3 := [5] float64{1,3,2,4,5};
    fmt.Println("printing array vals with index")
    for i,value := range favNums3{
        fmt.Println(value, i)
    }
    fmt.Println("printing array vals with NO index")

    for _,value := range favNums3{
        fmt.Println(value)
    }














}