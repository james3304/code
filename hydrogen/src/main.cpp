#include <iostream>
#include <fstream>
#include <sstream>
#include <optional>
#include <vector>

using namespace std;

enum class TokenType
{
    _return,
    int_lit,
    semi
};

struct Token
{
    TokenType type;
    optional<string> value {};
};

vector<Token> tokenize(const string& str)
{
    vector<Token> tokens;

    string buf;
    for (int i = 0; i < str.length(); i++)
    {
        char c = str.at(i);
        if (isalpha(c))
        {
            buf.push_back(c);
            i++;
            while (isalpha(str.at(i)))
            {
                buf.push_back(str.at(i));
                i++;
            }
            i--;

            if (buf =="return")
            {
                tokens.push_back({.type = TokenType::_return});
            }
            else
            {
                cerr << "You messed up" << endl;
                exit(EXIT_FAILURE);
            }
        }
    }
};

int main(int argc, char* argv[]) {
    if (argc != 2)
    {
        cerr << "Incorrect Usage. Use hydro <input.hy>" <<endl;
        return EXIT_FAILURE;
    }

    string contents;
    {
        stringstream contents_stream;
        fstream input(argv[1], ios::in);
        contents_stream << input.rdbuf();
        contents = contents_stream.str();
    }

    //cout << contents << endl;
    tokenize(contents);
    return EXIT_SUCCESS;
}
