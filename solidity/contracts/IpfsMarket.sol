pragma solidity 0.6.2;

import "@openzeppelin/contracts/payment/PullPayment.sol";

contract IpfsMarket is PullPayment {

    struct Publication {
        address author;
        bytes authorAccessKey;
        uint256 price;

        address[] purchasers;
        mapping (address => bytes) purchasersPublicKeys;
        mapping (address => bytes) purchasersAccessKeys;
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

        Publication memory publication;
        publication.author = msg.sender;
        publication.authorAccessKey = accessKey;
        publication.price = price;

        publications[cid] = publication;
    }

    function getPrice(string memory cid) public view returns (uint256) {
        require(
            publications[cid].author != address(0),
            "Publication with given CID doesn't exists"
        );

        return publications[cid].price;
    }

    function purchase(
        string memory cid,
        bytes memory publicKey
    ) public payable {
        require(
            publications[cid].author != address(0),
            "Publication with given CID doesn't exists"
        );

        require(
            publications[cid].author != msg.sender,
            "Only non-authors can purchase"
        );

        address purchaser = publicKeyToAddress(publicKey);

        require(
            purchaser == msg.sender,
            "Sender must use their own public key"
        );

        require(
            publications[cid].purchasersPublicKeys[purchaser].length == 0,
            "Purchase can be made only once"
        );

        require(
            publications[cid].price == msg.value,
            "Incorrect payment amount"
        );

        publications[cid].purchasers.push(purchaser);
        publications[cid].purchasersPublicKeys[purchaser] = publicKey;
    }

    function publicKeyToAddress (
        bytes memory publicKey
    ) private pure returns (address) {
        require (
            publicKey.length == 64,
            "Incorrect public key length"
        );

        return address(bytes20(uint160(uint256(keccak256(publicKey)))));
    }

    function answerPurchase(
        string memory cid,
        address purchaser,
        bytes memory accessKey
    ) public {
        require(
            publications[cid].author != address(0),
            "Publication with given CID doesn't exists"
        );

        require(
            publications[cid].author == msg.sender,
            "Only publication author can answer to the purchase"
        );

        require(
            publications[cid].purchasersPublicKeys[purchaser].length != 0,
            "Could not answer a non existing purchase"
        );

        require(
            publications[cid].purchasersAccessKeys[purchaser].length == 0,
            "Purchase answer can be made only once"
        );

        publications[cid].purchasersAccessKeys[purchaser] = accessKey;

        _asyncTransfer(publications[cid].author, publications[cid].price);
    }

    function hasPurchased(string memory cid) public view returns (bool) {
        return publications[cid].purchasersPublicKeys[msg.sender].length > 0;
    }

    function getAccessKey(string memory cid) public view returns (bytes memory) {
        require(
            publications[cid].author != address(0),
            "Publication with given CID doesn't exists"
        );

        return publications[cid].purchasersAccessKeys[msg.sender];
    }
}
