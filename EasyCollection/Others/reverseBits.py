# Reverse bits of a given 32 bits unsigned integer.
# https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/648/
class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        binary = bin(n).replace("0b", "")
        remain = 32 - len(binary)
        binary = remain * "0" + binary
        return int(binary[::-1], 2)
