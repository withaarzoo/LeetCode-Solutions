class Solution:
    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        result = []
        
        # Try all possible hours
        for hour in range(12):
            
            # Try all possible minutes
            for minute in range(60):
                
                # Count total set bits
                if (bin(hour).count('1') + bin(minute).count('1')) == turnedOn:
                    
                    # Format minute with leading zero
                    time = f"{hour}:{minute:02d}"
                    result.append(time)
        
        return result
