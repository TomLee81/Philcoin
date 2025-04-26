// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Escrow {
    address public buyer;
    address public seller;
    address public arbiter;

    uint public amount;
    bool public isReleased;

    constructor(address _buyer, address _seller, address _arbiter) payable {
        buyer = _buyer;
        seller = _seller;
        arbiter = _arbiter;
        amount = msg.value;
        isReleased = false;
    }

    function releaseFunds() external {
        require(msg.sender == buyer || msg.sender == arbiter, "Not authorized");
        require(!isReleased, "Already released");

        payable(seller).transfer(amount);
        isReleased = true;
    }

    function refundBuyer() external {
        require(msg.sender == arbiter, "Only arbiter");
        require(!isReleased, "Already released");

        payable(buyer).transfer(amount);
        isReleased = true;
    }
}
