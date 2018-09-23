# Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
# The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
# You may assume the integer does not contain any leading zero, except the number 0 itself.
# https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/

class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        string_list = [str(i) for i in digits]
        strings = str(int("".join(string_list)) + 1)
        return [int(var) for var in strings]
