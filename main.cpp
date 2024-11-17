#include "copy.h"
#include <iostream>
#include <exception> //for c++ exception
using namespace std;

int main(int argc, char *argv[])
{

    try
    {
        Copy c;
        if (argc < 3)
            throw std::runtime_error("./prog src dest");

        // validate and source
        c.validateSource(argv[1]);
        c.validateDestination(argv[2]);

        // making a copy to desired location
        c.copy();
    }
    catch (const std::exception &e)
    {
        std::cout <<"Error : "<< e.what()<<"\n";
    }
    return 0;
}