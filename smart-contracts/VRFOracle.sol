// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@chainlink/contracts/src/v0.8/VRFConsumerBase.sol";

contract VRFOracle is VRFConsumerBase {
    bytes32 internal keyHash;
    uint256 internal fee;

    uint256 public randomResult;
    address public lastRequester;

    mapping(bytes32 => address) public requestToSender;

    event RandomRequested(bytes32 requestId, address requester);
    event RandomFulfilled(bytes32 requestId, uint256 randomness);

    constructor(
        address vrfCoordinator,
        address linkToken,
        bytes32 _keyHash,
        uint256 _fee
    ) VRFConsumerBase(vrfCoordinator, linkToken) {
        keyHash = _keyHash;
        fee = _fee; // Example: 0.1 * 10 ** 18 for 0.1 LINK
    }

    function requestRandomNumber() public returns (bytes32 requestId) {
        require(LINK.balanceOf(address(this)) >= fee, "Not enough LINK");
        requestId = requestRandomness(keyHash, fee);
        requestToSender[requestId] = msg.sender;
        emit RandomRequested(requestId, msg.sender);
    }

    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
        randomResult = randomness;
        lastRequester = requestToSender[requestId];
        emit RandomFulfilled(requestId, randomness);
    }

    // Helper function to withdraw LINK in case of emergency
    function withdrawLink() external {
        require(LINK.transfer(msg.sender, LINK.balanceOf(address(this))), "Transfer failed");
    }
}
