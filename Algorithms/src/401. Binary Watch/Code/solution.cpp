class Solution {
public:
    vector<string> readBinaryWatch(int turnedOn) {
        vector<string> result;
        
        // Try all possible hours (0-11)
        for (int hour = 0; hour < 12; hour++) {
            // Try all possible minutes (0-59)
            for (int minute = 0; minute < 60; minute++) {
                
                // Count total number of set bits in hour + minute
                if (__builtin_popcount(hour) + __builtin_popcount(minute) == turnedOn) {
                    
                    // Format time string
                    string time = to_string(hour) + ":";
                    
                    // Add leading zero for minute if needed
                    if (minute < 10) {
                        time += "0";
                    }
                    
                    time += to_string(minute);
                    
                    result.push_back(time);
                }
            }
        }
        
        return result;
    }
};
