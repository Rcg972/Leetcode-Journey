class Solution {
public:
    int largestRectangleArea(vector<int>& heights) {
        vector<int>stack;
        int n = heights.size();
        int max_luas = 0;

        for (int i = 0; i < n ; i++){
            while (!stack.empty() && heights[i] < heights[stack.back()]){
                int top = stack.back();
                int tinggi = heights[top];

                stack.pop_back();

                int lebar;

                if (stack.empty()){
                    lebar = i;
                }else{
                    int kiri = stack.back();
                    lebar = i - kiri - 1;
                }

                int curr_luas = tinggi * lebar;
                max_luas = max(max_luas, curr_luas);
            }

            stack.push_back(i);
        }

        while (!stack.empty()){
                int top = stack.back();
                int tinggi = heights[top];

                stack.pop_back();

                int lebar;

                if (stack.empty()){
                    lebar = n;
                }else{
                    int kiri = stack.back();
                    lebar = n - kiri - 1;
                }

                int curr_luas = tinggi * lebar;
                max_luas = max(max_luas, curr_luas);
            }

        return max_luas;
    }
};
