package main

func toBool(b bool) bool {
	if(b == true || b == 1 || b == "true" || b == "1") {
		return true;
	}
	return false;
}
func main() {
	var bool = "true"
	var b = toBool(bool)
	fmt.Printf("toBool(%s) = %t\n", bool, b)

}
