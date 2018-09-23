class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        index_dict = {}
        for i in range(0, len(nums)):
            if target - nums[i] in index_dict:
                return [index_dict[target - nums[i]], i]
            else:
                index_dict[nums[i]] = i
