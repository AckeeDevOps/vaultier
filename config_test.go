package main

log = New(func(string) {})

func TestTimeConsuming(t *testing.T) {
	log.Fatal("bla")
}
