import random
import math
outputPoints = open("points.txt", "w")
outputData = open("graph.dat", "w")

numPoints = 100
maxval = 1000.
xVals = []
yVals = []

for i in range(0, numPoints):
    x = random.uniform(0, maxval)
    y = random.uniform(0, maxval)
    outputPoints.write(str(x) + " " + str(y) + "\n")
    xVals.append(x)
    yVals.append(y)

outputData.write(str(numPoints) + "\n")
for x1, y1, i in zip(xVals, yVals, range(0, numPoints)):
    for x2, y2, j in zip(xVals, yVals, range(0, numPoints)):
        distance = math.hypot(x2 - x1, y2 - y1)
        outputData.write(str(i) + " " + str(j) + " " + str(int(distance)) + "\n")
outputData.write("-1 -1 -1\n")
