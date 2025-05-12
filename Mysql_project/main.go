package main

import(
	
)

func main() {
	connectDB()
	GetTables()
	// InsertInto(101, "Ganesh", "ganeshmahity@example.com", "Mumbai", time.Now())
	UpdateUserEmail(101, "bob.new@example.com")
	DeleteUser(21)
	defer dbm.Close()
}