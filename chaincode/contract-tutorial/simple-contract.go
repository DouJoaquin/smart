package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"io/ioutil"
)

// SimpleContract contract for handling writing and reading from the world state
type SimpleContract struct {
	contractapi.Contract
}
// Create adds a new key with value to the world state
func (sc *SimpleContract) Create(ctx contractapi.TransactionContextInterface, key string, value string) error {
	existing, err := ctx.GetStub().GetState(key)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with key %s. Already exists", key)
	}
	err = ctx.GetStub().PutState(key, []byte(value))

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	return nil
}
// Update changes the value with key in the world state
func (sc *SimpleContract) Update(ctx contractapi.TransactionContextInterface, key string, value string) error {
	existing, err := ctx.GetStub().GetState(key)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing == nil {
		return fmt.Errorf("Cannot update world state pair with key %s. Does not exist", key)
	}

	err = ctx.GetStub().PutState(key, []byte(value))

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	return nil
}
// Read returns the value at key in the world state
func (sc *SimpleContract) Read(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	existing, err := ctx.GetStub().GetState(key)

	if err != nil {
		return "", errors.New("Unable to interact with world state")
	}

	if existing == nil {
		return "", fmt.Errorf("Cannot read world state pair with key %s. Does not exist", key)
	}

	return string(existing), nil
}

// AddDocument adds a new key with value to the world state
func (sc *SimpleContract) AddDocument (ctx contractapi.TransactionContextInterface, key string, documentName string) error {
	existing, err := ctx.GetStub().GetState(key)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with key %s. Already exists", key)
	}
	//filepath := "/home/mylinux/go/src/github.com/hyperledger/fabric/scripts/fabric-samples/txtdata/"+ documentName
	filepath := "/home/mylinux/go/src/github.com/hyperledger/fabric/scripts/fabric-samples/txtdata/1.txt"
	//filepath :="/data/1.txt"
	content ,err :=ioutil.ReadFile(filepath)
	if err !=nil {
		return errors.New("Unable to open file in basic3")
	}
	err = ctx.GetStub().PutState(key, []byte(string(content)))
	if err != nil {
		return errors.New("Unable to interact with world state")
	}
	//fmt.Println("OK")

	return nil
}
