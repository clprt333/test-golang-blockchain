package blockchain

//Imports
import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//Implementing the Proof of Work Algorithm.
//Difficulty uses a Static Number for this Scenario.
const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//Take the Data from a Block. Compare with a Static Difficulty Value.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}

	return pow
}

//Create a Nonce that begins at 0.
//Create a Hash from the Data + Counter.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToBytes(int64(nonce)),
			ToBytes(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

//Infinite Loop
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	//Read the Hashes, Return the Nonce.
	fmt.Println()

	return nonce, hash[:]
}

//Check the Hash to see if it meets a set of Requirements.
//Validate using a Boolean Function.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

//Create a Buffer for the Data. (Converts Bytes from Binary to Hex)
func ToBytes(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
