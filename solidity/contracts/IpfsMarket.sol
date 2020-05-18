pragma solidity 0.5.17;

contract IpfsMarket {

    struct Publication {
        address author;
        bytes authorAccessKey;
        uint256 price;
    }

    mapping (string => Publication) publications;

    function publish(
        string memory hash,
        bytes memory accessKey,
        uint256 price
    ) public {
        require(
            publications[hash].author == address(0),
            "Publication with given hash already exists"
        );

        publications[hash] = Publication(msg.sender, accessKey, price);
    }
}
