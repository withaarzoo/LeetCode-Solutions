function checkOverlap(radius: number, xCenter: number, yCenter: number, x1: number, y1: number, x2: number, y2: number): boolean {
    let closestX = Math.max(x1, Math.min(xCenter, x2));
    let closestY = Math.max(y1, Math.min(yCenter, y2));
    let dx = closestX - xCenter;
    let dy = closestY - yCenter;
    return dx * dx + dy * dy <= radius * radius;
};