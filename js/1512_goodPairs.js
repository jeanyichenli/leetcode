/** @format */

/**
 * Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
 */

/**
 *
 * @param {number[]} nums
 * @returns {number}
 */

var numIdenticalPairs = function (nums) {
  let cnt = 0; // output count

  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] === nums[j]) {
        cnt++;
      }
    }
  }
  return cnt;
};
