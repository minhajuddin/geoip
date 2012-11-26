package main

type Location struct {
	CountryCode string
	IpStart     int32
	IpEnd       int32
}

//sort interface
type Locations []Location

func (self Locations) Len() int           { return len(self) }
func (self Locations) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self Locations) Less(i, j int) bool { return self[i].IpStart < self[j].IpStart }

//search
//copied straight from sort/search.go
func SearchLocation(ip int32) string {
	i, j := 0, len(db)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if db[h].IpStart < ip {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return db[i].CountryCode
}
