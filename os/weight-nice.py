import matplotlib.pyplot as plt
import math


def get_weights(nices):
    weights = []
    for i, _ in enumerate(nices):
        weights.append(1024 / (math.pow(1.25, nices[i])))        
    return weights

def get_nices():
    nices = []
    for i in range (0,40):
        nices.append(20 - i)        
    return nices

def main():
    nices = get_nices()
    weights = get_weights(nices)
    plt.figure()
    plt.plot(nices, weights)
    plt.show()



if __name__ == "__main__":
    main()