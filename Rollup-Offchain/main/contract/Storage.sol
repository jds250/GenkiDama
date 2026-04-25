// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.22 < 0.9.0;
/**
* @title Storage
* @dev store or retrieve variable value
*/

import {Show} from "./Show.sol";

contract Storage {

    uint256 public value;
    address public showAddr;
    Show showInstance;

    event ABIEvent(bool success, bytes returnData);

    constructor (address _showAddr) public {
        showAddr = _showAddr;
        showInstance = Show(_showAddr);
    }

    function store(uint256 number) public{
        value = number;
    }

    function retrieve() public view returns (uint256){
        return value;
    }

    function setShowStr(string memory _str) public returns(bool, string memory) {
        bool success;
        bytes memory returnData;

        (success, returnData) = address(showInstance).call(
            abi.encodeWithSelector(
                showInstance
                .setShowStr
                .selector,
                _str
            )
        );

        bool res1;
        string memory res2;
        // If the call was successful let's decode!
        if (success) {
            (res1, res2) = abi.decode(
                (returnData),
                (bool, string)
            );
        } else {
            return (false, "call fail !!!");
        }

        bytes memory resABI = abi.encode(res1, res2);

        emit ABIEvent(true, resABI);
        return (res1, res2);
    }

    function getShowAddr() public view returns (address) {
        return showAddr;
    }

    function getShowRes() public view returns (string memory) {
        return showInstance.show();
    }

    function getInfo() public view returns (bytes memory) {
        string memory tmp = showInstance.show();
        bytes memory res = abi.encode(
            showAddr,
            tmp
        );

        return res;
    }
}