package main

import (
	"errors"
	"fmt"
	"maps"
)

type user struct {
	name        string
	phoneNumber int
}

type users struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

func deleteIfNecessary(users map[string]users, name string) (deleted bool, err error) {
	if _, ok := users[name]; !ok {
		return false, errors.New("not found")
	}
	ele := users[name]
	if ele.scheduledForDeletion {
		delete(users, name)
		return true, nil
	} else {
		return false, nil
	}
}

func getCounts(userIds []string) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(userIds); i++ {
		
		counts[userIds[i]]++
		
	}
	return counts
}

func main() {
	names := []string{"keshav", "ruchi"}
	phoneNumber := []int{123423443, 12342314}
	usermap, err := getUserMap(names, phoneNumber)
	if err != nil {
		fmt.Println("some eror occured")
	}
	for key, value := range usermap {
		fmt.Println(key, value)
	}

	m := make(map[string]int)
	m["k1"] = 3
	m["k2"] = 2

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k2"]
	fmt.Println("v2: ", v3)

	fmt.Println("length of map m: ", len(m))

	delete(m, "k2")
	fmt.Println("map m: ", m)

	clear(m)
	fmt.Println("map m: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("new map: ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n2)
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}

}
