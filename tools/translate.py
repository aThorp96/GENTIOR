outFilename = "sgb128_size128.dat"
outFile = open(outFilename, "w")

filename = "./sgb128_dist.txt"
edgeFile = open(filename)

string = ""

i = 0
for line in edgeFile:
    if not line.startswith("#"):
        j = 0
        for word in line.split():
           string += (str(i) + " " + str(j) + " " + word + "\n")
           j += 1
        i += 1
outFile.write(str(i) + "\n")
outFile.write(string)
outFile.write("-1 -1 -1")

