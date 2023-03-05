package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name   string
	Scores []float64
}

func (s *Student) AddScore(score float64) {
	s.Scores = append(s.Scores, score)
}

func (s Student) AverageScore() float64 {
	var total float64
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

func (s Student) MinScore() float64 {
	minScore := s.Scores[0]
	for _, score := range s.Scores[1:] {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}

func (s Student) MaxScore() float64 {
	maxScore := s.Scores[0]
	for _, score := range s.Scores[1:] {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	var students [5]Student

	for i := 0; i < 5; i++ {
		var name string
		var scores []float64
		fmt.Printf("Masukkan nama siswa %d: ", i+1)
		fmt.Scanln(&name)

		for j := 0; j < 1; j++ {
			var score float64
			fmt.Printf("Masukkan nilai ke-%d untuk siswa %s: ", j+1, name)
			fmt.Scanln(&score)
			scores = append(scores, score)
		}

		students[i] = Student{Name: name, Scores: scores}
	}

	var totalScore float64
	var allScores []float64
	for _, student := range students {
		totalScore += student.AverageScore()
		allScores = append(allScores, student.Scores...)
	}

	fmt.Printf("Rata-rata skor dari semua siswa adalah: %.2f\n", totalScore/float64(len(students)))

	sort.Float64s(allScores)
	fmt.Printf("Siswa dengan skor terendah adalah %s dengan skor %0.2f\n", findStudentNameByScore(students, allScores[0]), allScores[0])
	fmt.Printf("Siswa dengan skor tertinggi adalah %s dengan skor %0.2f\n", findStudentNameByScore(students, allScores[len(allScores)-1]), allScores[len(allScores)-1])
}

func findStudentNameByScore(students [5]Student, score float64) string {
	for _, student := range students {
		if student.MinScore() == score {
			return student.Name
		}
		if student.MaxScore() == score {
			return student.Name
		}
	}
	return ""
}
