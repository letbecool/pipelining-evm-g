// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.

// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

// import (
// 	"fmt"
// 	"hash"
// 	"sync/atomic"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/math"
// 	"github.com/ethereum/go-ethereum/log"
// )

// // Config are the configuration options for the Interpreter
// type Config struct {
// 	Debug                   bool   // Enables debugging
// 	Tracer                  Tracer // Opcode logger
// 	NoRecursion             bool   // Disables call, callcode, delegate call and create
// 	EnablePreimageRecording bool   // Enables recording of SHA3/keccak preimages

// 	JumpTable [256]*operation // EVM instruction table, automatically populated if unset

// 	EWASMInterpreter string // External EWASM interpreter options
// 	EVMInterpreter   string // External EVM interpreter options

// 	ExtraEips []int // Additional EIPS that are to be enabled
// }

// // Interpreter is used to run Ethereum based contracts and will utilise the
// // passed environment to query external sources for state information.
// // The Interpreter will run the byte code VM based on the passed
// // configuration.
// type Interpreter interface {
// 	// Run loops and evaluates the contract's code with the given input data and returns
// 	// the return byte-slice and an error if one occurred.
// 	Run(contract *Contract, input []byte, static bool) ([]byte, error)
// 	// CanRun tells if the contract, passed as an argument, can be
// 	// run by the current interpreter. This is meant so that the
// 	// caller can do something like:
// 	//
// 	// ```golang
// 	// for _, interpreter := range interpreters {
// 	//   if interpreter.CanRun(contract.code) {
// 	//     interpreter.Run(contract.code, input)
// 	//   }
// 	// }
// 	// ```
// 	CanRun([]byte) bool
// }

// // callCtx contains the things that are per-call, such as stack and memory,
// // but not transients like pc and gas
// type callCtx struct {
// 	memory   *Memory
// 	stack    *Stack
// 	rstack   *ReturnStack
// 	contract *Contract
// }

// // keccakState wraps sha3.state. In addition to the usual hash methods, it also supports
// // Read to get a variable amount of data from the hash state. Read is faster than Sum
// // because it doesn't copy the internal state, but also modifies the internal state.
// type keccakState interface {
// 	hash.Hash
// 	Read([]byte) (int, error)
// }

// // EVMInterpreter represents an EVM interpreter
// type EVMInterpreter struct {
// 	evm *EVM
// 	cfg Config

// 	hasher    keccakState // Keccak256 hasher instance shared across opcodes
// 	hasherBuf common.Hash // Keccak256 hasher result array shared aross opcodes

// 	readOnly   bool   // Whether to throw on stateful modifications
// 	returnData []byte // Last CALL's return data for subsequent reuse
// }

// // NewEVMInterpreter returns a new instance of the Interpreter.
// func NewEVMInterpreter(evm *EVM, cfg Config) *EVMInterpreter {
// 	// We use the STOP instruction whether to see
// 	// the jump table was initialised. If it was not
// 	// we'll set the default jump table.

// 	fmt.Println("file: intepreter.go \t func: NewEVMInterpreter, \t  Descr: jump table is set, as a default frontierInstructionSet is going to be choosen. \n jump table is :")

// 	if cfg.JumpTable[STOP] == nil {
// 		var jt JumpTable

// 		fmt.Println(jt)

// 		switch {
// 		case evm.chainRules.IsYoloV1:
// 			jt = yoloV1InstructionSet
// 		case evm.chainRules.IsIstanbul:
// 			jt = istanbulInstructionSet
// 		case evm.chainRules.IsConstantinople:
// 			jt = constantinopleInstructionSet
// 		case evm.chainRules.IsByzantium:
// 			jt = byzantiumInstructionSet
// 		case evm.chainRules.IsEIP158:
// 			jt = spuriousDragonInstructionSet
// 		case evm.chainRules.IsEIP150:
// 			jt = tangerineWhistleInstructionSet
// 		case evm.chainRules.IsHomestead:
// 			jt = homesteadInstructionSet
// 		default:
// 			jt = frontierInstructionSet
// 		}
// 		for i, eip := range cfg.ExtraEips {
// 			if err := EnableEIP(eip, &jt); err != nil {
// 				// Disable it, so caller can check if it's activated or not
// 				cfg.ExtraEips = append(cfg.ExtraEips[:i], cfg.ExtraEips[i+1:]...)
// 				log.Error("EIP activation failed", "eip", eip, "error", err)
// 			}
// 		}
// 		cfg.JumpTable = jt
// 	}

