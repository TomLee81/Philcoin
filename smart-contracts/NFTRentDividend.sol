// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IERC721 {
    function transferFrom(address from, address to, uint256 tokenId) external;
}

contract NFTLease {
    struct Lease {
        address lessor;
        address lessee;
        uint256 tokenId;
        uint256 rentAmount;
        uint256 duration;
        uint256 startTime;
        bool active;
    }

    mapping(uint256 => Lease) public leases;
    IERC721 public nftContract;
    uint256 public leaseIdCounter;

    constructor(address _nftAddress) {
        nftContract = IERC721(_nftAddress);
    }

    function createLease(address lessee, uint256 tokenId, uint256 rentAmount, uint256 duration) public {
        nftContract.transferFrom(msg.sender, address(this), tokenId);
        leases[leaseIdCounter] = Lease(msg.sender, lessee, tokenId, rentAmount, duration, block.timestamp, true);
        leaseIdCounter++;
    }

    function endLease(uint256 leaseId) public {
        Lease storage lease = leases[leaseId];
        require(block.timestamp >= lease.startTime + lease.duration, "Lease not ended");
        require(lease.active, "Already inactive");
        lease.active = false;
        nftContract.transferFrom(address(this), lease.lessor, lease.tokenId);
    }

    function payRent(uint256 leaseId) public payable {
        Lease memory lease = leases[leaseId];
        require(msg.value == lease.rentAmount, "Incorrect rent");
        payable(lease.lessor).transfer(msg.value);
    }
}
