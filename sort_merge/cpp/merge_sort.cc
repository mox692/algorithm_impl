#include <stdio.h>
#include <vector>
#include <string>
#include <iostream>
using namespace std;


// [3,4,2,6]
// [4,2,6]
vector<int> merge_sort(vector<int> v) {
    int len = v.size();
    if(len == 1) {
        return v;
    }
    vector<int> l, r;
    // left 
    l = merge_sort(v);
    // right
    r = merge_sort(v);
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
    // merge_sort();
}
