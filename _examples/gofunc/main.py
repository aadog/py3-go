gofunc=__import__("gofunc")


def main():
    print("add:{}".format(gofunc.Call("gofunc.add",1,2)))
    print("py:{}".format(gofunc.Call("gofunc.py","py string")))


if __name__ == '__main__':
    main()