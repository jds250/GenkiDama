// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.22 < 0.9.0;
/**
* @title Storage
* @dev store or retrieve variable value
*/

contract Show {
    string public showStr;

    constructor() public {
        showStr = "show the detail";
    }

    function setShowStr(string memory _newStr) public returns(bool, string memory) {
        showStr = _newStr;
        return (true, _newStr);
    }

    function show() public view returns (string memory){
        return showStr;
    }
}