// 	fmt.Println("file: intepreter.go \t func: NewEVMInterpreter, \t  Descr: the field of evm and cfg are as below:")

// 	fmt.Println(evm)
// 	fmt.Println(cfg)
// 	return &EVMInterpreter{
// 		evm: evm,
// 		cfg: cfg,
// 	}
// }

// // Run loops and evaluates the contract's code with the given input data and returns
// // the return byte-slice and an error if one occurred.
// //
// // It's important to note that any errors returned by the interpreter should be
// // considered a revert-and-consume-all-gas operation except for
// // ErrExecutionReverted which means revert-and-keep-gas-left.
// func (in *EVMInterpreter) Run(contract *Contract, input []byte, readOnly bool) (ret []byte, err error) {

// 	fmt.Println("file: intepreter.go \t func: Run, \t  Descr: Run loops and evaluates the contract's code with the given input data and returns the return byte-slice and an error if one occurred. current evm depth value is below:")

// 	fmt.Println(in.evm.depth)
// 	// Increment the call depth which is restricted to 1024

// 	in.evm.depth++
// 	defer func() { in.evm.depth-- }()

// 	// Make sure the readOnly is only set if we aren't in readOnly yet.
// 	// This makes also sure that the readOnly flag isn't removed for child calls.
// 	if readOnly && !in.readOnly {
// 		in.readOnly = true
// 		defer func() { in.readOnly = false }()
// 	}

// 	fmt.Println("file: intepreter.go \t func: Run, \t  Descr: return data of previous call is given below in.returndata variable. Here this value is set to null thereafter so that not to preserve old buffer, return data is: ")
// 	fmt.Println(in.returnData)

// 	// Reset the previous call's return data. It's unimportant to preserve the old buffer
// 	// as every returning call will return new data anyway.
// 	in.returnData = nil

// 	fmt.Println("file: intepreter.go \t func: Run, \t  Descr:  check the contract code length. if no code then return safely with nill. the length is as belw:")
// 	fmt.Println(len(contract.Code))

// 	// Don't bother with the execution if there's no code.
// 	if len(contract.Code) == 0 {
// 		return nil, nil
// 	}

// 	var (
// 		op          OpCode             // current opcode
// 		mem         = NewMemory()      // bound memory
// 		stack       = newstack()       // local stack
// 		returns     = newReturnStack() // local returns stack
// 		callContext = &callCtx{
// 			memory:   mem,
// 			stack:    stack,
// 			rstack:   returns,
// 			contract: contract,
// 		}
// 		// For optimisation reason we're using uint64 as the program counter.
// 		// It's theoretically possible to go above 2^64. The YP defines the PC
// 		// to be uint256. Practically much less so feasible.
// 		pc   = uint64(0) // program counter
// 		cost uint64
// 		// copies used by tracer
// 		pcCopy  uint64 // needed for the deferred Tracer
// 		gasCopy uint64 // for Tracer to log gas remaining before execution
// 		logged  bool   // deferred Tracer should ignore already logged steps
// 		//res     []byte // result of the opcode execution function
// 	)

// 	fmt.Println("file: intepreter.go \t func: Run, \t  Descr: newly created variables are as below:")

// 	fmt.Println("op")
// 	fmt.Println(op)

// 	fmt.Println("mem")
// 	fmt.Println(mem)

// 	fmt.Println("stack")
// 	fmt.Println(stack)

// 	fmt.Println("returns")
// 	fmt.Println(returns)

// 	fmt.Println("callContext")
// 	fmt.Println(callContext)

