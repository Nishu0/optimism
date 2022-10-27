// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":25185,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":28178,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":28181,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":28792,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":28050,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":28170,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":28343,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":28448,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":28463,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":28507,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":25237,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":25242,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":25247,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":25250,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":25253,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":25258,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"receivedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":25263,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80638129fc1c116100cb578063b28ade251161007f578063ecc7042811610059578063ecc70428146103f6578063f2fde38b1461045b578063f69f81511461047b57600080fd5b8063b28ade251461038f578063d764ad0b146103af578063db505d80146103c257600080fd5b80638da5cb5b116100b05780638da5cb5b14610301578063a71198691461032c578063b1b1b2091461035f57600080fd5b80638129fc1c146102d75780638456cb59146102ec57600080fd5b80633f827a5a1161012d5780636e296e45116101075780636e296e4514610271578063715018a6146102ab5780637dea7cc3146102c057600080fd5b80633f827a5a1461020357806354fd4d501461022b5780635c975abb1461024d57600080fd5b80632828d7e81161015e5780632828d7e8146101c35780633dbb202b146101d95780633f4ba83a146101ee57600080fd5b8063028f85f71461017a5780630c568498146101ad575b600080fd5b34801561018657600080fd5b5061018f601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101b957600080fd5b5061018f6103e881565b3480156101cf57600080fd5b5061018f6103f881565b6101ec6101e7366004611cec565b6104ab565b005b3480156101fa57600080fd5b506101ec61070f565b34801561020f57600080fd5b50610218600181565b60405161ffff90911681526020016101a4565b34801561023757600080fd5b50610240610721565b6040516101a49190611dcb565b34801561025957600080fd5b5060655460ff165b60405190151581526020016101a4565b34801561027d57600080fd5b506102866107c4565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a4565b3480156102b757600080fd5b506101ec6108b0565b3480156102cc57600080fd5b5061018f62030d4081565b3480156102e357600080fd5b506101ec6108c2565b3480156102f857600080fd5b506101ec610abf565b34801561030d57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610286565b34801561033857600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610286565b34801561036b57600080fd5b5061026161037a366004611de5565b60cb6020526000908152604090205460ff1681565b34801561039b57600080fd5b5061018f6103aa366004611dfe565b610acf565b6101ec6103bd366004611e52565b610b1b565b3480156103ce57600080fd5b506102867f000000000000000000000000000000000000000000000000000000000000000081565b34801561040257600080fd5b5061044d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a4565b34801561046757600080fd5b506101ec610476366004611ed4565b61121c565b34801561048757600080fd5b50610261610496366004611de5565b60ce6020526000908152604090205460ff1681565b6105e47f00000000000000000000000000000000000000000000000000000000000000006104da858585610acf565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061054660cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016105629796959493929190611f38565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526112ec565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561066960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161067b959493929190611f97565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b61071761137a565b61071f6113fb565b565b606061074c7f0000000000000000000000000000000000000000000000000000000000000000611478565b6107757f0000000000000000000000000000000000000000000000000000000000000000611478565b61079e7f0000000000000000000000000000000000000000000000000000000000000000611478565b6040516020016107b093929190611fe5565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6108b861137a565b61071f60006115ad565b6000547501000000000000000000000000000000000000000000900460ff161580801561090d575060005460017401000000000000000000000000000000000000000090910460ff16105b8061093f5750303b15801561093f575060005474010000000000000000000000000000000000000000900460ff166001145b6109cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161088a565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610a5157600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610a59611624565b8015610abc57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610ac761137a565b61071f61171b565b600062030d40610ae060108561208a565b6103e8610af56103f863ffffffff871661208a565b610aff91906120e9565b610b099190612110565b610b139190612110565b949350505050565b600260975403610b87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161088a565b6002609755610b94611776565b60f087901c60018114610c4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605560248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2031206d657373616765732061726520737570706f72746564206166746560648201527f722074686520426564726f636b20757067726164650000000000000000000000608482015260a40161088a565b6000610c95898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117e392505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610d2d57853414610d0957610d0961213c565b600081815260ce602052604090205460ff1615610d2857610d2861213c565b610e7f565b3415610de1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161088a565b600081815260ce602052604090205460ff16610e7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161088a565b610e8887611806565b15610f3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161088a565b600081815260cb602052604090205460ff1615610fda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161088a565b610fe661afc88661216b565b5a1015611075576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d657373616765000000000000000000606482015260840161088a565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556000611111886110c961138861afc8612183565b5a6110d49190612183565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061185b92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015156001036111ac57600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261120b565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a25b505060016097555050505050505050565b61122461137a565b73ffffffffffffffffffffffffffffffffffffffff81166112c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161088a565b610abc816115ad565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac9084906113429088908890879060040161219a565b6000604051808303818588803b15801561135b57600080fd5b505af115801561136f573d6000803e3d6000fd5b505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161088a565b611403611875565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6060816000036114bb57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156114e557806114cf816121e2565b91506114de9050600a8361221a565b91506114bf565b60008167ffffffffffffffff8111156115005761150061222e565b6040519080825280601f01601f19166020018201604052801561152a576020820181803683370190505b5090505b8415610b135761153f600183612183565b915061154c600a8661225d565b61155790603061216b565b60f81b81838151811061156c5761156c612271565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506115a6600a8661221a565b945061152e565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000547501000000000000000000000000000000000000000000900460ff166116cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790556117036118e1565b61170b61198c565b611713611a40565b61071f611b15565b611723611776565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861144e3390565b60655460ff161561071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161088a565b60006117f3878787878787611bc7565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611855575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080600080845160208601878a8af19695505050505050565b60655460ff1661071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161088a565b6000547501000000000000000000000000000000000000000000900460ff1661071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b6000547501000000000000000000000000000000000000000000900460ff16611a37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b61071f336115ad565b6000547501000000000000000000000000000000000000000000900460ff16611aeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000547501000000000000000000000000000000000000000000900460ff16611bc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b6001609755565b6060868686868686604051602401611be4969594939291906122a0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611c8a57600080fd5b919050565b60008083601f840112611ca157600080fd5b50813567ffffffffffffffff811115611cb957600080fd5b602083019150836020828501011115611cd157600080fd5b9250929050565b803563ffffffff81168114611c8a57600080fd5b60008060008060608587031215611d0257600080fd5b611d0b85611c66565b9350602085013567ffffffffffffffff811115611d2757600080fd5b611d3387828801611c8f565b9094509250611d46905060408601611cd8565b905092959194509250565b60005b83811015611d6c578181015183820152602001611d54565b83811115611d7b576000848401525b50505050565b60008151808452611d99816020860160208601611d51565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611dde6020830184611d81565b9392505050565b600060208284031215611df757600080fd5b5035919050565b600080600060408486031215611e1357600080fd5b833567ffffffffffffffff811115611e2a57600080fd5b611e3686828701611c8f565b9094509250611e49905060208501611cd8565b90509250925092565b600080600080600080600060c0888a031215611e6d57600080fd5b87359650611e7d60208901611c66565b9550611e8b60408901611c66565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611eb557600080fd5b611ec18a828b01611c8f565b989b979a50959850939692959293505050565b600060208284031215611ee657600080fd5b611dde82611c66565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611f8a60c083018486611eef565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611fc7608083018688611eef565b905083604083015263ffffffff831660608301529695505050505050565b60008451611ff7818460208901611d51565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612033816001850160208a01611d51565b6001920191820152835161204e816002840160208801611d51565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516818304811182151516156120b1576120b161205b565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680612104576121046120ba565b92169190910492915050565b600067ffffffffffffffff8083168185168083038211156121335761213361205b565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000821982111561217e5761217e61205b565b500190565b6000828210156121955761219561205b565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff831660208201526060604082015260006121d96060830184611d81565b95945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036122135761221361205b565b5060010190565b600082612229576122296120ba565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008261226c5761226c6120ba565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526122eb60c0830184611d81565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
