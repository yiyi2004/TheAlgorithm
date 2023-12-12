package main

func main() {

}

// 我感觉我的头不是很舒服
// 一个课程竟然需要两个前置课程

func canFinish(numCourses int, prerequisites [][]int) bool {
	preCourseNum := make([]int, numCourses)
	postCourses := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		preCourseNum[prerequisites[i][0]]++
		postCourses[prerequisites[i][1]] = append(postCourses[prerequisites[i][1]], prerequisites[i][0])
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if preCourseNum[i] == 0 {
			queue = append(queue, i)
		}
	}
	learned := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		learned++
		for _, v := range postCourses[course] {
			preCourseNum[v]--
			if preCourseNum[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return learned == numCourses
}