// 	// Don't move this deferrred function, it's placed before the capturestate-deferred method,
// 	// so that it get's executed _after_: the capturestate needs the stacks before
// 	// they are returned to the pools
// 	defer func() {

// 		fmt.Println("returning from contract execution, values of returnStack and returnRStack are as below")

// 		returnStack(stack)
// 		returnRStack(returns)

// 	}()
// 	contract.Input = input

// 	if in.cfg.Debug {
// 		defer func() {
// 			if err != nil {
// 				if !logged {
// 					in.cfg.Tracer.CaptureState(in.evm, pcCopy, op, gasCopy, cost, mem, stack, returns, in.returnData, contract, in.evm.depth, err)
// 				} else {
// 					in.cfg.Tracer.CaptureFault(in.evm, pcCopy, op, gasCopy, cost, mem, stack, returns, contract, in.evm.depth, err)
// 				}
// 			}
// 		}()
// 	}
// 	// The Interpreter main run loop (contextual). This loop runs until either an
// 	// explicit STOP, RETURN or SELFDESTRUCT is executed, an error occurred during
// 	// the execution of one of the operations or until the done flag is set by the
// 	// parent context.

// 	modifiedpc := make(chan uint64)

// 	steps := 0
// 	for {

// 		steps++

// 		fmt.Println("steps")
// 		fmt.Println(steps)

// 		if steps%1000 == 0 && atomic.LoadInt32(&in.evm.abort) != 0 {
// 			break
// 		}
// 		if in.cfg.Debug {
// 			// Capture pre-execution values for tracing.
// 			logged, pcCopy, gasCopy = false, pc, contract.Gas
// 		}

// 		// Get the operation from the jump table and validate the stack to ensure there are
// 		// enough stack items available to perform the operation.

// 		//temppc := pc

// 		//fetch and decode operation

// 		operation, conditionfetch, fetchedresult, err := fetch(modifiedpc, pc, contract, in, callContext, cost, pcCopy, gasCopy, mem, returns, logged)
// 		if !conditionfetch {
// 			return fetchedresult, err

// 		}

// 		//execute the operation

// 		mpc, conditionexec, executedresult, errs := execute(pc, in, callContext, operation)
// 		if !conditionexec {
// 			return executedresult, errs
// 		}

// 		pc = mpc

// 	}
// 	return nil, nil
// }

// // CanRun tells if the contract, passed as an argument, can be
// // run by the current interpreter.
// func (in *EVMInterpreter) CanRun(code []byte) bool {

// 	fmt.Println("approve for contract to run just setting true")

// 	return true
// }

// func fetch(modifiedpc <-chan uint64,  pc uint64, contract *Contract, in *EVMInterpreter, callContext *callCtx, cost uint64, pcCopy uint64, gasCopy uint64, mem *Memory, returns *ReturnStack, logged bool) (operation *operation, condition bool, ret []byte, errs error) {

// 	go func() {
// 		op := contract.GetOp(pc)

// 		fmt.Println("opcode fetched")
// 		fmt.Println(op)

// 		operation = in.cfg.JumpTable[op]

// 		fmt.Println("operation variable is set with opcode using in.cfg.jumptable. Operation detail is given below")
// 		fmt.Println(operation)

// 		if operation == nil {
// 			return nil, false, nil, &ErrInvalidOpCode{opcode: op}
// 		}

// 		fmt.Println("Stack length is :")
// 		fmt.Println(callContext.stack.len())

// 		fmt.Println("minStack")
// 		fmt.Println(operation.minStack)

// 		// Validate stack
// 		if sLen := callContext.stack.len(); sLen < operation.minStack {
// 			return nil, false, nil, &ErrStackUnderflow{stackLen: sLen, required: operation.minStack}

// 		} else if sLen > operation.maxStack {

// 			return nil, false, nil, &ErrStackOverflow{stackLen: sLen, limit: operation.maxStack}
// 		}

