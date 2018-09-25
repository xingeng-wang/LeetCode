# Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
# https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/565/
class Solution:
    # @param n, an integer
    # @return an integer
    def hammingWeight(self, n):
        return bin(n).count("1")
