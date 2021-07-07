// kullanıcıdan belli bir x y ve yarı çap değeri alır ona göre çember çizen uygulama
package main

import (
  	"fmt"
    "math"
)

const lx = 30
const ly = 30


func main() {
	var circleX int
	var circleY int
	var circleR int
	var mainMatris [lx][ly]string

//	matrisi dolduruyorum
	for i := 0; i < ly; i++{
        for j := 0; j < lx; j++{
            mainMatris[j][i] = "."
        }
    }



	fmt.Println("DİKKAT!! matris", lx, ly, "uzunluğunda lütfen taşırmayınız")
	for unimportantVar := 0; unimportantVar < 10; unimportantVar++{
		fmt.Print("Lütfen çemberin x koordinatını giriniz = ")
		fmt.Scanln(&circleX)
		fmt.Print("Lütfen çemberin y koordinatını giriniz = ")
		fmt.Scanln(&circleY)
		fmt.Print("Lütfen çemberin yarı çapını giriniz = ")
		fmt.Scanln(&circleR)


//		çemberi yazdırıyorum
		R := float64(circleR)


		for i := 0; i <= circleR; i++{
			cos := float64(i)
			//  BİRİNCİ BÖLGE   //
			mainMatris [circleX + i] [circleY + int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] = "#"
			mainMatris [circleX + int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] [circleY + i] = "#"
			//   İKİNCİ BÖLGE   //
			mainMatris [circleX - i] [circleY + int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] = "#"
			mainMatris [circleX - int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] [circleY + i] = "#"
			//   ÜÇÜNCÜ BÖLGE   //
			mainMatris [circleX - i] [circleY - int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] = "#"
			mainMatris [circleX - int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] [circleY - i] = "#"
			//  DÖRDÜNCÜ BÖLGE  //
			mainMatris [circleX + i] [circleY - int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] = "#"
			mainMatris [circleX + int(math.RoundToEven(math.Sqrt(R * R - cos * cos)))] [circleY - i] = "#"
		}

//		matrisi yazdırıyorum

	    for i := 0; i < ly; i++{
	        for j := 0; j < lx; j++{
	            fmt.Print(" ",mainMatris[j][i])
	        }
	        fmt.Println()
	    }
	}
}
