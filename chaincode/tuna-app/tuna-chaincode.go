// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
*/

package main

/* Imports
* 4 utility libraries for handling bytes, reading and writing JSON,
formatting, and string manipulation
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Vote structure, with 7 properties.
Structure tags are used by encoding/json library
*/
type Vote struct {
	Codigo    string `json:"codigo"`
	Eleccion  string `json:"eleccion"`
	Entidad   string `json:"entidad"`
	Distrito  string `json:"distrito"`
	Municipio string `json:"municipio"`
	Seccion   string `json:"seccion"`
	Localidad string `json:"localidad"`
}

/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function
 -- see initLedger()
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
*/
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordVote" {
		return s.recordVote(APIstub, args)
	} else if function == "queryVote" {
		return s.queryVote(APIstub, args)
	} else if function == "queryAllVotes" {
		return s.queryAllVotes(APIstub)
	} else if function == "queryVotesByDomain" {
		return s.queryVotesByDomain(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryVote method *
Used to view the records of one particular vote
It takes one argument -- the key for the vote in question
*/
func (s *SmartContract) queryVote(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	voteAsBytes, _ := APIstub.GetState(args[0])
	fmt.Println(args[0])
	if voteAsBytes == nil {
		return shim.Error("Could not locate vote")
	}
	return shim.Success(voteAsBytes)
}

/*
 * The queryVotesByDomain method *
Used to view the records of several votes with couchdb rich queries
It takes two arguments -- the domain of the query and the specific value
*/
func (s *SmartContract) queryVotesByDomain(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	searchDomain := args[0]
	param := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"%s\":\"%s\"}}", searchDomain, param)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

/*
 * The initLedger method *
Will add test data (10 tuna catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	votes := []Vote{}

	i := 0
	for i < len(votes) {
		fmt.Println("i is ", i)
		voteAsBytes, _ := json.Marshal(votes[i])
		APIstub.PutState(strconv.Itoa(i+1), voteAsBytes)
		fmt.Println("Added", votes[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordVote method *
This method takes in seven arguments (attributes to be saved in the ledger).
*/
func (s *SmartContract) recordVote(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var vote = Vote{Codigo: args[1], Eleccion: args[2], Entidad: args[3], Distrito: args[4],
		Municipio: args[5], Seccion: args[6], Localidad: args[7]}

	voteAsBytes, _ := json.Marshal(vote)
	err := APIstub.PutState(args[0], voteAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record vote: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllVote method *
allows for assessing all the records added to the ledger(all votes)
This method does not take any arguments. Returns JSON string containing results.
*/
func (s *SmartContract) queryAllVotes(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := ""
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllVote:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

/*
 * main function *
calls the Start function
The main function starts the chaincode in the container during instantiation.
*/
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
