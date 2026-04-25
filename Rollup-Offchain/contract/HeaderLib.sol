// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

library HeaderLib {
    function canonicalHash(bytes memory _header) internal pure returns (bytes32) {
        if (_header.length == 32) {
            return _decodeBytes32(_header);
        }

        return keccak256(_header);
    }

    function tryParentHash(bytes memory _header) internal pure returns (bool, bytes32) {
        return _tryBytes32Field(_header, 0);
    }

    function tryStateRoot(bytes memory _header) internal pure returns (bool, bytes32) {
        return _tryBytes32Field(_header, 3);
    }

    function _tryBytes32Field(bytes memory _header, uint256 _index) private pure returns (bool, bytes32) {
        if (_header.length == 32 || _header.length == 0 || uint8(_header[0]) < 0xc0) {
            return (false, bytes32(0));
        }

        (uint256 cursor, uint256 payloadEnd) = _listPayloadBounds(_header);
        for (uint256 i = 0; i < _index; i++) {
            if (cursor >= payloadEnd) {
                return (false, bytes32(0));
            }
            cursor += _itemLength(_header, cursor);
        }

        if (cursor >= payloadEnd) {
            return (false, bytes32(0));
        }

        (uint256 itemOffset, uint256 itemLengthValue, bool isList) = _payloadBounds(_header, cursor);
        if (isList || itemLengthValue != 32 || itemOffset + itemLengthValue > _header.length) {
            return (false, bytes32(0));
        }

        bytes32 value;
        assembly {
            value := mload(add(add(_header, 32), itemOffset))
        }
        return (true, value);
    }

    function _listPayloadBounds(bytes memory _header) private pure returns (uint256, uint256) {
        (uint256 payloadOffset, uint256 payloadLength, bool isList) = _payloadBounds(_header, 0);
        require(isList, "header must be an rlp list");
        return (payloadOffset, payloadOffset + payloadLength);
    }

    function _itemLength(bytes memory _header, uint256 _offset) private pure returns (uint256) {
        (uint256 payloadOffset, uint256 payloadLength, ) = _payloadBounds(_header, _offset);
        return (payloadOffset - _offset) + payloadLength;
    }

    function _payloadBounds(
        bytes memory _header,
        uint256 _offset
    ) private pure returns (uint256 payloadOffset, uint256 payloadLength, bool isList) {
        require(_offset < _header.length, "rlp item out of bounds");
        uint8 prefix = uint8(_header[_offset]);

        if (prefix <= 0x7f) {
            return (_offset, 1, false);
        }
        if (prefix <= 0xb7) {
            return (_offset + 1, prefix - 0x80, false);
        }
        if (prefix <= 0xbf) {
            uint256 lenOfLen = prefix - 0xb7;
            return (_offset + 1 + lenOfLen, _readBigEndian(_header, _offset + 1, lenOfLen), false);
        }
        if (prefix <= 0xf7) {
            return (_offset + 1, prefix - 0xc0, true);
        }

        uint256 listLenOfLen = prefix - 0xf7;
        return (_offset + 1 + listLenOfLen, _readBigEndian(_header, _offset + 1, listLenOfLen), true);
    }

    function _readBigEndian(bytes memory _header, uint256 _offset, uint256 _length) private pure returns (uint256 value) {
        require(_offset + _length <= _header.length, "rlp length out of bounds");
        for (uint256 i = 0; i < _length; i++) {
            value = (value << 8) | uint8(_header[_offset + i]);
        }
    }

    function _decodeBytes32(bytes memory _value) private pure returns (bytes32 result) {
        require(_value.length == 32, "invalid bytes32");
        assembly {
            result := mload(add(_value, 32))
        }
    }
}
