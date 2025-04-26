// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DAOVoting {
    struct Proposal {
        string description;
        uint256 voteCount;
        bool executed;
    }

    Proposal[] public proposals;
    mapping(address => bool) public voted;

    function createProposal(string memory desc) public {
        proposals.push(Proposal({ description: desc, voteCount: 0, executed: false }));
    }

    function vote(uint256 proposalId) public {
        require(!voted[msg.sender], "Already voted");
        require(proposalId < proposals.length, "Invalid proposal");
        proposals[proposalId].voteCount++;
        voted[msg.sender] = true;
    }

    function execute(uint256 proposalId) public {
        Proposal storage proposal = proposals[proposalId];
        require(!proposal.executed, "Already executed");
        require(proposal.voteCount > 0, "Not enough votes");
        proposal.executed = true;
        // 행동 실행은 오프체인 또는 확장 필요
    }

    function getProposal(uint256 id) public view returns (string memory, uint256, bool) {
        Proposal memory p = proposals[id];
        return (p.description, p.voteCount, p.executed);
    }

    function getProposalCount() public view returns (uint256) {
        return proposals.length;
    }
}
