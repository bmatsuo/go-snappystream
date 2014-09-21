package snappystream

import "io/ioutil"

var testDataMan []byte

// curl -s https://api.github.com/users/mreiferson/repos
var testDataJSON []byte

func init() {
	var err error
	testDataMan, err = ioutil.ReadFile("fixturedata/man-xargs.txt")
	if err != nil {
		panic(err)
	}
	testDataJSON, err = ioutil.ReadFile("fixturedata/github-users-repos.json")
	if err != nil {
		panic(err)
	}
}
