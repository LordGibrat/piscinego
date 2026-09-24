package piscine
import "github.com/01-edu/z01"
func Printstr(s string){
    for _, r :=range s{
        z01.PrintRune(r)
    }
}
