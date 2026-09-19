class Solution {
public:
    bool checkOverlap(int radius, int xCenter, int yCenter, int x1, int y1, int x2, int y2) {
        int closestX = max(x1, min(xCenter, x2));
        int closestY = max(y1, min(yCenter, y2));
        int dx = closestX - xCenter;
        int dy = closestY - yCenter;
        return (long long)dx * dx + (long long)dy * dy <= (long long)radius * radius;
    }
};