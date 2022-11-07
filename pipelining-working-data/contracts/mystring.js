// // var Web3 = require('web3')
// // var solc = require('solc')
// // var fs = require('fs')
//  const Web3 = require('web3'); 
// //  let web3 = new Web3(Web3.givenProvider || "127.0.0.1:8551");
// var web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));

// // var storageContract = new web3.eth.Contract([{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]);
// // var storage = storageContract.deploy({
// //      data: '0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212203e95ea2da4e939dfa48418cfc087fcbe688196e94ac891ac6a89ecd0d7fca55164736f6c63430008070033', 
// //      arguments: [
// //      ]
// // }).send({
// //      from: web3.eth.accounts[0], 
// //      gas: '4700000'
// //    }, function (e, contract){
// //     console.log(e, contract);
// //     if (typeof contract.address !== 'undefined') {
// //          console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
// //     }
// //  }
var Web3 = require('web3');
var web3 = new Web3('http://localhost:8545');
// or
// var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));

// change provider
// web3.setProvider('ws://localhost:8546');
// or
// web3.setProvider(new Web3.providers.WebsocketProvider('ws://localhost:8546'));

// Using the IPC provider in node.js
// var net = require('net');
// var web3 = new Web3('/Users/myuser/Library/Ethereum/geth.ipc', net); // mac os path
// // or
// var web3 = new Web3(new Web3.providers.IpcProvider('/Users/myuser/Library/Ethereum/geth.ipc', net)); // mac os path
// on windows the path is: "\\\\.\\pipe\\geth.ipc"
// on linux the path is: "/users/myuser/.ethereum/geth.ipc"
console.log(web3.givenProvider)