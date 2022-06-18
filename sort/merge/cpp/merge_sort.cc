#include <stdio.h>
#include <vector>
#include <string>
#include <iostream>
using namespace std;


// [3,4,2]
//
vector<int> merge_sort(vector<int> v) {

}
int main() {
    int n;
    cin >> n;
    vector<int> v(n);
    for(int i = 0; i < n; i++) {
        cin >> v[i];
    }

    // check
    for(int n:v) {
        cout << n << " ";
    }
    cout << endl;
}
