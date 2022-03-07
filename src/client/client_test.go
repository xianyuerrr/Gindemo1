package client

import (
	"fmt"
	"grayRelease/src/client/operationTest"
	"io/ioutil"
	"log"
	"testing"
)

func Test_testOperationAdd(t *testing.T) {
	add := operationTest.AddTest
	resp := testOperation(add)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func Test_testOperationUpdate(t *testing.T) {
	update := operationTest.UpdateTest
	resp := testOperation(update)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func Test_testOperationDelete(t *testing.T) {
	delete := operationTest.DeleteTest
	resp := testOperation(delete)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func Test_testOperationRelease(t *testing.T) {
	release := operationTest.ReleaseTest
	resp := testOperation(release)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func Test_testOperationOffline(t *testing.T) {
	offline := operationTest.OfflineTest
	resp := testOperation(offline)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func Test_testOperationCheck(t *testing.T) {
	check := operationTest.CheckTest
	resp := testOperation(check)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
