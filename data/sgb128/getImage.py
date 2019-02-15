import matplotlib.pyplot as plt
filename = "./sgb128_xy.txt"
edgeFile = open(filename)

xVals= []
yVals= []
maxX = -9999
minX = 9999

minY = 9999
maxY = -9999

solution = [20, 51, 64, 14, 117, 92, 116, 68, 70, 47, 69, 79, 53, 91, 108, 88, 103, 60, 106, 119, 63, 95, 21, 71, 46, 58, 50, 65, 66, 40, 26, 0, 93, 3, 74, 105, 5, 82, 73, 22, 101, 72, 6, 12, 42, 27, 15, 96, 104, 113, 124, 16, 112, 120, 52, 35, 18, 55, 43, 87, 109, 44, 111, 29, 37, 123, 33, 48, 49, 78, 83, 62, 39, 125, 1, 89, 97, 59, 80, 126, 9, 77, 28, 25, 54, 115, 90, 2, 8, 30, 57, 81, 102, 31, 84, 107, 86, 41, 10, 11, 7, 45, 19, 76, 34, 67, 36, 56, 99, 23, 61, 122, 75, 32, 128, 121, 127, 85, 100, 13, 114, 17, 110, 24, 38, 98, 4, 118, 94]

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
print(len(xVals))
for i in solution:
    print(i)
    pathX.append(xVals[i])
    pathY.append(yVals[i])


plt.plot(pathX, pathY, color = "red")
plt.plot(xVals, yVals, 'bo')
plt.axis([minX, maxX, minY, maxY])
plt.show()
