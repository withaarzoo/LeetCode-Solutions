class Solution:
    def countOdds(self, low: int, high: int) -> int:
        # Helper: count odd numbers from 1 to x
        def odds_up_to(x: int) -> int:
            # // is integer division (floor)
            return (x + 1) // 2
        
        # Odds in [low, high]
        return odds_up_to(high) - odds_up_to(low - 1)
