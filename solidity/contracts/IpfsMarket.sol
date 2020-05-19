pragma solidity 0.5.17;

contract IpfsMarket {

    struct Publication {
        address author;
        bytes authorAccessKey;
        uint256 price;
    }

    mapping (string => Publication) publications;

    function publish(
        string memory cid,
        bytes memory accessKey,
        uint256 price
    ) public {
        require(
            publications[cid].author == address(0),
            "Publication with given CID already exists"
        );

        publications[cid] = Publication(msg.sender, accessKey, price);
    }
}
