// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function transfer(address recipient, uint256 amount) external returns (bool);
}

contract Staking {
    IERC20 public token;
    mapping(address => uint256) public staked;
    mapping(address => uint256) public stakingTime;

    constructor(address _tokenAddress) {
        token = IERC20(_tokenAddress);
    }

    function stake(uint256 amount) external {
        require(amount > 0, "amount > 0");
        token.transferFrom(msg.sender, address(this), amount);
        staked[msg.sender] += amount;
        stakingTime[msg.sender] = block.timestamp;
    }

    function unstake() external {
        uint256 balance = staked[msg.sender];
        require(balance > 0, "No staked tokens");
        require(block.timestamp >= stakingTime[msg.sender] + 1 days, "Staking period not finished");
        staked[msg.sender] = 0;
        token.transfer(msg.sender, balance);
    }

    function getStaked(address user) external view returns (uint256) {
        return staked[user];
    }
}
