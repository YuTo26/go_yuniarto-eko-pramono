# (10) Recursive - Number Theory - Sorting - Searching

*******************

1/  Materi Recursive:

- Recursive adalah teknik pemrograman yang digunakan untuk menyelesaikan masalah dengan memecahkannya menjadi sub-masalah yang lebih kecil dan sama jenisnya, hingga mencapai kasus dasar yang dapat dipecahkan secara langsung.
- Di Golang, fungsi rekursif ditulis seperti fungsi biasa dengan pemanggilan dirinya sendiri, namun perlu diperhatikan kasus dasar dan kondisi berhenti rekursi agar tidak terjadi infinite loop.
  
  // contoh sintaks dasar rekursif

  ```source code
  func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
    }
    ```

*******************

2/ Materi Number Theory:

- Number theory adalah studi tentang sifat-sifat bilangan bulat dan hubungannya dengan operasi matematika tertentu seperti faktorisasi, permutasi, dan kombinatorik.
- Di Golang, package math/big menyediakan struktur BigInt yang dapat digunakan untuk melakukan operasi aritmatika pada bilangan bulat dengan jumlah digit yang besar.

    //contoh sintaks dasar number theory

    ```source code
    import "math"
    func isPrime(n int) bool {
        if n < 2 {
            return false
        }
        for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }
    ```

*******************

3/ Materi Sorting:

- Sorting adalah proses pengurutan data dalam urutan tertentu, seperti ascending atau descending, dengan menggunakan algoritma tertentu seperti bubble sort, insertion sort, quicksort, dan lain-lain.
- Di Golang, package sort menyediakan fungsi untuk melakukan sorting dengan algoritma quicksort, heapsort, dan insertion sort.

    //contoh sintaks dasar sorting

    ```source code
    func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
    }
    ```

*******************

4/ Materi Searching:

- Searching adalah proses mencari data tertentu dalam kumpulan data yang besar, dengan menggunakan algoritma tertentu seperti linear search, binary search, dan lain-lain.
- Di Golang, package sort juga menyediakan fungsi untuk melakukan binary search pada data yang telah diurutkan dengan algoritma binary search. Selain itu, package strings juga menyediakan fungsi untuk melakukan pencarian string dengan algoritma Boyer-Moore dan Knuth-Morris-Pratt.

    //contoh sintaks dasar Searching

    ```source code
    func binarySearch(arr []int, x int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] < x {
            low = mid + 1
        } else if arr[mid] > x {
            high = mid - 1
        } else {
            return mid
        }
    }
    return -1
    }
    ```
