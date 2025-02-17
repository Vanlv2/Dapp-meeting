// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract LivestreamRoom {
    struct Livestream {
        string sessionId;
        string whipUrl;
        string whepUrl;
        string host;
        string status;
        uint256 createdAt;
    }

    mapping(string => Livestream) public livestreams;

    event LivestreamCreated(string roomId, string sessionId, string whipUrl, string host);
    event LivestreamEnded(string roomId, string sessionId);
    event WHEPUrlUpdated(string roomId, string whepUrl);

    // Tạo phòng livestream mới (sẽ được gọi bởi streamer)
    function createLivestream(string memory roomId, string memory sessionId, string memory whipUrl, string memory host) public {
        require(bytes(roomId).length > 0, "Room ID cannot be empty");
        require(bytes(sessionId).length > 0, "Session ID cannot be empty");
        require(bytes(whipUrl).length > 0, "WHIP URL cannot be empty");

        livestreams[roomId] = Livestream({
            sessionId: sessionId,
            whipUrl: whipUrl,
            whepUrl: "",
            host: host,
            status: "active",
            createdAt: block.timestamp
        });

        emit LivestreamCreated(roomId, sessionId, whipUrl, host);
    }

    // Cập nhật WHEP URL khi viewer tham gia (Cloudflare trả về WHEP URL)
    function updateWHEPUrl(string memory roomId, string memory whepUrl) public {
        Livestream storage livestream = livestreams[roomId];
        require(bytes(livestream.sessionId).length > 0, "Livestream does not exist");
        require(bytes(whepUrl).length > 0, "WHEP URL cannot be empty");

        livestream.whepUrl = whepUrl;

        emit WHEPUrlUpdated(roomId, whepUrl);
    }

    // Lấy thông tin phòng livestream (Viewer gọi để lấy session và URL)
    function getLivestreamInfo(string memory roomId) public view returns (string memory sessionId, string memory whepUrl, string memory host, string memory status, uint256 createdAt) {
        Livestream storage livestream = livestreams[roomId];
        require(bytes(livestream.sessionId).length > 0, "Livestream does not exist");

        return (livestream.sessionId, livestream.whepUrl, livestream.host, livestream.status, livestream.createdAt);
    }

    // Kết thúc livestream (Streamer gọi khi kết thúc phiên livestream)
    function endLivestream(string memory roomId) public {
        Livestream storage livestream = livestreams[roomId];
        require(bytes(livestream.sessionId).length > 0, "Livestream does not exist");

        livestream.status = "ended";

        emit LivestreamEnded(roomId, livestream.sessionId);
    }
}
