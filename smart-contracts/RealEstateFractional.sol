// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract RealEstateFractional is ERC721Enumerable, Ownable {
	uint256 public currentId = 0;
	mapping(uint256 => string) public propertyMetadata;

	constructor() ERC721("RealEstateFraction", "REF") {}

	function mintProperty(address to, string memory metadata) public onlyOwner {
		currentId++;
		_safeMint(to, currentId);
		propertyMetadata[currentId] = metadata;
	}

	function getMetadata(uint256 tokenId) public view returns (string memory) {
		require(_exists(tokenId), "Token does not exist");
		return propertyMetadata[tokenId];
	}
} // End of Smart Contract