class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        bool isDup = false;

        unordered_set<int> hashset;
        for(int i = 0; i < nums.size(); i++) {
            if (hashset.find(nums[i]) != hashset.end()) {
                isDup = true;
                break;
            }

            hashset.insert(nums[i]);
            if(i >= k) {
                hashset.erase(nums[i - k]);
            }
        }
        return isDup;
    }
};