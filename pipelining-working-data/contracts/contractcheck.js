var Web3 = require('web3');

// var solc = require('solc')
var fs = require('fs')

var abi = [
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "store",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "retrieve",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
var provider = 'http://localhost:8545';
var web3Provider = new Web3.providers.HttpProvider(provider);
var web3 = new Web3(web3Provider);
web3.eth.getBlockNumber().then((result) => {
  console.log("Latest Ethereum Block is ",result);
});


var storageContract = new web3.eth.Contract([{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]);

// var storage = storageContract.deploy({
//      data: '0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212204322126feb1ce8c38fa9d4dd9722adaaf9df7a40a8bccfc2a1d7d7a2b11f817464736f6c63430008070033', 
//      arguments: [
//      ]
// }).send({
//      from: "b6a076c3b9e36f63dfbc920dcb462207f771695b",
//      gas: '4700000'
//    }, function (e, contract){
//     console.log(e, contract);
    
//  }).on('receipt', function(receipt){
//   console.log("contract address is " + receipt.contractAddress)
//  })


var address = "0x75D32A43BDee4C4bAF29d9A631Eb3B09Bfb809ad"
var MyContract = new web3.eth.Contract(abi, address);

MyContract.methods.retrieve().call()
.then(console.log);

// MyContract.methods.store(19921103).send({from: 'b6a076c3b9e36f63dfbc920dcb462207f771695b'})
// .on('receipt', function(receipt){
//     // receipt example
//     console.log(receipt);
// })
