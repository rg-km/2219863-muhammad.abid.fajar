// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {
<<<<<<< HEAD
	hash := make(map[string]bool)
    for _, e := range str1 {
        hash[e] = true
    }
    for _, e := range str2 {
        // If elements present in the hashmap then append intersection list.
        if hash[e] {
            inter = append(inter, e)
        }
    }
    //Remove dups from slice.
    inter = RemoveDuplicates(inter)
    return // TODO: replace this
}

func RemoveDuplicates(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return // TODO: replace this
=======
	return []string{} // TODO: replace this
}

func RemoveDuplicates(elements []string) (nodups []string) {
	return []string{} // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
