package main
import (
	"fmt";
	"reflect"
)

// func cal(x int, y int) (a int, b int) {
// 	a = x / y
// 	b = x * y

// 	return
// }

func cal(x int, y int) (int, int) {
	a := x / y
	b := x * y

	return a, b
}

func calsum(x,y int) {
	fmt.Println(x + y)
}

func confirmCal(x int, y int) int {
	return x + y
}

type Student struct {
	name string
	math, english float64
}

func (s Student) avg(math, english float64) float64 {
	return (math + english) / 2
}

// func (s Student) avg() {
// 	fmt.Println(s.name, (s.math + s.english) / 2)
// }

type User struct {
	gender string
	age int
}

type Userbmi struct {
	name string
}

func (ubmi Userbmi) calbmi(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

type StudentTest struct {
	name string
}

func (s StudentTest) calAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s StudentTest) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main(){
	fmt.Println("Good morning")
	fmt.Println("Good afternoon")
	fmt.Println("Good evening")

	// var num int
	// num = 1

	// var num = 1

	num := 1
	fmt.Println(num)

	num01 := 123
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))

	var string_a string = "Hello,World!"
	string_b := "Hello,World!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

	a := 10
	b := 1
	var num_bool bool = a > b
	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

	// arr_a := [...]string{"sato", "takanashi"}
	// fmt.Println(arr_a[0])

	arr_a := [2][2]string{{"sato", "suzuki"}, { "takahashi", "tanaka"}}
	arr_a[1][1] = "yamaguchi"

	fmt.Println(arr_a[0][0])
	fmt.Println(arr_a[0][1])
	fmt.Println(arr_a[1][0])
	fmt.Println(arr_a[1][1])

	x := 10
	y := 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	xx := 10
	yy := 2
	zz := 10

	fmt.Println(xx <= yy)
	fmt.Println(xx >= zz)

	fmt.Println(xx == yy)
	fmt.Println(xx != yy)

	fmt.Println(xx >= 3 && xx <= 10)
	fmt.Println(yy >= 2 && yy <= 3)

	fmt.Println(xx == 3 || yy == 2)
	fmt.Println(xx == 1 || yy == 3)

	xxx := 8
	yyy := 12
	zzz := 20
	xxx += 10
	zzz += yyy
	xxx++
	yyy--

	fmt.Println(xxx)
	fmt.Println(yyy)
	fmt.Println(zzz)

	xxxx := 10
	yyyy := 12

	if age := xxxx + yyyy ; age >= 20 {
		// 簡易文はif文の中でしか使用できない
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

	carrer := 1
	if carrer += 1 ; 5 <= carrer && carrer < 10 {
		fmt.Println(carrer)
	}

	for i := 0; i <= 4; i++ {
		if i == 3 {
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	// 外側のループのiが0周目の時にjが0から2までまわり次にiが1周目に入る
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Println(i, "-", j)
		}
	}

	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
		sum += arr[i]
	}
	fmt.Println(sum)

	ii := 0
	for ii <=4 {
		fmt.Println(ii)
		ii++
	}

	iii := 0
	for {
		fmt.Println(iii)
		if iii == 4 {
			break
		}
		iii++
	}

	for rank := 1; rank <= 10; rank++ {
		fmt.Println(rank)
	}

	rank := 1
	for rank <= 10 {
		fmt.Println(rank)
		rank++
	}

	ranking := 1
	for {
		if ranking == 11 {
			break
		}
		fmt.Println(ranking)
		ranking++
	}

	hello := func(greeting string) {
		fmt.Println(greeting)
	}

	func(greeting string) {
		fmt.Println(greeting)
	}("Good evening")

	hello("Good morning")
	
	result01, result02 := cal(6, 3)
	fmt.Println(result01, result02)
	calsum(6, 3)

	fmt.Println(confirmCal(10, 5))

	// 構造体の初期化
	// var s Student
	// s.name = "sato"
	// s.math = 80
	// s.english = 70

	// 構造体の初期化
	// s := Student{"sato", 80, 70}
	s := Student{english: 70}

	fmt.Println(s)

	// var u User
	// u.gender = "male"
	// u.age = 20
	u := User{"male", 20}
	
	fmt.Println(u)

	// 構造体の初期化
	var a001 Student
	a001.name = "sato"
	fmt.Println(a001.avg(80, 70))

	// 構造体の初期化
	ubmi := Userbmi{"haruki"}
	fmt.Println(ubmi.name ,ubmi.calbmi(178, 70))

	// 総仕上げ
	t001 := StudentTest{"sato"}
	data := []float64{70, 65, 50, 90, 30}
	var avgTest float64 = t001.calAvg(data)
	resultTest := t001.judge(avgTest)
	fmt.Println(avgTest)
	fmt.Println(t001.name + " " + resultTest)
}
