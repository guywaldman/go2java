package greeting

// GreetPerson returns a greeting for the received name.
// This is an example of an existing function you have in Go,
// which you would like to expose to Java (or JVM in general).
func GreetPerson(name string) *string {
	greeting := "Hello, " + name + "!"
	return &greeting
}
