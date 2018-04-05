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

/* Define Tuna structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Vote struct {
	// Vessel    string `json:"vessel"`
	// Timestamp string `json:"timestamp"`
	// Location  string `json:"location"`
	// Owner     string `json:"owner"`
	Codigo    string `json:"codigo"`
	Eleccion  string `json:"eleccion"`
	Entidad   string `json:"entidad"`
	Distrito  string `json:"distrito"`
	Municipio string `json:"municipio"`
	Seccion   string `json:"seccion"`
	Localidad string `json:"localidad"`
}

// type Vote struct {
// 	Codigo    string `json:"codigo"`
// 	Eleccion  string `json:"eleccion"`
// 	Entidad   string `json:"entidad"`
// 	Distrito  string `json:"distrito"`
// 	Municipio string `json:"municipio"`
// 	Seccion   string `json:"seccion"`
// 	Localidad string `json:"localidad"`
// }

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

	// if function == "queryTuna" {
	// 	return s.queryTuna(APIstub, args)
	// } else
	// else if function == "recordTuna" {
	// 	return s.recordTuna(APIstub, args)
	// } else if function == "queryAllTuna" {
	// 	return s.queryAllTuna(APIstub)
	// } else if function == "changeTunaHolder" {
	// 	return s.changeTunaHolder(APIstub, args)
	// }

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryTuna method *
Used to view the records of one particular tuna
It takes one argument -- the key for the tuna in question
*/
// func (s *SmartContract) queryTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	tunaAsBytes, _ := APIstub.GetState(args[0])
// 	if tunaAsBytes == nil {
// 		return shim.Error("Could not locate tuna")
// 	}
// 	return shim.Success(tunaAsBytes)
// }

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
	votes := []Vote{
		Vote{Codigo: "1234", Eleccion: "AMLO", Entidad: "CDMX", Distrito: "1",
			Municipio: "Coyoacan", Seccion: "3", Localidad: "23023"},
		Vote{Codigo: "1434", Eleccion: "Anaya", Entidad: "Chiapas", Distrito: "45",
			Municipio: "Tlalpan", Seccion: "3458", Localidad: "230"},
		Vote{Codigo: "18712", Eleccion: "Meade", Entidad: "Chihuaha", Distrito: "108",
			Municipio: "Coyoacan", Seccion: "60890", Localidad: "130890"},
		Vote{Codigo: "10980", Eleccion: "AMLO", Entidad: "Colima", Distrito: "99",
			Municipio: "Talpan", Seccion: "12309", Localidad: "34567"},
		Vote{Codigo: "12345", Eleccion: "Anaya", Entidad: "Durango", Distrito: "234",
			Municipio: "Milpa Alta", Seccion: "67120", Localidad: "22000"},
		Vote{Codigo: "6574", Eleccion: "Meade", Entidad: "Guerrero", Distrito: "129",
			Municipio: "Tlahuac", Seccion: "230", Localidad: "200000"},
		Vote{Codigo: "34658", Eleccion: "AMLO", Entidad: "Hidalgo", Distrito: "13",
			Municipio: "Coyoacan", Seccion: "340", Localidad: "9943"},
		Vote{Codigo: "956734", Eleccion: "Anaya", Entidad: "Jalisco", Distrito: "9",
			Municipio: "Tlalpan", Seccion: "908", Localidad: "77712"},
		Vote{Codigo: "4513", Eleccion: "Meade", Entidad: "Morelos", Distrito: "134",
			Municipio: "Coyoacan", Seccion: "1230", Localidad: "666"},
		Vote{Codigo: "12087", Eleccion: "AMLO", Entidad: "Nayarit", Distrito: "189",
			Municipio: "Xochimilco", Seccion: "8976", Localidad: "12353"},
		// Tuna{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Owner: "Dave"},
		// Tuna{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Owner: "Igor"},
		// Tuna{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Owner: "Amalea"},
		// Tuna{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Owner: "Rafa"},
		// Tuna{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Owner: "Shen"},
		// Tuna{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Owner: "Leila"},
		// Tuna{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Owner: "Yuan"},
		// Tuna{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Owner: "Carlo"},
		// Tuna{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Owner: "Fatima"},
	}

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
 * The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
// func (s *SmartContract) recordTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 5 {
// 		return shim.Error("Incorrect number of arguments. Expecting 5")
// 	}

// 	var tuna = Tuna{Vessel: args[1], Location: args[2], Timestamp: args[3], Owner: args[4]}

// 	tunaAsBytes, _ := json.Marshal(tuna)
// 	err := APIstub.PutState(args[0], tunaAsBytes)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
// 	}

// 	return shim.Success(nil)
// }

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
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results.
*/
// func (s *SmartContract) queryAllTuna(APIstub shim.ChaincodeStubInterface) sc.Response {

// 	startKey := "0"
// 	endKey := "999"

// 	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add comma before array members,suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(queryResponse.Key)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.WriteString(string(queryResponse.Value))
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

// 	return shim.Success(buffer.Bytes())
// }

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
 * The changeTunaHolder method *
The data in the world state can be updated with who has possession.
This function takes in 2 arguments, tuna id and new holder name.
*/
// func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	tunaAsBytes, _ := APIstub.GetState(args[0])
// 	if tunaAsBytes == nil {
// 		return shim.Error("Could not locate tuna")
// 	}
// 	tuna := Tuna{}

// 	json.Unmarshal(tunaAsBytes, &tuna)
// 	// Normally check that the specified argument is a valid holder of tuna
// 	// we are skipping this check for this example
// 	tuna.Owner = args[1]

// 	tunaAsBytes, _ = json.Marshal(tuna)
// 	err := APIstub.PutState(args[0], tunaAsBytes)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to change tuna holder: %s", args[0]))
// 	}

// 	return shim.Success(nil)
// }

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
