import math

outFilename = "./xqf131.wdat"
outFile = open(outFilename, "w")

filename = "./xqf131.tsp"
edgeFile = open(filename)

string = ""

numPoints = 0
xVals = []
yVals = []
beganPoints = False

# Read in file
for line in edgeFile:
    if not beganPoints and line.startswith("DIMENSION"):
        words = line.split()
        numPoints = int(words[2])
        print("there are " + str(numPoints) + "words")
    elif line.startswith("NODE_COORD_SECTION"):
        beganPoints = True
    elif beganPoints and not line.startswith("EOF"):
        words = line.split()
        # First number on line is the node number,
        # and frankly the nodes seem to be
        # 1-indexed so I don't care
        xVals.append(int(words[1]))
        yVals.append(int(words[2]))

# Print order of graph
outFile.write(str(numPoints) + "\n")

# print weighted data
for x1, y1, i in zip(xVals, yVals, range(0, numPoints)):
    for x2, y2, j in zip(xVals, yVals, range(0, numPoints)):
        distance = math.hypot(x2 - x1, y2 - y1)
        outFile.write(str(i) + " " + str(j) + " " + str(int(distance)) + "\n")

#print footer
outFile.write("-1 -1 -1")