// 		// If the operation is valid, enforce and write restrictions
// 		if in.readOnly && in.evm.chainRules.IsByzantium {
// 			// If the interpreter is operating in readonly mode, make sure no
// 			// state-modifying operation is performed. The 3rd stack item
// 			// for a call operation is the value. Transferring value from one
// 			// account to the others means the state is modified and should also
// 			// return with an error.
// 			if operation.writes || (op == CALL && callContext.stack.Back(2).Sign() != 0) {
// 				return nil, false, nil, ErrWriteProtection
// 			}
// 		}
// 		// Static portion of gas
// 		cost = operation.constantGas // For tracing
// 		if !contract.UseGas(operation.constantGas) {
// 			return nil, false, nil, ErrOutOfGas
// 		}

// 		var memorySize uint64
// 		// calculate the new memory size and expand the memory to fit
// 		// the operation
// 		// Memory check needs to be done prior to evaluating the dynamic gas portion,
// 		// to detect calculation overflows
// 		if operation.memorySize != nil {
// 			memSize, overflow := operation.memorySize(callContext.stack)
// 			if overflow {
// 				return nil, false, nil, ErrGasUintOverflow
// 			}
// 			// memory is expanded in words of 32 bytes. Gas
// 			// is also calculated in words.
// 			if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
// 				return nil, false, nil, ErrGasUintOverflow
// 			}
// 		}
// 		// Dynamic portion of gas
// 		// consume the gas and return an error if not enough gas is available.
// 		// cost is explicitly set so that the capture state defer method can get the proper cost
// 		var err error
// 		if operation.dynamicGas != nil {
// 			var dynamicCost uint64

// 			dynamicCost, err = operation.dynamicGas(in.evm, contract, callContext.stack, mem, memorySize)
// 			cost += dynamicCost // total cost, for debug tracing
// 			if err != nil || !contract.UseGas(dynamicCost) {
// 				return nil, false, nil, ErrOutOfGas
// 			}
// 		}
// 		if memorySize > 0 {
// 			mem.Resize(memorySize)
// 		}

// 		if in.cfg.Debug {
// 			in.cfg.Tracer.CaptureState(in.evm, pc, op, gasCopy, cost, mem, callContext.stack, returns, in.returnData, contract, in.evm.depth, err)
// 			logged = true
// 		}

// 		fmt.Println("address of pc")
// 		fmt.Println(&pc)

// 		fmt.Println("interpreter context")
// 		fmt.Println(in)

// 		fmt.Println("callContext")
// 		fmt.Println(callContext)
// 	}()

// 	//operation variable is now passed to the execute phase
// 	return operation, true, nil, nil
// }

// func execute(pc uint64, in *EVMInterpreter, callContext *callCtx, operation *operation) (mpc uint64, condition bool, ret []byte, err error) {

// 	go func() {
// 		// execute the operation
// 		res, err := operation.execute(&pc, in, callContext)

// 		//end of execute and store
// 		if err != nil {
// 			fmt.Println("Error on intepreter execute operation")
// 			fmt.Println(err)
// 		}
// 		fmt.Println("execution result is")
// 		fmt.Println(res)

// 		// fmt.Println("execution result is")
// 		// fmt.Println(res)

// 		// if the operation clears the return data (e.g. it has returning data)
// 		// set the last return to the result of the operation.
// 		if operation.returns {
// 			in.returnData = common.CopyBytes(res)
// 		}

// 		switch {
// 		case err != nil:
// 			return pc, false, nil, err
// 		case operation.reverts:
// 			return pc, false, res, ErrExecutionReverted
// 		case operation.halts:
// 			fmt.Println("halted with result as:")
// 			fmt.Println(res)

// 			return pc, false, res, nil
// 		case !operation.jumps:
// 			fmt.Println("no jump operation, and pc before and after")
// 			fmt.Println(pc)

// 			pc++

// 			fmt.Println(pc)

// 		}
// 		modifiedpc <- pc
// 	}()

// 	return 1, true, nil, nil

// 	//one case is added here for if operation.jumps:
// 	//case operation.jumps:
// 	// set pc of phase 1 to current  phase 2 jumpdest pc
// }
