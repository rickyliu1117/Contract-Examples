// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Evidence {
    struct EvidenceData {
        bytes32 hash;
        address owner;
        bytes remaks;
        uint256 timestamp;
    }

    mapping(bytes32 => EvidenceData) private _evidences;

    function setData(
        bytes32 hash,
        address owner,
        bytes memory remaks,
        uint256 timestamp
    ) public {
        require(hash != 0, "Not valid hash");
        _evidences[hash].hash = hash;
        _evidences[hash].owner = owner;
        _evidences[hash].remaks = remaks;
        _evidences[hash].timestamp = timestamp;
    }

    function getData(bytes32 hash)
        public
        view
        returns (
            bytes32,
            address,
            bytes memory,
            uint256
        )
    {
        EvidenceData storage evidence = _evidences[hash];
        require(evidence.hash == hash, "Evidence not exist");
        return (
            evidence.hash,
            evidence.owner,
            evidence.remaks,
            evidence.timestamp
        );
    }
}
