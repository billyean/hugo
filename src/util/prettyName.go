package util

func GetPrefix() map[string]string {
	//var nameWithPrefix map[string]string
	//nameWithPrefix = make(map[string]string)
	//nameWithPrefix["Tristan"] = "Mr"
	//nameWithPrefix["Tina"] = "Mrs"
	//nameWithPrefix["Ruby"] = "Miss"
	//nameWithPrefix["Zachary"] = "Buddy"
	//nameWithPrefix["Rachael"] = "Buddy"

	nameWithPrefix := map[string]string {
		"Tristan" : "Mr",
		"Tina" : "Mrs",
		"Ruby" : "Miss",
		"Zachary" : "Buddy",
		"Rachael" : "Buddy",
	}

	// Update
	nameWithPrefix["Ruby"] = "Dr"
	nameWithPrefix["Rachael"] = "PHD"

	//Insertion
	nameWithPrefix["Frank"] = "Dude"
	nameWithPrefix["JohnDoe"] = "Hello"

	//Deletion
	delete(nameWithPrefix, "Frank")



	return nameWithPrefix
}