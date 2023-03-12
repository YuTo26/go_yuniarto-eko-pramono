# (11) Problem Solving Paradigm - Brute Force , Greedy anda Dynamic Programming

****************

Problem Solving Paradigm adalah teknik atau pendekatan yang digunakan dalam pemrograman untuk menyelesaikan suatu masalah.
Beberapa paradigma yang umum digunakan dalam pemrograman termasuk Brute Force, Greedy, dan Dynamic Programming.

****************

1/ Brute Force:

- Brute Force adalah teknik pemrograman yang digunakan untuk menyelesaikan masalah dengan mencoba semua kemungkinan solusi secara sistematis.
- Kelebihan dari Brute Force adalah dapat menemukan solusi yang optimal untuk masalah, namun kelemahannya adalah waktu eksekusi yang cukup lama terutama untuk masalah yang kompleks.
- Di Golang, Brute Force dapat diimplementasikan dengan cara melakukan loop untuk mencoba semua kemungkinan solusi.

    //contoh sintaks dasar brute Force

    ```source
    func BruteForce(numbers []int, target int) bool {
    for _, num1 := range numbers {
        for _, num2 := range numbers {
            if num1+num2 == target {
                return true
            }
        }
    }
    return false
    }
    ```

****************

2/ Greedy:

- Greedy adalah teknik pemrograman yang digunakan untuk menyelesaikan masalah dengan memilih solusi yang terbaik pada setiap langkah, tanpa mempertimbangkan dampak jangka panjang.
- Kelebihan dari Greedy adalah waktu eksekusi yang cepat, namun kelemahannya adalah tidak menjamin solusi optimal untuk masalah.
- Di Golang, Greedy dapat diimplementasikan dengan cara memilih solusi terbaik pada setiap langkah dan mengulanginya hingga mencapai solusi akhir.

    //contoh sintaks dasar Greedy

    ```source code
    func Greedy(coins []int, amount int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(coins)))
    count := 0
    for _, coin := range coins {
        for amount >= coin {
            count++
            amount -= coin
        }
    }
    return count
    }
    ```

****************

3/ Dynamic Programming:

- Dynamic Programming adalah teknik pemrograman yang digunakan untuk menyelesaikan masalah dengan memecahkannya menjadi sub-masalah yang lebih kecil dan menyimpan solusi dari setiap sub-masalah tersebut.
- Kelebihan dari Dynamic Programming adalah dapat menemukan solusi yang optimal dan waktu eksekusi yang lebih cepat daripada Brute Force.
- Di Golang, Dynamic Programming dapat diimplementasikan dengan cara menyimpan solusi dari setiap sub-masalah dalam bentuk tabel atau array, dan menggunakan solusi tersebut untuk menyelesaikan masalah secara keseluruhan.

    //contoh sintaks dynamic Programming

    ```source code
    func DynamicProgramming(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
    }
    ```
