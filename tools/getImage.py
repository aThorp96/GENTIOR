import matplotlib.pyplot as plt
filename = "../data/randomDataset/points.txt"
edgeFile = open(filename)

xVals= []
yVals= []
maxX = -9999
minX = 9999

minY = 9999
maxY = -9999

solution = [97, 65, 84, 60, 27, 57, 68, 44, 29, 79, 92, 16, 58, 14, 95, 38, 59, 98, 48, 61, 72, 39, 80, 51, 87, 21, 70, 89, 49, 17, 93, 8, 99, 90, 3, 31, 5, 56, 32, 45, 22, 33, 13, 50, 24, 0, 75, 91, 30, 26, 86, 42, 36, 85, 6, 63, 9, 18, 67, 76, 4, 7, 35, 10, 47, 94, 66, 1, 96, 83, 19, 2, 20, 52, 53, 25, 28, 82, 81, 55, 43, 69, 74, 78, 37, 23, 54, 41, 15, 77, 73, 46, 64, 88, 12, 11, 40, 34, 71, 62]

for line in edgeFile:
    if not line.startswith("#"):
        words = line.split()
        x = float(words[0])
        y = float(words[1])

        if x < minX:
            minX = x
        elif x > maxX:
            maxX = x

        if y < minY:
            minY = y
        elif y > maxY:
            maxY = y

        xVals.append(x)
        yVals.append(y)

pathX = []
pathY = []
for i in solution:
    pathX.append(xVals[i])
    pathY.append(yVals[i])


#plt.plot(xVals, yVals, markers)
for x, y, i in zip(xVals, yVals, range(0, 128)):
    plt.text(x, y, i, color="blue", fontsize=10)
plt.plot(pathX, pathY, color = "red")
plt.axis([minX, maxX, minY, maxY])
plt.show()
